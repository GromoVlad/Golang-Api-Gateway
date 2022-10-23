package localContext

import (
	"gin_tonic/internal/response/baseResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LocalContext struct {
	Context *gin.Context
}

func (localContext *LocalContext) CheckInternalServerError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusInternalServerError, result)
		panic("CheckInternalServerError")
	}
}

func (localContext *LocalContext) CheckNotFoundError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusNotFound, result)
		panic("CheckNotFoundError")
	}
}

func (localContext *LocalContext) CheckAlreadyExistsError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusForbidden, result)
		panic("CheckAlreadyExistsError")
	}
}

func (localContext *LocalContext) CheckBadRequestError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusBadRequest, result)
		panic("CheckBadRequestError")
	}
}

func (localContext *LocalContext) CheckStatusConflictError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusConflict, result)
		panic("CheckStatusConflictError")
	}
}

func (localContext *LocalContext) SuccessStatusCreated(result gin.H) {
	localContext.Context.JSON(http.StatusCreated, result)
}

func (localContext *LocalContext) SuccessStatusOK(result gin.H) {
	localContext.Context.JSON(http.StatusOK, result)
}
