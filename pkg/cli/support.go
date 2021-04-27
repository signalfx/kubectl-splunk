// Copyright 2014 The Kubernetes Authors.
// Modifications Copyright 2020 Splunk, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"archive/zip"
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/kubectl/pkg/polymorphichelpers"

	"github.com/signalfx/kubectl-splunk/pkg/kubectl"
)

const supportFile = "signalfx-support.zip"

var (
	accessor     = meta.NewAccessor()
	apiResources = []string{"pods", "daemonsets", "configmaps", "clusterroles", "clusterrolebindings", "serviceaccounts"}
	// From https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/logs/logs.go
	containerNameFromRefSpecRegexp = regexp.MustCompile(`spec\.(?:initContainers|containers|ephemeralContainers){(.+)}`)
)

// supportCmd represents the support command
var supportCmd = &cobra.Command{
	Use:   "support",
	Short: "Collects Kubernetes resources into a zip file",
	Long:  "Collects Kubernetes resources into a zip file",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Create(supportFile)
		if err != nil {
			return fmt.Errorf("unable to create support.zip file: %v", err)
		}
		defer file.Close()
		z := zip.NewWriter(file)

		if err := logs(z); err != nil {
			return err
		}

		if err := resourceFiles(z); err != nil {
			return err
		}

		if err := z.Close(); err != nil {
			return fmt.Errorf("unable to close zip: %v", err)
		}

		return nil
	},
}

func logs(z *zip.Writer) error {
	if err := resource.NewBuilder(kubectl.CfgFlags).
		ResourceTypes("pods").
		AllNamespaces(true).
		LabelSelector(Selector).
		WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
		ContinueOnError().
		Flatten().
		Do().
		Visit(func(info *resource.Info, err error) error {
			res, err := polymorphichelpers.LogsForObjectFn(
				kubectl.CfgFlags,
				info.Object,
				&v1.PodLogOptions{},
				20*time.Second,
				true,
			)
			if err != nil {
				return err
			}

			for ref, wrapper := range res {
				log, err := wrapper.DoRaw(context.Background())
				if err != nil {
					return err
				}

				// From https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/logs/logs.go#L389
				containerName := "unknown"
				containerNameMatches := containerNameFromRefSpecRegexp.FindStringSubmatch(ref.FieldPath)
				if len(containerNameMatches) == 2 {
					containerName = containerNameMatches[1]
				}

				pth := strings.Join([]string{ref.Namespace, "pods", ref.Name, containerName + ".log"}, "/")
				zipEntry, err := z.Create(pth)
				if err != nil {
					return err
				}

				if _, err := zipEntry.Write(log); err != nil {
					return err
				}
			}

			return nil
		}); err != nil {
		return err
	}

	return nil
}

func resourceFiles(z *zip.Writer) error {
	serializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme,
		scheme.Scheme, json.SerializerOptions{
			Yaml:   true,
			Pretty: true,
		})

	if err := resource.NewBuilder(kubectl.CfgFlags).
		ResourceTypes(apiResources...).
		AllNamespaces(true).
		LabelSelector(Selector).
		WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
		ContinueOnError().
		Flatten().
		Do().
		Visit(func(info *resource.Info, err error) error {
			obj := info.Object
			namespace, err := accessor.Namespace(obj)
			if err != nil {
				return fmt.Errorf("could not determine namespace: %v", err)
			}

			if namespace == "" {
				namespace = "global"
			}

			name, err := accessor.Name(obj)
			if err != nil {
				return fmt.Errorf("could not determine name: %v", err)
			}

			pth := strings.Join([]string{namespace, info.Mapping.Resource.Resource, name, "spec.yml"}, "/")

			zipEntry, err := z.Create(pth)
			if err != nil {
				return fmt.Errorf("failed creating zip entry: %v", err)
			}

			if err := serializer.Encode(obj, zipEntry); err != nil {
				return fmt.Errorf("failed encoding runtime object: %v", err)
			}

			return nil
		}); err != nil {
		return fmt.Errorf("visiting resources failed: %v", err)
	}

	return nil
}

func init() {
	RootCmd.AddCommand(supportCmd)
}
