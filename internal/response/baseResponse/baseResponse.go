package baseResponse

type BaseResponse struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}
