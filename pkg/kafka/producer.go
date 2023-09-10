package kafka

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string, username string, password string) *Producer {
	mechanism, err := scram.Mechanism(scram.SHA256, username, password)
	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
		Dialer:  dialer,
	})

	return &Producer{writer: writer}
}

func (p *Producer) PublishMessage(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() {
	p.writer.Close()
}
