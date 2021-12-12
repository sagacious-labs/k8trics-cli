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
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/sagacious-labs/kcli/cmd/plugin"
	"github.com/sagacious-labs/kcli/cmd/system"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kcli",
	Short: "kcli - K8trics CLI allows to interact with K8trics API Server",
	Long: `kcli - K8trics CLI allows to interact with K8trics API Server.

kcli bundles some default Hyperion Wodules (plugins) but can be used to spin up any compatible plugin`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.k8trics.yaml)")

	setupCommands()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("KCLI")

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".k8trics")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		color.Yellow("⚠️  Failed to read config file")

		color.Yellow("⚠️  Attempting to create default config file...")
		if err := setupDefaultConfig(); err != nil {
			color.Red("❌ Failed to create default config file")
			os.Exit(1)
		}
	}
}

func setupCommands() {
	rootCmd.AddCommand(plugin.PluginCmd)
	rootCmd.AddCommand(system.SystemCmd)
}

func setupDefaultConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(home, ".k8trics.yaml"), nil, 0644)
}
