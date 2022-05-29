package requests

type Question struct {
	Question        string `form:"question,omitempty" binding:"required"`
	Form            string `form:"form,omitempty" binding:"required"`
	ResponseType    string `form:"responseType,omitempty" binding:"required"`
	Order           int16  `form:"orderNumber,omitempty" binding:"required"`
	Option          string `form:"option,omitempty"`
	MediaAndContact string `form:"mediaAndContact,omitempty"`
	Feedback        string `form:"feedback,omitempty"`
}

type Form struct {
	UserID int64  `form:"userId" binding:"required"`
	Name   string `form:"formName"`
}

//question
//answer
//response
//form
type Response struct {
	UserID     int64  `form:"userId" binding:"required"`
	FormID     string `form:"formId" binding:"required"`
	ResponseId string `form:"responseId,omitempty"`
	Status     bool   `form:"status,omitempty"`
}
