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

	// Wait for all messages to be delivered
	kafkaClient.Flush(10000)
	kafkaClient.Close()

	// if AnswerData.AnswerID != "" {
	// 	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] update answer")
	// 	multiMatchQuery, err := es.Client().Update().Index("answers").Id(answerId).Doc(map[string]interface{}{"Answer": AnswerData.Answer}).Do(ctx)

	// 	dbSpan1.Finish()

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		sentry.CaptureException(err)
	// 		logger.Client().Error(err.Error())
	// 		return "null", err
	// 	}

	// 	return multiMatchQuery.Id, nil
	// } else {

	// 	dbSpan1 := sentry.StartSpan(span.Context(), "[DB] Insert into /answer")
	// 	multiMatchQuery, err := es.Client().Index().Id(answerId).Index("answers").BodyJson(AnswerData).Do(ctx)
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
