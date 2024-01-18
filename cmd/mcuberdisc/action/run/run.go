package run

import (
	"context"
	"github.com/spf13/viper"
	"github.com/uptrace/uptrace-go/uptrace"
	"go.uber.org/zap"
	"tyr.codes/tyr/mcuberdisc/cmd/mcuberdisc/action"
	"tyr.codes/tyr/mcuberdisc/internal/config"
	logic "tyr.codes/tyr/mcuberdisc/internal/logic/v1"
)

var Run action.Action = func(ctx context.Context, args []string) error {
	uptrace.ConfigureOpentelemetry(
		uptrace.WithServiceName(viper.GetString(config.Keys.ApplicationName)),
		uptrace.WithServiceVersion(viper.GetString(config.Keys.SoftwareVersion)),
	)
	ctx, cancel := context.WithCancel(ctx)

	logicMod := logic.New()
	watcher := logicMod.NewLogWatcher(viper.GetString(config.Keys.LogFilePath))
	if err := watcher.Watch(ctx); err != nil {
		zap.L().Error("Error watching log file", zap.Error(err))
		cancel()
		return err
	}

	cancel()
	return nil
}
