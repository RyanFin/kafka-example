package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		log.Fatal("Error creating Kafka producer: ", err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     "test-topic",
		Value:     sarama.StringEncoder("Hello from Go Kafka!!!"),
		Timestamp: time.Now(),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Failed to send message: ", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
