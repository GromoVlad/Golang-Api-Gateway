package deleteBook

import (
	"context"
	"fmt"
	protobuf "gin_tonic/internal/controllers/apiGateway/book/deleteBook/gRPC"
	"gin_tonic/internal/gRPC"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// Endpoint - Удалить запись о книге
// UpdateBook godoc
// @Summary      Удалить запись о книге
// @Tags         Api Gateway Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор книги"
// @Success      200  {object}  Response
// @Router       /api-gateway/book/{bookId} [delete]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	bookId, err := strconv.Atoi(context.Context.Param("bookId"))
	context.BadRequestError(err)
	request := &protobuf.Request{BookId: int32(bookId)}

	connection := gRPC.Connection()
	defer connection.Close()

	client := protobuf.NewDeleteBookClient(connection)
	response := deleteBook(client, request)

	if response.ErrorMessage != "" {
		context.NotFoundError(fmt.Errorf(response.ErrorMessage))
	}

	result := Response{
		Data:    fmt.Sprintf("Книга с идентификатором = %d удалена", bookId),
		Success: response.Success,
	}

	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}

func deleteBook(client protobuf.DeleteBookClient, request *protobuf.Request) *protobuf.Response {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.DeleteBook(context, request)
	if err != nil {
		log.Fatalf("%v \n", err.Error())
	}

	return response
}

type Response struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}
