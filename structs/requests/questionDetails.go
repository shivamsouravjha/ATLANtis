package requests

type Question struct {
	QuestionID      string `json:"questionId,omitempty"`
	Question        string `json:"question,omitempty" binding:"required"`
	Form            string `json:"form,omitempty" binding:"required"`
	ResponseType    string `json:"responseType,omitempty" binding:"required"`
	Order           int16  `json:"orderNumber,omitempty" binding:"required"`
	Option          string `json:"option,omitempty"`
	MediaAndContact string `json:"mediaAndContact,omitempty"`
	Feedback        string `json:"feedback,omitempty"`
}

type Form struct {
	UserID int64  `json:"userId" binding:"required"`
	Name   string `json:"formName"`
	FormID string `json:"formId"`
}
type GetQuestion struct {
	QuestionID string `json:"questionId" binding:"required"`
}

type GetForm struct {
	FormID string `json:"formId" binding:"required"`
}

type GetResponse struct {
	ResponseId string `json:"responseId" binding:"required"`
}

type GetAnswer struct {
	AnswerID string `json:"answerId" binding:"required"`
}
type Response struct {
	UserID     int64  `json:"userId" binding:"required"`
	FormID     string `json:"formId" binding:"required"`
	ResponseId string `json:"responseId,omitempty"`
	Status     bool   `json:"status,omitempty"`
}

type Answer struct {
	FormID     string   `json:"formId" binding:"required"`
	ResponseId string   `json:"responseId"  binding:"required"`
	Answer     []string `json:"answer" binding:"required"`
	AnswerType string   `json:"answerType" binding:"required"`
	AnswerID   string   `json:"answerId"`
}
