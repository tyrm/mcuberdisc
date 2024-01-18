package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"tyr.codes/tyr/mcuberdisc/cmd/mcuberdisc/action"
	"tyr.codes/tyr/mcuberdisc/cmd/mcuberdisc/flag"
	"tyr.codes/tyr/mcuberdisc/internal/config"
)

// Version is the software version.
var Version string

// Commit is the git commit.
var Commit string

func main() {
	// init logger
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		_ = zapLogger.Sync()
	}()
	zap.ReplaceGlobals(zapLogger)

	// set software version
	var v string
	if len(Commit) < 7 {
		v = "v" + Version
	} else {
		v = "v" + Version + "-" + Commit[:7]
	}

	viper.Set(config.Keys.SoftwareVersion, v)

	rootCmd := &cobra.Command{
		Use:           "mcuderdisc",
		Version:       v,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	flag.Global(rootCmd, config.Defaults)

	err = viper.BindPFlag(config.Keys.ConfigPath, rootCmd.PersistentFlags().Lookup(config.Keys.ConfigPath))
	if err != nil {
		zap.L().Fatal("Error binding config flag", zap.Error(err))

		return
	}

	// add commands
	rootCmd.AddCommand(runCommand())

	err = rootCmd.Execute()
	if err != nil {
		zap.L().Fatal("Error executing command", zap.Error(err))
	}
}

func preRun(cmd *cobra.Command) error {
	if err := config.Init(cmd.Flags()); err != nil {
		return fmt.Errorf("error initializing config: %s", err)
	}

	if err := config.ReadConfigFile(); err != nil {
		return fmt.Errorf("error reading config: %s", err)
	}

	return nil
}

func run(ctx context.Context, action action.Action, args []string) error {
	return action(ctx, args)
}
