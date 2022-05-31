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

func CreateResponse(ctx context.Context, ResponseData *requests.Response, sentryCtx context.Context) {
	defer sentry.Recover()
	span := sentry.StartSpan(sentryCtx, "[DAO] AddResponse")
	defer span.Finish()

	kafkaClient := kafkaFunc.InitProducer()
	topic := "Responses"
	exampleBytes, err := json.Marshal(ResponseData)
	fmt.Println(string(exampleBytes), err)

	kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(topic),
		Value:          []byte(exampleBytes),
	}, nil)

	// Wait for all messages to be delivered
	kafkaClient.Flush(10000)
	kafkaClient.Close()

	// if ResponseData.ResponseId != " " && ResponseData.Status {
	// 	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] update responses")
	// 	multiMatchQuery, err := es.Client().Update().Index("responses").Id(responseId).Script(elastic.NewScriptInline(`ctx._source.Status = true`)).Do(ctx)

	// 	dbSpan1.Finish()

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		sentry.CaptureException(err)
	// 		logger.Client().Error(err.Error())
	// 		return "null", err
	// 	}

	// 	return multiMatchQuery.Id, nil
	// } else {

	// 	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /responses")
	// 	multiMatchQuery, err := es.Client().Index().Id(responseId).Index("responses").BodyJson(ResponseData).Do(ctx)
	// 	dbSpan1.Finish()

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		sentry.CaptureException(err)
	// 		logger.Client().Error(err.Error())
	// 		return "null", err
	// 	}

	// 	return multiMatchQuery.Id, nil
	// }

}
