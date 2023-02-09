package updateBook

import (
	protobuf "gin_tonic/internal/controllers/apiGateway/book/updateBook/gRPC"
	"gin_tonic/internal/support/localContext"
	"github.com/golang/protobuf/ptypes/wrappers"
	"strconv"
)

type DTO struct {
	Name        string `form:"name"                   json:"name"                    binding:"required"`
	Category    string `form:"category"               json:"category"                binding:"required"`
	AuthorId    int    `form:"author_id"              json:"author_id"               binding:"required,number"`
	Description string `form:"description,omitempty"  json:"description,omitempty"   binding:"omitempty"`
}

func GetRequest(context localContext.LocalContext) *protobuf.Request {
	var dto DTO
	err := context.Context.ShouldBindJSON(&dto)
	context.BadRequestError(err)

	bookId, err := strconv.Atoi(context.Context.Param("bookId"))
	context.BadRequestError(err)

	request := &protobuf.Request{
		BookId:      int32(bookId),
		Name:        dto.Name,
		Category:    dto.Category,
		AuthorId:    int32(dto.AuthorId),
		Description: nil,
	}

	if dto.Description != "" {
		request.Description = &wrappers.StringValue{Value: dto.Description}
	}

	return request
}
