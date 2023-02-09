package baseErrorResponse

type BaseResponse struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}

type Response struct {
	Message string `json:"error_message"`
}
