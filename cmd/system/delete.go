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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete removes Hyperion Agents and K8trics server from the Kubernetes Cluster",
	Run: func(cmd *cobra.Command, args []string) {
		delete()
	},
}

func init() {
	SystemCmd.AddCommand(deleteCmd)
}

func delete() {
	if !kubectl.Exists() {
		color.Red("kubectl not found in $PATH")
		return
	}

	_, serr, err := kubectl.Delete([]string{getHyperionManifest(), getK8tricsManifest()})
	if err != nil {
		color.Red(err.Error())
		return
	}

	if serr != "" {
		color.Red("Something went wrong while deleted K8trics and its components")
		return
	}

	color.Green("Successfully deleted K8trics from Kubernetes Cluster")
}
