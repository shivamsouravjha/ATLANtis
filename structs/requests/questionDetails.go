package requests

type Question struct {
	Question     string `form:"question" binding:"required"`
	Form         string `form:"form" binding:"required"`
	ResponseType string `form:"responseType" binding:"required"`
	Order        string `form:"orderNumber" binding:"required"`
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
