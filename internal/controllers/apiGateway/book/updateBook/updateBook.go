package updateBook

import (
	"context"
	"fmt"
	protobuf "gin_tonic/internal/controllers/apiGateway/book/updateBook/gRPC"
	"gin_tonic/internal/gRPC"
	updateBookRequest "gin_tonic/internal/requests/book/updateBook"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Endpoint - Обновить запись о книге
// UpdateBook godoc
// @Summary      Обновить запись о книге
// @Tags         Api Gateway Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор книги"
// @Param  		 RequestBody  body  updateBook.DTO  true	"Тело запроса"
// @Success      200  {object}  updateBook.Response
// @Router       /api-gateway/book/{bookId} [put]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := updateBookRequest.GetRequest(context)

	connection := gRPC.Connection()
	defer connection.Close()

	client := protobuf.NewUpdateBookClient(connection)
	response := updateBook(client, request)

	if response.ErrorMessage != "" {
		context.NotFoundError(fmt.Errorf(response.ErrorMessage))
	}

	result := Response{Data: response, Success: response.Success}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}

func updateBook(client protobuf.UpdateBookClient, request *protobuf.Request) *protobuf.Response {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	book, err := client.UpdateBook(context, request)
	if err != nil {
		log.Fatalf("%v \n", err.Error())
	}

	return book
}

type Response struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}
