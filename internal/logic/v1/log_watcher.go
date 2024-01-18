package v1

import (
	"context"
	"github.com/hpcloud/tail"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"regexp"
	"sync"
)

var tracerAttrsLogWatcher = []trace.SpanStartOption{
	trace.WithAttributes(
		attribute.String("type", "LogWatcher"),
	),
}

type LogWatcher struct {
	filepath string
	logic    *Logic
}

var (
	reLogLine = regexp.MustCompile(`^\[([0-9]+):([0-9]+):([0-9]+)\] \[(.+)/(.+)\]: (.+)$`)

	reJoined = regexp.MustCompile(`^([a-zA-Z0-9_-]+) joined the game$`)
	reLeft   = regexp.MustCompile(`^([a-zA-Z0-9_-]+) left the game$`)
	reChat   = regexp.MustCompile(`^<([a-zA-Z0-9_-]+)> (.+)$`)
)

func (lw *LogWatcher) Watch(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)

	t, err := tail.TailFile(lw.filepath, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		zap.L().Error("Error tailing file", zap.Error(err))
		cancel()
		return err
	}
	defer t.Cleanup()

	var wg sync.WaitGroup
	logLines := make(chan string, 255)

	wg.Add(1)
	go lw.processor(ctx, &wg, logLines)

	for {
		select {
		case reason := <-ctx.Done():
			// Context is cancelled, stop processing.
			zap.L().Debug("Context canceled, exiting", zap.Any("reason", reason))
			cancel()
			return nil
		case line, ok := <-t.Lines:
			if !ok {
				// Channel is closed, stop processing.
				zap.L().Debug("Channel closed, exiting")
				cancel()
				return nil
			}

			if !reLogLine.MatchString(line.Text) {
				zap.L().Debug("Line skip", zap.String("line", line.Text))
				continue
			}

			zap.L().Debug("Line read", zap.String("line", line.Text))
			logLines <- line.Text
		}
	}
}

func (lw *LogWatcher) processor(ctx context.Context, wg *sync.WaitGroup, logLines <-chan string) {
	defer wg.Done()

	for {
		select {
		case reason := <-ctx.Done():
			// Context is cancelled, stop processing.
			zap.L().Debug("Context canceled, exiting", zap.Any("ctx", reason))
			return
		case line, ok := <-logLines:
			if !ok {
				// Channel is closed, stop processing.
				zap.L().Debug("Channel closed, exiting")
				return
			}

			// parse line
			matches := reLogLine.FindStringSubmatch(line)
			if err := lw.processLine(ctx, matches[6]); err != nil {
				zap.L().Warn("Error processing line", zap.Error(err), zap.String("line", line), zap.Strings("matches", matches))
				continue
			}
		}
	}
}

func (lw *LogWatcher) processLine(ctx context.Context, line string) error {
	_, span := tracer.Start(ctx, "processLine", tracerAttrsLogWatcher...)
	defer span.End()

	switch {
	case reJoined.MatchString(line):
		// player joined
		player := reJoined.FindStringSubmatch(line)[1]

		return lw.logic.PlayerJoined(ctx, player)
	case reLeft.MatchString(line):
		// player left
		player := reLeft.FindStringSubmatch(line)[1]

		return lw.logic.PlayerLeft(ctx, player)
	case reChat.MatchString(line):
		// chat
		matches := reChat.FindStringSubmatch(line)
		player := matches[1]
		message := matches[2]

		return lw.logic.PlayerChat(ctx, player, message)
	default:
		// unknown
		zap.L().Debug("Unknown line", zap.String("line", line))
		return nil
	}
}
