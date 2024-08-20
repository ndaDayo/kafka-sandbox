package main

import (
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:29092",
	})

	if err != nil {
		panic(err)
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

	topic := "first-app"
	for i := 1; i <= 100; i++ {
		key := i
		value := strconv.Itoa(i)

		record := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(strconv.Itoa(key)),
			Value:          []byte(value),
		}

		err := p.Produce(record, nil)
		if err != nil {
			fmt.Printf("Produce failed: %v\n", err)
		}
	}

	p.Flush(15 * 1000)
}
