/*
Copyright © 2021 Utkarsh Srivastava

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
package system

import (
	"github.com/fatih/color"
	"github.com/sagacious-labs/kcli/pkg/kubectl"
	"github.com/spf13/cobra"
)

// setupCmd represents the start command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup installs Hyperion agents and K8trics server in the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	SystemCmd.AddCommand(setupCmd)
}

func start() {
	if !kubectl.Exists() {
		color.Red("kubectl not found in $PATH")
		return
	}

	_, serr, err := kubectl.Apply([]string{getHyperionManifest(), getK8tricsManifest()})
	if err != nil {
		color.Red(err.Error())
		return
	}

	if serr != "" {
		color.Red("Something went wrong while starting K8trics and its components")
		return
	}

	color.Green("Successfully started K8trics in Kubernetes Cluster")
}
