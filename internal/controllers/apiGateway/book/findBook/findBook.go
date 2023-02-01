package findBook

import (
	"encoding/json"
	"fmt"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/support/localContext"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Endpoint - Найти книгу по идентификатору
// FindBook godoc
// @Summary      Найти книгу
// @Tags         Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор пользователя"
/** @Success      200  	 {object}  	findBook.Response */
// @Router       /api-gateway/book/{bookId} [get]
func Endpoint(ginContext *gin.Context) {
	fmt.Println("Запрос пришел")
	context := localContext.LocalContext{Context: ginContext}
	bookId, err := strconv.Atoi(ginContext.Param("bookId"))
	context.BadRequestError(err)

	response, err := http.Get("http://mcrsrv-book:8083/book/" + strconv.Itoa(bookId))
	context.InternalServerError(err)

	var result baseResponse.BaseResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	context.InternalServerError(err)

	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}
