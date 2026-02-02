package kafka

import (
	"context"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	logger *slog.Logger
}

func NewConsumer(brokers []string, topic, groupID string, logger *slog.Logger) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			GroupID:  groupID,
			Topic:    topic,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
		logger: logger,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	defer c.reader.Close()

	c.logger.Info("Kafka Consumer started", "topic", c.reader.Config().Topic)

	for {

		m, err := c.reader.ReadMessage(ctx)
		if err != nil {

			if ctx.Err() != nil {
				return
			}
			c.logger.Error("failed to read message", "err", err)
			continue
		}

		c.logger.Info("message received",
			"topic", m.Topic,
			"partition", m.Partition,
			"offset", m.Offset,
			"key", string(m.Key),
			"value", string(m.Value),
		)
	}
}
