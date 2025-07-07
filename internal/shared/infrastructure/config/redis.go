package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type RedisConfig struct {
	Host     string `env:"REDIS_HOST, default=localhost"`
	Port     int    `env:"REDIS_PORT, default=6379"`
	Password string `env:"REDIS_PASSWORD, default=password"`
	DB       int    `env:"REDIS_DB, default=0"`
}

func LoadRedis(ctx context.Context) (*RedisConfig, error) {
	var c RedisConfig

	return &c, envconfig.Process(ctx, &c)
}

func (c RedisConfig) Addr() string {

	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
