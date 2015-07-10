package commands

import (
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "github.com/MustWin/go-base/app"
)

// Options
var Port int

func buildServerCommand() *cobra.Command {
  var serverCommand = &cobra.Command{
    Use:   "server",
    Short: "Start your application",
    Long: `Start your application, takes options:`,
    Run: func(cmd *cobra.Command, args []string) {
      app.Main()
    },
  }
  serverCommand.Flags().IntVarP(&Port, "port", "p", 8888, "The port to listen on")
  viper.BindPFlag("port", serverCommand.Flags().Lookup("port"))
  return serverCommand
}
