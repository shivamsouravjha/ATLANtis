package response

type InsertEventResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type EventResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
