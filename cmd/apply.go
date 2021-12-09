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
package cmd

import (
	"fmt"

	"github.com/sagacious-labs/kcli/pkg/k8trics"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	locations []string
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a configuration to every Hyperion",
	Long:  `Apply a configuration to every Hyperion - Command will request K8trics Server to forward request to every Hyperion Agent in the cluster`,
	Example: `
  # Apply a single file
  kcli plugin apply -f /path/to/file
  
  # Apply multiple files
  kcli plugin apply -f /path/to/file1/ -f /path/to/file2`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := k8trics.New(viper.GetString("host")).Plugin().Apply(locations); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	pluginCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringArrayVarP(&locations, "files", "f", nil, "Paths to the files which are to be applied")

	applyCmd.MarkFlagRequired("files")
}
