package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:29092",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	defer p.Close()
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "myTopic"
	for _, word := range []string{"Hello!!!!!", "ndaDayo!!!!"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	p.Flush(15 * 1000)
}
