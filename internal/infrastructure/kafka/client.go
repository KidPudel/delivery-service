package kafka

import (
	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	Writer *kafka.Writer
}

func NewKafkaClient() *KafkaClient {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "quickstart-events",
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaClient{
		Writer: writer,
	}
}
