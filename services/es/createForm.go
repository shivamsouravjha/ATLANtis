package es

import (
	kafkaFunc "Atlantis/helpers/kafkaConsumer"
	"Atlantis/structs"
	"Atlantis/structs/requests"
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
)

func CreateForm(ctx context.Context, FormData *requests.Form, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddForm")
	defer span.Finish()

	if FormData.Name == "" {
		FormData.Name = "No Title"
	}

	kafkaClient := kafkaFunc.InitProducer()
	topic := "Forms"

	data := structs.FormKafka{
		Data:     FormData,
		IsUpdate: false,
	}

	exampleBytes, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&data)
	fmt.Println(err)
	kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(topic),
		Value:          []byte(exampleBytes),
	}, nil)

	kafkaClient.Flush(10000)
	kafkaClient.Close()

}
