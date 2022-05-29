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
type UpdateRedisRequest struct {
	HandleName string `form:"handleName"`
	Status     int    `form:"status"`
}

type GetCreatorDetailsRequest struct {
	HandleName string `form:"handleName" binding:"required"`
}

type HandleDetailsRequest struct {
	HandleName string      `form:"handleName" binding:"required"`
	UserId     interface{} `form:"userId"`
}

type HandleDetailsAdminRequest struct {
	UserId interface{} `form:"userId"`
}

type UpdateAvatarRequest struct {
	Email string `form:"userEmail" binding:"required"`
}
