package helpers

import (
	"Atlantis/services/es"
	"Atlantis/services/logger"
	"Atlantis/structs"
	"context"
	"fmt"

	"github.com/getsentry/sentry-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic/v7"
)

func EsUploader(topic string, data []byte) {

	if topic == "Answers" {
		var input structs.AnswerKafka
		fmt.Println("vdd")
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(data, &input)
		fmt.Println(input.IsUpdate)
		if input.IsUpdate {
			_, err := es.Client().Update().Index("answers").Id(input.Data.AnswerID).Doc(map[string]interface{}{"Answer": input.Data.Answer}).Do(context.Background())
			if err != nil {
				fmt.Println(err)
				return
			}

			return
		} else {

			_, err := es.Client().Index().Id(input.Data.AnswerID).Index("answers").BodyJson(input.Data).Do(context.Background())

			if err != nil {
				fmt.Println(err)
				return
			}

			return
		}
	} else if topic == "Forms" {
		var input structs.FormKafka
		fmt.Println("vdd")
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(data, &input)
		fmt.Println(input.IsUpdate)
		_, err := es.Client().Index().Id(input.Data.FormID).Index("forms").BodyJson(input.Data).Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if topic == "Responses" {
		var input structs.ResponseKafka
		fmt.Println("vdd")
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(data, &input)
		fmt.Println(input.IsUpdate)
		if input.IsUpdate && input.Data.Status {
			_, err := es.Client().Update().Index("responses").Id(input.Data.ResponseId).Script(elastic.NewScriptInline(`ctx._source.Status = true`)).Do(context.Background())

			if err != nil {
				fmt.Println(err)
				return
			}

			return
		} else {
			_, err := es.Client().Index().Id(input.Data.ResponseId).Index("responses").BodyJson(input.Data).Do(context.Background())

			if err != nil {
				fmt.Println(err)
				sentry.CaptureException(err)
				logger.Client().Error(err.Error())
				return
			}

			return
		}
	} else if topic == "Questions" {
		var input structs.QuestionKafka
		fmt.Println("vdd")
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(data, &input)
		fmt.Println(input.IsUpdate) // if AnswerData.AnswerID != "" {
		_, err := es.Client().Index().Id(input.Data.QuestionID).Index("questions").BodyJson(input.Data).Do(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
