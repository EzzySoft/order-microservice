package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type DBConfig struct {
	Host     string `env:"DB_HOST,default=localhost"`
	Port     int    `env:"DB_PORT,default=5432"`
	User     string `env:"DB_USER,default=orders"`
	Password string `env:"DB_PASSWORD,default=example"`
	Name     string `env:"DB_NAME,default=orders"`
	SSLMode  string `env:"DB_SSL_MODE,default=disable"`
}

func LoadDB(ctx context.Context) (*DBConfig, error) {
	var c DBConfig

	return &c, envconfig.Process(ctx, &c)
}

func (c DBConfig) DSN() string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SSLMode,
	)
}
