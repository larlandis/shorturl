package config

import (
	"context"
)

func GetString(ctx context.Context, key string) string {
	if cfg, ok := ctx.Value(ctxKey{}).(Config); ok {
		return cfg.c.GetString(key)
	}
	return ""
}

func GetInt(ctx context.Context, key string) uint {
	if cfg, ok := ctx.Value(ctxKey{}).(Config); ok {
		return cfg.c.GetUint(key)
	}
	return 0
}
