package createBook

import (
	"context"
	"fmt"
	protobuf "gin_tonic/internal/controllers/apiGateway/book/createBook/gRPC"
	"gin_tonic/internal/gRPC"
	createBookRequest "gin_tonic/internal/requests/book/createBook"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Endpoint - Создать запись о книге
// CreateBook godoc
// @Summary      Создать запись о книге
// @Tags         Api Gateway Books
// @Produce      json
// @Param  		 RequestBody  body  DTO  true  "Тело запроса"
// @Success      201  {object}  baseResponse.Response
// @Router       /api-gateway/book [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := createBookRequest.GetRequest(context)

	connection := gRPC.Connection()
	defer connection.Close()

	client := protobuf.NewCreateBookClient(connection)
	response := createBook(client, request)

	if response.Success == false {
		context.BadRequestError(fmt.Errorf(response.Message))
	}

	result := baseResponse.BaseResponse{
		Data:    baseResponse.Response{Status: "Книга создана"},
		Success: response.Success,
	}
	context.StatusCreated(gin.H{"data": result.Data, "success": result.Success})
}

func createBook(client protobuf.CreateBookClient, request *protobuf.Request) *protobuf.Response {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	book, err := client.CreateBook(context, request)
	if err != nil {
		log.Fatalf("%v \n", err.Error())
	}

	return book
}
