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

func CreateForm(ctx context.Context, FormData *requests.Form, formID string, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddForm")
	defer span.Finish()

	if FormData.Name == "" {
		FormData.Name = "No Title"
	}

	kafkaClient := kafkaFunc.InitProducer()
	topic := "Forms"
	exampleBytes, err := json.Marshal(FormData)
	fmt.Println(string(exampleBytes), err)

	kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(formID),
		Value:          []byte(exampleBytes),
	}, nil)

	// Wait for all messages to be delivered
	kafkaClient.Flush(10000)
	kafkaClient.Close()

	// dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /forms")
	// _, err := es.Client().Index().Id(formID).Index("forms").BodyJson(FormData).Do(ctx)
	// dbSpan1.Finish()

	// if err != nil {
	// 	sentry.CaptureException(err)
	// 	logger.Client().Error(err.Error())
	// 	return
	// }

}
