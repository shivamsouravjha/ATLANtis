package structs

import "Atlantis/structs/requests"

type AnswerKafka struct {
	Data     *requests.Answer
	IsUpdate bool
}
type FormKafka struct {
	Data     *requests.Form
	IsUpdate bool
}
type ResponseKafka struct {
	Data     *requests.Response
	IsUpdate bool
}
type QuestionKafka struct {
	Data     *requests.Question
	IsUpdate bool
}
