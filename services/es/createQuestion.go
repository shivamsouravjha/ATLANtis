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

func CreateQuestion(ctx context.Context, QuestionData *requests.Question, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddQuestion")
	defer span.Finish()

	kafkaClient := kafkaFunc.InitProducer()
	topic := "Questions"

	data := structs.QuestionKafka{
		Data:     QuestionData,
		IsUpdate: false,
	}

	exampleBytes, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(&data)

	kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(topic),
		Value:          []byte(exampleBytes),
	}, nil)

	// Wait for all messages to be delivered
	kafkaClient.Flush(15000)
	kafkaClient.Close()

	// dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /questions")
	// QuestionInsert, err := es.Client().Index().Id(questionID).Index("questions").BodyJson(decodedStr).Do(ctx)
	// dbSpan1.Finish()

	// if err != nil {
	// 	sentry.CaptureException(err)
	// 	logger.Client().Error(err.Error())
	// 	return "null", err
	// }

}
