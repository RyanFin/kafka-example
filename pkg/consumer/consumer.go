package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

// func main() {
// 	brokers := []string{"localhost:9092"}

// 	consumer, err := sarama.NewConsumer(brokers, nil)
// 	if err != nil {
// 		log.Fatal("Error creating consumer: ", err)
// 	}

// 	defer consumer.Close()

// 	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
// 	if err != nil {
// 		log.Fatal("Error creating partition consumer: ", err)
// 	}

// 	defer partitionConsumer.Close()

// 	fmt.Println("listening for messages...")

// 	for msg := range partitionConsumer.Messages() {
// 		fmt.Printf("Received message: %s\n", msg.Value)
// 	}

// }

func main() {
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // Start from the beginning

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}
	defer partitionConsumer.Close()

	fmt.Println("Reading old messages...")
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Message received: %s\n", string(msg.Value))
	}
}
