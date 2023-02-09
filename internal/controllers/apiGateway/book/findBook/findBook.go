package findBook

import (
	"context"
	"fmt"
	protobuf "gin_tonic/internal/controllers/apiGateway/book/findBook/gRPC"
	"gin_tonic/internal/gRPC"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// Endpoint - Найти книгу по идентификатору
// FindBook godoc
// @Summary      Найти книгу
// @Tags         Api Gateway Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор пользователя"
// @Success      200  	 {object}  	findBookResponse.Response
// @Router       /api-gateway/book/{bookId} [get]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	bookId, err := strconv.Atoi(context.Context.Param("bookId"))
	context.BadRequestError(err)

	connection := gRPC.Connection()
	defer connection.Close()

	client := protobuf.NewFindBookClient(connection)
	response := findBook(client, &protobuf.Request{BookId: int32(bookId)})

	if response.BookId == 0 {
		context.NotFoundError(fmt.Errorf(response.ErrorMessage))
	}

	result := Response{Data: response, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}

func findBook(client protobuf.FindBookClient, request *protobuf.Request) *protobuf.Response {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	book, err := client.FindBook(context, request)
	if err != nil {
		log.Fatalf("%v \n", err.Error())
	}

	return book
}

type Response struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}
