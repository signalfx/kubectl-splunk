// Copyright 2020 Splunk, Inc.
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
	"errors"

	"github.com/spf13/cobra"

	"github.com/signalfx/kubectl-signalfx/pkg/kubectl"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get [flags] [kubernetes resource] -- [kubectl flags]",
	Aliases: []string{"g"},
	Short:   "Get a resource filtered by " + Selector,
	Long: `This command can be used to retrieve any Kubernetes resource (e.g. pods, daemonsets, configmaps)
that is automatically filtered by ` + Selector + `.

Any additional flags are passed directly through to the underlying kubectl get command.

Examples:
	# Get list of all agent pods.
	kubectl signalfx get pods

	# Get list of all agent daemonsets.
	kubectl signalfx get daemonset

	# Standard kubectl aliases still work.
	kubectl signalfx get po

	# Additional kubectl flags can be passed by using --.
	kubectl signalfx get po -- -o yaml
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a Kubernetes resource type (pod, daemonset, etc.)")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return kubectl.Exec(append([]string{
			"get", "--all-namespaces=true", "--selector", Selector, "--output=wide"}, args...))
	},
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
