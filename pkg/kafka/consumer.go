package kafka

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic, groupID, username, password string) *KafkaConsumer {
	mechanism, err := scram.Mechanism(scram.SHA256, username, password)
	if err != nil {
		log.Fatalf("Error creating scram mechanism: %v", err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		// GroupID:  groupID,
		Topic:    topic,
		Dialer:   dialer,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		// StartOffset: kafka.LastOffset,
	})

	return &KafkaConsumer{
		reader: reader,
	}
}

func (kc *KafkaConsumer) Start(callback func([]byte) error) error {
	fmt.Println("Consumer started. Waiting for messages...")
	for {
		m, err := kc.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			return err
		}

		// Call the callback to process the message
		if err := callback(m.Value); err != nil {
			log.Printf("Error processing message: %v", err)
			// decide if you want to continue or exit
			continue
		}
	}
}

func (kc *KafkaConsumer) Close() {
	kc.reader.Close()
}
