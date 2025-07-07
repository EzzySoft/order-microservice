package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type KafkaConfig struct {
	Brokers []string `env:"KAFKA_BROKERS, required, delimiter=;"`
	Topic   string   `env:"KAFKA_TOPIC, default=orders"`
	GroupID string   `env:"KAFKA_GROUP_ID, default=order-service"`
	Version string   `env:"KAFKA_VERSION, default=4.0.0"`
}

func LoadKafka(ctx context.Context) (*KafkaConfig, error) {
	var c KafkaConfig
	return &c, envconfig.Process(ctx, &c)
}
