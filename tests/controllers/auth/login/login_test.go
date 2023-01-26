package login_test

import (
	"bytes"
	"gin_tonic/internal/controllers/auth/login"
	"gin_tonic/internal/support/localContext"
	"gin_tonic/testHelpers/createUser"
	"gin_tonic/testHelpers/deleteUser"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var context localContext.LocalContext

func init() {
	/** Подгружаем данные из .env */
	if err := godotenv.Load("../../../../.env"); err != nil {
		log.Fatalf("Ошибка загрузки переменных из .env: %s", err.Error())
	}
	context = localContext.LocalContext{Context: &gin.Context{}}
}

// Flow: CreateUser -> TestAuthLogin -> DeleteUser
func TestAuthLogin(t *testing.T) {
	createUser.CreateUser(context) // setUp

	router := gin.Default()
	router.POST("/auth/login", func(c *gin.Context) {
		login.Endpoint(c)
	})

	writer := httptest.NewRecorder()
	body := bytes.NewReader([]byte(`{"email": "` + createUser.EMAIL + `", "password": "` + createUser.PASSWORD + `"}`))
	request, _ := http.NewRequest(http.MethodPost, "/auth/login", body)
	router.ServeHTTP(writer, request)

	assert.Equal(t, 200, writer.Code)
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "access_token"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "data"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "refresh_token"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "expires_at"))
	assert.Equal(t, true, strings.Contains(writer.Body.String(), "\"success\":true}"))

	deleteUser.DeleteUser(context) // tearDown
}
