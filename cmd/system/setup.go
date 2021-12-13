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
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/sagacious-labs/kcli/pkg/kubectl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		color.Red("❌ ", err.Error())
		return
	}

	if serr != "" {
		color.Red("❌ Something went wrong while starting K8trics and its components")
		return
	}

	color.Green("✅ Successfully started K8trics in Kubernetes Cluster")

	ep, err := findK8tricsLBEndpoint(30 * time.Second)
	if err != nil {
		color.Red("❌ Failed to get location of K8trics API Server")
		return
	}

	port, err := findK8tricsLBEndpointPort(30 * time.Second)
	if err != nil {
		color.Red("❌ Failed to get location of K8trics API Server")
		return
	}

	viper.Set("host", fmt.Sprintf("%s:%s", ep, port))
	if err := viper.WriteConfig(); err != nil {
		color.Red("❌ Failed to write config file")
		return
	}

	color.Green("✅ Successfully written config")
	color.Green("✅ K8trics is ready to be used")
}

func findK8tricsLBEndpoint(timeout time.Duration) (string, error) {
	sleeper := 1 * time.Second
	start := time.Now()

	for {
		stdout, stderr, err := kubectl.GenericExec([]string{
			"get",
			"services",
			"-n",
			"k8trics",
			"k8trics",
			"-o=jsonpath={$.status.loadBalancer.ingress[0].ip}",
		})

		if stderr != "" || err != nil || stdout == "" {
			if start.Add(timeout).Before(time.Now()) {
				// Sleep as the timeout time hasn't reached yet
				time.Sleep(sleeper)

				// Try again
				continue
			}

			if stderr != "" {
				return "", fmt.Errorf(stderr)
			}

			if err != nil {
				return "", err
			}
		}

		return stdout, nil
	}
}

func findK8tricsLBEndpointPort(timeout time.Duration) (string, error) {
	sleeper := 1 * time.Second
	start := time.Now()

	for {
		stdout, stderr, err := kubectl.GenericExec([]string{
			"get",
			"services",
			"-n",
			"k8trics",
			"k8trics",
			"-o=jsonpath={$.spec.ports[0].port}",
		})

		if stderr != "" || err != nil {
			if start.Add(timeout).Before(time.Now()) {
				// Sleep as the timeout time hasn't reached yet
				time.Sleep(sleeper)

				// Try again
				continue
			}

			if stderr != "" {
				return "", fmt.Errorf(stderr)
			}

			if err != nil {
				return "", err
			}
		}

		return stdout, nil
	}
}
