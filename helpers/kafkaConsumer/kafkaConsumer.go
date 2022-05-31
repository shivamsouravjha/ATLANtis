package kafkaFunc

import (
	"Atlantis/config"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var clientConsumer *kafka.Consumer

func NewConsumerClient(topicName string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Get().KafkaServer,
		"group.id":          topicName,
		"auto.offset.reset": "earliest",
		"sasl.mechanisms":   config.Get().Mechanisms,
		"security.protocol": config.Get().Protocol,
		"sasl.username":     config.Get().Username,
		"sasl.password":     config.Get().Password,
	})
	return consumer, err
}

func InitConsumer(topicName string) *kafka.Consumer {
	clientConsumer, err = NewConsumerClient(topicName)
	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	subscriptionError := clientConsumer.SubscribeTopics([]string{topicName}, nil)
	if subscriptionError != nil {
		fmt.Printf("Failed to create consumer: %s", subscriptionError)
		os.Exit(1)
	}

	return clientConsumer
}
