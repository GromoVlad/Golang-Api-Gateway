package findBook

import (
	"fmt"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
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

	url := strings.Replace(fmt.Sprintf("%s", ginContext.Request.URL), "/api-gateway", "", -1)
	response, err := http.Get(os.Getenv("MICROSERVICE_BOOKS_URL") + url)
	context.InternalServerError(err)

	buffer, err := io.ReadAll(response.Body)
	context.DetermineStatus(response.StatusCode, buffer)
	ginContext.Writer.Write(buffer)
}
