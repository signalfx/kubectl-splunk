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

	"github.com/signalfx/kubectl-splunk/pkg/kubectl"
)

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	Use:     "describe [flags] [kubernetes resource] -- [kubectl flags]",
	Aliases: []string{"de", "des", "desc"},
	Short:   "Describe a resource filtered by " + Selector,
	Long: `This command can be used to describe any Kubernetes resource (e.g. pods, daemonsets, configmaps)
that is automatically filtered by ` + Selector + `.

Any additional flags are passed directly through to the underlying kubectl describe command.

Examples:
	# Describe all agent pods.
	kubectl signalfx describe pods

	# Describe all agent daemonsets.
	kubectl signalfx describe daemonset

	# Standard kubectl aliases still work.
	kubectl signalfx describe po

	# Additional kubectl flags can be passed by using --.
	kubectl signalfx describe po -- -o yaml
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a Kubernetes resource type (pod, daemonset, etc.)")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return kubectl.Exec(append([]string{
			"describe", "--all-namespaces=true", "--selector", Selector}, args...))
	},
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
}

func init() {
	RootCmd.AddCommand(describeCmd)
}
