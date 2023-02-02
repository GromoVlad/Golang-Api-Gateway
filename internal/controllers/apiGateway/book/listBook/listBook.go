package listBook

import (
	"fmt"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
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
// @Success      200  {object}  listBook.Response
// @Router       /api-gateway/book/list [get]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	url := strings.Replace(fmt.Sprintf("%s", ginContext.Request.URL), "/api-gateway", "", -1)
	response, err := http.Get(os.Getenv("MICROSERVICE_BOOKS_URL") + url)
	context.InternalServerError(err)

	buffer, err := io.ReadAll(response.Body)
	context.DetermineStatus(response.StatusCode, buffer)
	ginContext.Writer.Write(buffer)
}
