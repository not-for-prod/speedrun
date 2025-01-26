package main

import (
	"github.com/not-for-prod/speedrun/cmd/crud"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "speedrun",
		Short: "Speedrun most annoying parts of development",
		Long:  `Speedrun is a bunch of dirty hacks that removes most boring parts of development by generating code that is not perfect but works.`,
	}
)

func main() {
	// fill root with other cmds
	rootCmd.AddCommand(crud.Cmd)

	_ = rootCmd.Execute()
}
