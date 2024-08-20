package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "broker:29092",
		"group.id":           "FirstAppConsumerGroup",
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{"first-app"}, nil)
	if err != nil {
		panic(err)
	}

	defer c.Close()

	for count := 0; count < 300; count++ {

		msg, err := c.ReadMessage(1000 * time.Millisecond)

		if err != nil {
			if kafkaError, ok := err.(kafka.Error); ok && kafkaError.Code() == kafka.ErrTimedOut {
				continue
			}
			fmt.Printf("Consumer error: %v\n", err)
			continue
		}

		fmt.Printf("key:%s, value:%s, topic:%s, partition:%d, offset:%d\n",
			string(msg.Key), string(msg.Value), *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)

		_, err = c.CommitMessage(msg)
		if err != nil {
			fmt.Printf("Failed to commit message: %v\n", err)
		}

		time.Sleep(1 * time.Second)
	}
}
