package kafkaFunc

import (
	"Atlantis/config"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var client *kafka.Producer
var err error

func NewProducerClient() (*kafka.Producer, error) {
	fmt.Println(config.Get().KafkaServer, config.Get().Mechanisms, config.Get().Protocol, config.Get().Username, config.Get().Password)
	consumer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Get().KafkaServer,
		"sasl.mechanisms":   config.Get().Mechanisms,
		"security.protocol": config.Get().Protocol,
		"sasl.username":     config.Get().Username,
		"sasl.password":     config.Get().Password,
	})
	return consumer, err
}

func InitProducer() *kafka.Producer {
	client, err = NewProducerClient()
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	return client
}
