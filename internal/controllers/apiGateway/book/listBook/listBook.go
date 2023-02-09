package listBook

import (
	"context"
	protobuf "gin_tonic/internal/controllers/apiGateway/book/listBook/gRPC"
	"gin_tonic/internal/gRPC"
	listBookRequest "gin_tonic/internal/requests/book/listBook"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Endpoint - Возвращает пагинированый список книг
// ListBook godoc
// @Summary      Возвращает пагинированый список книг
// @Tags         Api Gateway Books
// @Produce      json
// @Param  		 page   	query	int	 	false	"Номер страницы"  minimum(1)
// @Param  		 limit  	query	int	 	false	"Кол-во записей на странице" minimum(1)	maximum(20)
// @Param  		 book_id  	query	int		false	"Идентификатор книги"
// @Param  		 name  		query	string	false	"Поиск по названию книги"
// @Param  		 author_id  query	int		false	"Идентификатор автора"
// @Param  		 category  	query	string	false	"Категория"
// @Success      200  {object}  listBookResponse.Response
// @Router       /api-gateway/book/list [get]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	request := listBookRequest.GetRequest(context)

	connection := gRPC.Connection()
	defer connection.Close()

	client := protobuf.NewListBookClient(connection)
	response := listBook(client, request)

	result := Response{Data: response, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}

func listBook(client protobuf.ListBookClient, request *protobuf.Request) *protobuf.Response {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	books, err := client.ListBook(context, request)
	if err != nil {
		log.Fatalf("%v \n", err.Error())
	}

	return books
}

type Response struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}
