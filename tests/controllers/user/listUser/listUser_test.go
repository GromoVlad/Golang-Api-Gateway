package listUser

import (
	"gin_tonic/internal/controllers/user/listUser"
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

var context localContext.LocalContext
var accessToken string

func init() {
	/** Подгружаем данные из .env */
	if err := godotenv.Load("../../../../.env"); err != nil {
		log.Fatalf("Ошибка загрузки переменных из .env: %s", err.Error())
	}
	context = localContext.LocalContext{Context: &gin.Context{}}
}

// Flow: CreateUser -> FindUser -> Login -> TestListUser -> DeleteUser
func TestListUser(t *testing.T) {
	setUp()

	router := gin.Default()
	router.GET("/user/list", func(c *gin.Context) {
		listUser.Endpoint(c)
	})

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/user/list", nil)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	router.ServeHTTP(writer, request)

	assert.Equal(t, 200, writer.Code)
	assert.Equal(
		t,
		true,
		strings.Contains(writer.Body.String(), `"email":"`+createUser.EMAIL+`"`),
	)

	deleteUser.DeleteUser(context) // tearDown
}

func setUp() {
	createUser.CreateUser(context)
	user := findUser.FindUser(context)
	accessToken, _, _ = loginService.Login(context, user)
}
