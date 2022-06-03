package response

import "Atlantis/structs/requests"

type InsertEventResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type EventResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type RespResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    Response `json:"data"`
}
type AnyResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}
type FormResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Form   `json:"data"`
}
type AnswerResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    requests.Answer `json:"data"`
}
type QuestionResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    requests.Question `json:"data"`
}
type Form struct {
	requests.Form
	Question []requests.Question
}
type Response struct {
	requests.Response
	QandA []UnitResponse
}

type UnitResponse struct {
	Question requests.Question
	Answer   requests.Answer
}
