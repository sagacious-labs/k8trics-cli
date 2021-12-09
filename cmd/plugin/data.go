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

// dataCmd represents the data command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "data can be used to stream data from the plugin",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("only 1 argument is required")
		}

		if err := k8trics.New(viper.GetString("host")).Plugin().StreamData(args[0]); err != nil {
			fmt.Println(err)
		}

		return nil
	},
}

func init() {
	PluginCmd.AddCommand(dataCmd)
}
