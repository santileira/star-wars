package cmd

import (
	"rings/cmd/server"
	"github.com/spf13/cobra"
)

func Cmds() *cobra.Command {
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(server.NewRunnable().Cmd())
	return rootCmd
}
