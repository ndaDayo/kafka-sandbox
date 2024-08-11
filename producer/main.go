package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func startProducer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		panic(err)
	}

	defer p.Close()
}
