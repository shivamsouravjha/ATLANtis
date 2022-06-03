package es

import (
	kafkaFunc "Atlantis/helpers/kafkaConsumer"
	"Atlantis/structs"
	"Atlantis/structs/requests"
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
)

func CreateResponse(ctx context.Context, ResponseData *requests.Response, isUpdate bool, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddResponse")
	defer span.Finish()

	kafkaClient := kafkaFunc.InitProducer()
	topic := "Responses"

	data := structs.ResponseKafka{
		Data:     ResponseData,
		IsUpdate: true,
	}
	exampleBytes, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&data)

	kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(topic),
		Value:          []byte(exampleBytes),
	}, nil)

	kafkaClient.Flush(10000)
	kafkaClient.Close()

}
