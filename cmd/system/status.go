/*
Copyright © 2021 Utkarsh

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
	"strconv"

	"github.com/fatih/color"
	"github.com/sagacious-labs/kcli/pkg/kubectl"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status shows the status of the K8trics components",
	Run: func(cmd *cobra.Command, args []string) {
		if checkHyperion() {
			color.Green("✅ Hyperion Daemons are running and are in Healthy state")
		} else {
			color.Red("❌ Hyperion Daemons are not in healthy state")
		}

		if checkK8trics() {
			color.Green("✅ K8trics server is running and is in Healthy state")
		} else {
			color.Red("❌ K8trics servier is not in a healthy state")
		}
	},
}

func init() {
	SystemCmd.AddCommand(statusCmd)
}

func checkHyperion() bool {
	args := []string{
		"get",
		"daemonsets.apps",
		"-n",
		"hyperion",
		"hyperion-daemon",
	}

	// Get number of daemons scheduled
	sout, serr, err := kubectl.GenericExec(
		append(args, "-o=jsonpath={$.status.currentNumberScheduled}"),
	)
	if err != nil || serr != "" {
		return false
	}

	currentScheduled, err := strconv.Atoi(sout)
	if err != nil {
		return false
	}

	// Get number of daemons scheduled
	sout, serr, err = kubectl.GenericExec(
		append(args, "-o=jsonpath={$.status.desiredNumberScheduled}"),
	)
	if err != nil || serr != "" {
		return false
	}

	desiredScheduled, err := strconv.Atoi(sout)
	if err != nil {
		return false
	}

	return currentScheduled >= desiredScheduled
}

func checkK8trics() bool {
	args := []string{
		"get",
		"deployments",
		"-n",
		"k8trics",
		"k8trics",
	}

	// Get number of ready replicas
	sout, serr, err := kubectl.GenericExec(
		append(args, "-o=jsonpath={$.status.readyReplicas}"),
	)
	if err != nil || serr != "" {
		return false
	}

	readyReplicas, err := strconv.Atoi(sout)
	if err != nil {
		return false
	}

	// Get number of desired replicas
	sout, serr, err = kubectl.GenericExec(
		append(args, "-o=jsonpath={$.status.replicas}"),
	)
	if err != nil || serr != "" {
		return false
	}

	desiredReplicas, err := strconv.Atoi(sout)
	if err != nil {
		return false
	}

	return readyReplicas >= desiredReplicas
}
