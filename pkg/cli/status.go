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

var pod string
var all bool

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Runs the status command in the specified pod.",
	Args: func(cmd *cobra.Command, args []string) error {
		if all {
			return errors.New("--all is unimplemented")
		}

		if all && pod != "" {
			return errors.New("both --all and --pod cannot be set")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO --all. Make calls to exec in-process so we can reuse connections.
		if err := kubectl.Spawn(append([]string{"exec", pod, "--", "agent-status"}, args...)); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	statusCmd.Flags().StringVarP(&pod, "pod", "p", "", "agent pod name")
	statusCmd.Flags().BoolVarP(&all, "all", "a", false, "run command on all agent pods")
	// XXX: Disabled, status command not available in OT.
	//RootCmd.AddCommand(statusCmd)
}
