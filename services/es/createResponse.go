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
