package storage

import (
	"context"
	"time"
)

type Storage interface {
	Increment(ctx context.Context, key string) (int64, error)
	SetExpiration(ctx context.Context, key string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
