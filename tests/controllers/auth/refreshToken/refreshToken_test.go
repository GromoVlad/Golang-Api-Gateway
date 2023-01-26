package refreshToken

import (
	"bytes"
	"gin_tonic/internal/controllers/auth/refreshToken"
	"gin_tonic/internal/service/auth/loginService"
	"gin_tonic/internal/support/localContext"
	"gin_tonic/testHelpers/createUser"
	"gin_tonic/testHelpers/deleteUser"
	"gin_tonic/testHelpers/findUser"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const BaseTestUrl = "/auth/refresh-token"

var context localContext.LocalContext
var refreshTokenString string

func init() {
	/** Подгружаем данные из .env */
	if err := godotenv.Load("../../../../.env"); err != nil {
		log.Fatalf("Ошибка загрузки переменных из .env: %s", err.Error())
	}
	context = localContext.LocalContext{Context: &gin.Context{}}
}

// Flow: CreateUser -> GetRefreshToken -> Test -> DeleteAuthToken -> DeleteUser
func TestRefreshToken(t *testing.T) {
	setUp()

	router := gin.Default()
	router.POST(BaseTestUrl, func(c *gin.Context) {
		refreshToken.Endpoint(c)
	})

	writer := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"refresh_token": "` + refreshTokenString + `"}`))
	request, _ := http.NewRequest(http.MethodPost, BaseTestUrl, body)
	router.ServeHTTP(writer, request)

	assert.Equal(t, 200, writer.Code)
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "data"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "access_token"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "refresh_token"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "expires_at"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "\"success\":true}"))

	deleteUser.DeleteUser(context) // tearDown
}

func setUp() {
	createUser.CreateUser(context)
	user := findUser.FindUser(context)
	_, refreshTokenString, _ = loginService.Login(context, user)
}
