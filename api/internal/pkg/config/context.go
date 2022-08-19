package config

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	v    *viper.Viper
)

type (
	config interface {
		GetString(key string) string
		GetUint(key string) uint
	}
	Config struct {
		c config
	}
	ctxKey struct{}
)

func Init(fileName string) {
	once.Do(func() {
		v = viper.New()
		wd, _ := os.Getwd()
		v.SetConfigFile(wd + fileName)
		if err := v.ReadInConfig(); err != nil {
			log.Panic(err)
		}
	})
}

func WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey{}, Config{c: v})
}
