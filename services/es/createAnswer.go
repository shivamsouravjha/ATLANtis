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

func CreateAnswer(ctx context.Context, AnswerData *requests.Answer, isUpdate bool, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddAnswer")
	defer span.Finish()

	kafkaClient := kafkaFunc.InitProducer()
	topic := "Answers"

	data := structs.AnswerKafka{
		Data:     AnswerData,
		IsUpdate: isUpdate,
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
