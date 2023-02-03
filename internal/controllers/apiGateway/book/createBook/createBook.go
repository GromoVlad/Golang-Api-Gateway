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
// @Param  		 RequestBody  body  DTO  true  "Тело запроса"
// @Success      201  {object}  Response
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

type DTO struct {
	Name        string `form:"name"                   json:"name"                    binding:"required"`
	Category    string `form:"category"               json:"category"                binding:"required"`
	AuthorId    int    `form:"author_id"              json:"author_id"               binding:"required,number"`
	Description string `form:"description,omitempty"  json:"description,omitempty"   binding:"omitempty"`
}

type Response struct {
	Success bool `json:"success"`
}
