/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"fmt"
	"github.com/signalfx/kubectl-signalfx/pkg/kubectl"
	"github.com/spf13/cobra"
	"io/ioutil"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
	"os"
	"path"
)

// supportCmd represents the support command
var supportCmd = &cobra.Command{
	Use:   "support",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := ioutil.TempDir("", "kubectl-signalfx-support")
		if err != nil {
			return err
		}
		defer func() {
			if err := os.RemoveAll(dir); err != nil {
				cmd.PrintErr("failed cleaning up support directory: ", dir)
			}
		}()

		apiResources := []string{"pods", "daemonsets", "configmaps", "clusterroles", "clusterrolebindings", "serviceaccounts"}

		resources := path.Join(dir, "resources")
		if err := os.Mkdir(resources, 0644); err != nil {
			return err
		}

		builder := resource.NewBuilder(kubectl.KubectlCfgFlags)
		return builder.ResourceTypes(apiResources...).
			ContinueOnError().AllNamespaces(true).
			LabelSelector("app=signalfx-agent").
			WithScheme(scheme.Scheme, scheme.Scheme.PrioritizedVersionsAllGroups()...).
			Do().
			Visit(func(info *resource.Info, err error) error {
				fmt.Println("visiting: ", info.Mapping.GroupVersionKind.Kind)
				return nil
			})

		// TODO: helm info
	},
}

func init() {
	RootCmd.AddCommand(supportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// supportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// supportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
