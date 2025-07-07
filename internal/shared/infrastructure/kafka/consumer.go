package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	handler func([]byte) error
}

func NewConsumer(brokers []string, topic, groupID string, handler func([]byte) error) *Consumer {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1e3,
		MaxBytes: 10e6,
		MaxWait:  100 * time.Millisecond,
	})
	return &Consumer{reader: reader, handler: handler}
}

func (c *Consumer) Start(ctx context.Context) {
	log.Println("[kafka-consumer] started...")
	go func() {
		defer c.reader.Close()
		for {
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("[kafka-consumer] read error: %v", err)
				continue
			}
			log.Printf("[kafka-consumer] new message: topic=%s partition=%d offset=%d", m.Topic, m.Partition, m.Offset)
			if err := c.handler(m.Value); err != nil {
				log.Printf("[kafka-consumer] handle error: %v", err)
			}
		}
	}()
}
