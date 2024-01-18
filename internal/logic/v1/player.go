package v1

import (
	"context"
	"go.uber.org/zap"
)

func (l *Logic) PlayerJoined(ctx context.Context, player string) error {
	_, span := tracer.Start(ctx, "PlayerJoined", tracerAttrs...)
	defer span.End()

	zap.L().Debug("Player joined", zap.String("player", player))

	return nil
}

func (l *Logic) PlayerLeft(ctx context.Context, player string) error {
	_, span := tracer.Start(ctx, "PlayerLeft", tracerAttrs...)
	defer span.End()

	zap.L().Debug("Player left", zap.String("player", player))

	return nil
}

func (l *Logic) PlayerChat(ctx context.Context, player, message string) error {
	_, span := tracer.Start(ctx, "PlayerChat", tracerAttrs...)
	defer span.End()

	zap.L().Debug("Player chat", zap.String("player", player), zap.String("message", message))

	return nil
}
