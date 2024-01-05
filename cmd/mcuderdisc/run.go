package main

import (
	"github.com/spf13/cobra"
)

func runCommand() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "run",
		Short: "run service",
	}

	return serverCmd
}
