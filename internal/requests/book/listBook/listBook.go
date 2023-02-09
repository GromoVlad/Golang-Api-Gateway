package listBook

import (
	protobuf "gin_tonic/internal/controllers/apiGateway/book/listBook/gRPC"
	"gin_tonic/internal/support/localContext"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type DTO struct {
	Page     int    `form:"page,omitempty"   json:"page,omitempty"               binding:"omitempty,number"`
	Limit    int    `form:"limit,omitempty"     json:"limit,omitempty"       binding:"omitempty,number"`
	BookId   int    `form:"book_id,omitempty"    json:"book_id,omitempty"       binding:"omitempty,number"`
	Name     string `form:"name,omitempty"       json:"name,omitempty"            binding:"omitempty"`
	Category string `form:"category,omitempty"   json:"category,omitempty"        binding:"omitempty"`
	AuthorId int    `form:"author_id,omitempty"  json:"author_id,omitempty"       binding:"omitempty,number"`
}

func GetRequest(context localContext.LocalContext) *protobuf.Request {
	var dto DTO
	err := context.Context.ShouldBindQuery(&dto)
	context.BadRequestError(err)

	request := &protobuf.Request{
		Page:     &wrappers.Int32Value{Value: 1},
		Limit:    &wrappers.Int32Value{Value: 10},
		BookId:   nil,
		Name:     nil,
		Category: nil,
		AuthorId: nil,
	}

	if dto.Page != 0 {
		request.Page = &wrappers.Int32Value{Value: int32(dto.Page)}
	}
	if dto.Limit != 0 {
		request.Limit = &wrappers.Int32Value{Value: int32(dto.Limit)}
	}
	if dto.BookId != 0 {
		request.BookId = &wrappers.Int32Value{Value: int32(dto.BookId)}
	}
	if dto.Name != "" {
		request.Name = &wrappers.StringValue{Value: dto.Name}
	}
	if dto.Category != "" {
		request.Category = &wrappers.StringValue{Value: dto.Category}
	}
	if dto.AuthorId != 0 {
		request.AuthorId = &wrappers.Int32Value{Value: int32(dto.AuthorId)}
	}

	return request
}
