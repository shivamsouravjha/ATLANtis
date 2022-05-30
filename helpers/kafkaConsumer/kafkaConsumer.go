package kafkaConsumer

import (
	"Atlantis/config"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var client *kafka.Consumer
var err error

func NewConsumerClient() (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Get().KafkaServer,
		"group.id":          config.Get().KafkaGroupID,
		"auto.offset.reset": "earliest",
	})
	return consumer, err
}

func Init(topicName string) *kafka.Consumer {
	client, err = NewConsumerClient()
	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	subscriptionError := client.SubscribeTopics([]string{topicName}, nil)
	if subscriptionError != nil {
		fmt.Printf("Failed to create consumer: %s", subscriptionError)
		os.Exit(1)
	}

	return client
}
