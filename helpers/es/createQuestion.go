package helpers

import (
	kafkaFunc "Atlantis/helpers/kafkaConsumer"
	"Atlantis/structs/requests"
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/getsentry/sentry-go"
)

func CreateQuestion(ctx context.Context, QuestionData *requests.Question, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddQuestion")
	defer span.Finish()

	kafkaClient := kafkaFunc.InitProducer()
	topic := "Questions"
	exampleBytes, err := json.Marshal(QuestionData)
	fmt.Println(string(exampleBytes), err)

	kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(topic),
		Value:          []byte(exampleBytes),
	}, nil)

	// Wait for all messages to be delivered
	kafkaClient.Flush(10000)
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
