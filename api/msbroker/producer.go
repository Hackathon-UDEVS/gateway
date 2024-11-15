package msbroker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

// NewProducer creates a new Kafka producer
func NewProducer(brokers []string) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{writer: writer}
}

// Produce sends a message to specified Kafka topic
func (p *Producer) Produce(ctx context.Context, topic string, key string, value interface{}) error {
	// Convert value to JSON bytes
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshaling message value: %w", err)
	}

	message := kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: jsonValue,
	}

	err = p.writer.WriteMessages(ctx, message)
	if err != nil {
		return fmt.Errorf("error writing message to kafka: %w", err)
	}

	return nil
}

// Close closes the Kafka producer
func (p *Producer) Close() error {
	return p.writer.Close()
}