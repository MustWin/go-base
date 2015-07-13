package commands

import (
  "github.com/spf13/viper"
  "github.com/spf13/cobra"
)

// Options
var config string;

func BaseCommand() *cobra.Command {
  var baseCommand = &cobra.Command{
    Use:   "gobase [command]",
    Short: "A simple production ready REST API framework",
    Long: `long usage`,
  }
/*
Supported environment variables include:
- ENVIRONMENT
*/
  viper.AutomaticEnv()

  baseCommand.Flags().StringVarP(&config, "config", "c", "config.json", "The path to a configuration file")
  viper.BindPFlag("config", baseCommand.Flags().Lookup("config"))

  // Add other commands
  baseCommand.AddCommand(buildServerCommand())
  return baseCommand
}
