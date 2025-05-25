package entity

type StandardResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
