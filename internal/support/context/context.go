package context

import (
	"gin_tonic/internal/response/baseResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Context *gin.Context
}

func (response *Response) CheckInternalServerError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		response.Context.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
}

func (response *Response) CheckBadRequestError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		response.Context.AbortWithStatusJSON(http.StatusBadRequest, result)
	}
}

func (response *Response) CheckStatusConflictError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		response.Context.AbortWithStatusJSON(http.StatusConflict, result)
	}
}

func (response *Response) SuccessStatusCreated(result gin.H) {
	response.Context.JSON(http.StatusCreated, result)
}

func (response *Response) SuccessStatusOK(result gin.H) {
	response.Context.JSON(http.StatusOK, result)
}
