package deleteBook

import (
	"fmt"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

// Endpoint - Удалить запись о книге
// UpdateBook godoc
// @Summary      Удалить запись о книге
// @Tags         Api Gateway Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор книги"
// @Success      200  {object}  deleteBook.Response
// @Router       /api-gateway/book/{bookId} [delete]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	url := strings.Replace(fmt.Sprintf("%s", ginContext.Request.URL), "/api-gateway", "", -1)
	request, err := http.NewRequest(
		http.MethodDelete,
		os.Getenv("MICROSERVICE_BOOKS_URL")+url,
		nil,
	)
	context.InternalServerError(err)

	response, err := http.DefaultClient.Do(request)
	context.InternalServerError(err)

	buffer, err := io.ReadAll(response.Body)
	context.DetermineStatus(response.StatusCode, buffer)
	ginContext.Writer.Write(buffer)
}
