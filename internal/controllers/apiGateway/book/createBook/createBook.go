package createBook

import (
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// Endpoint - Создать запись о книге
// CreateBook godoc
// @Summary      Создать запись о книге
// @Tags         Api Gateway Books
// @Produce      json
// @Param  		 RequestBody  body  createBook.DTO  true  "Тело запроса"
// @Success      201  {object}  createBook.Response
// @Router       /api-gateway/book [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	response, err := http.Post(
		os.Getenv("MICROSERVICE_BOOKS_URL")+"/book",
		"application/json",
		ginContext.Request.Body,
	)
	context.InternalServerError(err)

	buffer, err := io.ReadAll(response.Body)
	context.DetermineStatus(response.StatusCode, buffer)
	ginContext.Writer.Write(buffer)
}
