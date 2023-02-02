package localContext

import (
	"errors"
	"gin_tonic/internal/response/baseResponse"
	"gin_tonic/internal/support/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LocalContext struct {
	Context *gin.Context
}

func (localContext *LocalContext) InternalServerError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		logger.ErrorLog("InternalServerError", err.Error())
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusInternalServerError, result)
		panic("InternalServerError")
	}
}

func (localContext *LocalContext) NotFoundError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		logger.ErrorLog("NotFound", err.Error())
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusNotFound, result)
		panic("NotFoundError")
	}
}

func (localContext *LocalContext) AlreadyExistsError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		logger.ErrorLog("AlreadyExists", err.Error())
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusForbidden, result)
		panic("AlreadyExistsError")
	}
}

func (localContext *LocalContext) BadRequestError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		logger.ErrorLog("BadRequest", err.Error())
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusBadRequest, result)
		panic("BadRequestError")
	}
}

func (localContext *LocalContext) TokenError(err error) {
	localContext.TokenExpiredError(err)
	localContext.InvalidTokenError(err)
	localContext.BadRequestError(err)
}

func (localContext *LocalContext) TokenExpiredError(err error) {
	if err != nil && err.Error() == "Token is expired" {
		data := baseResponse.Response{Status: "Произошла ошибка: Истекло время жизни токена"}
		logger.ErrorLog("TokenExpired", "Истекло время жизни токена")
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(700, result)
		panic("TokenExpiredError")
	}
}

func (localContext *LocalContext) InvalidTokenError(err error) {
	if err != nil && err.Error() == "signature is invalid" {
		data := baseResponse.Response{Status: "Произошла ошибка: Невалидная структура токена"}
		logger.ErrorLog("InvalidToken", "Невалидная структура токена")
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(701, result)
		panic("InvalidTokenError")
	}
}

func (localContext *LocalContext) PasswordError(err error) {
	localContext.WrongPasswordError(err)
	localContext.BadRequestError(err)
}

func (localContext *LocalContext) WrongPasswordError(err error) {
	if err != nil && err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
		data := baseResponse.Response{Status: "Произошла ошибка: Передан некорректный пароль"}
		logger.ErrorLog("WrongPassword", " Передан некорректный пароль")
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusUnauthorized, result)
		panic("WrongPasswordError")
	}

	localContext.BadRequestError(err)
}

func (localContext *LocalContext) StatusConflictError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		logger.ErrorLog("StatusConflict", err.Error())
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusConflict, result)
		panic("StatusConflictError")
	}
}

func (localContext *LocalContext) UnauthorizedError(err error) {
	if err != nil {
		data := baseResponse.Response{Status: "Произошла ошибка: " + err.Error()}
		logger.ErrorLog("Unauthorized", err.Error())
		result := baseResponse.BaseResponse{Data: data, Success: false}
		localContext.Context.AbortWithStatusJSON(http.StatusUnauthorized, result)
		panic("UnauthorizedError")
	}
}

func (localContext *LocalContext) StatusCreated(result gin.H) {
	localContext.Context.JSON(http.StatusCreated, result)
}

func (localContext *LocalContext) StatusOK(result gin.H) {
	localContext.Context.JSON(http.StatusOK, result)
}

func (localContext *LocalContext) DetermineStatus(code int, buffer []byte) {
	errorBuffer := errors.New(string(buffer))
	switch code {
	case 400:
		localContext.BadRequestError(errorBuffer)
	case 401:
		localContext.UnauthorizedError(errorBuffer)
	case 403:
		localContext.AlreadyExistsError(errorBuffer)
	case 404:
		localContext.NotFoundError(errorBuffer)
	case 409:
		localContext.StatusConflictError(errorBuffer)
	case 500:
		localContext.InternalServerError(errorBuffer)
	case 700:
		localContext.TokenExpiredError(errorBuffer)
	case 701:
		localContext.InvalidTokenError(errorBuffer)
	}
}
