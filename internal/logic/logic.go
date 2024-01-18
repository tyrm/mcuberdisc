package logic

import "context"

type Logic interface {
	NewLogWatcher(filepath string) LogWatcher

	PlayerJoined(ctx context.Context, player string) error
	PlayerLeft(ctx context.Context, player string) error
	PlayerChat(ctx context.Context, player, message string) error
}

type LogWatcher interface {
	Watch(ctx context.Context) error
}
