package updateBook

import (
	"fmt"
	"gin_tonic/internal/models/book"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
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

	url := strings.Replace(fmt.Sprintf("%s", ginContext.Request.URL), "/api-gateway", "", -1)
	request, err := http.NewRequest(
		http.MethodPut,
		os.Getenv("MICROSERVICE_BOOKS_URL")+url,
		ginContext.Request.Body,
	)
	context.InternalServerError(err)

	response, err := http.DefaultClient.Do(request)
	context.InternalServerError(err)

	buffer, err := io.ReadAll(response.Body)
	context.DetermineStatus(response.StatusCode, buffer)
	ginContext.Writer.Write(buffer)
}

type DTO struct {
	Name        string `form:"name,omitempty"        json:"name,omitempty"          binding:"omitempty"`
	Category    string `form:"category,omitempty"    json:"category,omitempty"      binding:"omitempty"`
	AuthorId    int    `form:"author_id,omitempty"   json:"author_id,omitempty"     binding:"omitempty,number"`
	Description string `form:"description,omitempty" json:"description,omitempty"   binding:"omitempty"`
}

type Response struct {
	Data    book.Book `json:"data"`
	Success bool      `json:"success"`
}
