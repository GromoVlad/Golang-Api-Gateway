package listRepositoryRequest

import "github.com/gin-gonic/gin"

type Request struct {
	Page   int    `form:"page,omitempty"  json:"page,omitempty"   binding:"omitempty,number"`
	Limit  int    `form:"limit,omitempty" json:"limit,omitempty"  binding:"omitempty,number"`
	Search string `form:"search,omitempty" json:"search,omitempty"  binding:"omitempty,alphanum"`
	Offset int
}

func GetRequest(context *gin.Context) (Request, error) {
	var request Request
	if err := context.ShouldBindQuery(&request); err != nil {
		return request, err
	}
	if request.Page == 0 {
		request.Page = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}
	request.Offset = (request.Page - 1) * request.Limit
	return request, nil
}
