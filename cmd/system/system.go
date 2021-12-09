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
	"github.com/spf13/cobra"
)

const (
	hyperionManifest string = "https://raw.githubusercontent.com/sagacious-labs/k8trics/master/install/kubernetes/hyperion.yaml"
	k8tricsManifest  string = "https://raw.githubusercontent.com/sagacious-labs/k8trics/master/install/kubernetes/k8trics.yaml"
)

// SystemCmd represents the system command
var SystemCmd = &cobra.Command{
	Use:   "system",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}
