package main

import (
	"github.com/spf13/cobra"
	runcmd "tyr.codes/tyr/mcuberdisc/cmd/mcuberdisc/action/run"
)

func runCommand() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "run",
		Short: "run service",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return preRun(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context(), runcmd.Run, args)
		},
	}

	return serverCmd
}
