package run

import (
	"context"
	"github.com/spf13/viper"
	"github.com/uptrace/uptrace-go/uptrace"
	"tyr.codes/tyr/mcuberdisc/cmd/mcuderdisc/action"
	"tyr.codes/tyr/mcuberdisc/internal/config"
)

var Run action.Action = func(ctx context.Context, args []string) error {
	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(viper.GetString(config.Keys.ApplicationName)),
		uptrace.WithServiceVersion(viper.GetString(config.Keys.SoftwareVersion)),
	)
	ctx, cancel := context.WithCancel(ctx)

	cancel()
	return nil
}
