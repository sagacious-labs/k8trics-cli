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
package plugin

import (
	"fmt"

	"github.com/sagacious-labs/kcli/pkg/k8trics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete plugins from K8trics",
	Long:  `Delete will send delete requests for plugin to K8trics server which will be forwarded to all of the Hyperion Agents`,
	Example: `
  # Delete a plugin named "NetworkWatcher"
  kcli plugin delete NetworkWatcher
  
  # Delete multiple plugins
  kcli plugin delete <plugin1> <plugin2> ...`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := k8trics.New(viper.GetString("host")).Plugin().Delete(args); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	PluginCmd.AddCommand(deleteCmd)
}
