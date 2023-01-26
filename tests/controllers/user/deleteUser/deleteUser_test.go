package deleteUser

import (
	"gin_tonic/internal/controllers/user/deleteUser"
	"gin_tonic/internal/support/localContext"
	"gin_tonic/testHelpers/createUser"
	"gin_tonic/testHelpers/findUser"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

var UserId int
var context localContext.LocalContext

func init() {
	/** Подгружаем данные из .env */
	if err := godotenv.Load("../../../../.env"); err != nil {
		log.Fatalf("Ошибка загрузки переменных из .env: %s", err.Error())
	}
	ginContext := gin.Context{}
	context = localContext.LocalContext{Context: &ginContext}
}

// Flow: CreateUser -> FindUser -> TestDeleteUser
func TestDeleteUser(t *testing.T) {
	createUser.CreateUser(context)
	UserId = findUser.FindUser(context).UserId

	deleteUserRouter := gin.Default()
	deleteUserRouter.DELETE("/user/"+strconv.Itoa(UserId), func(context *gin.Context) {
		context.AddParam("userId", strconv.Itoa(UserId))
		deleteUser.Endpoint(context)
	})

	deleteUserRequest, _ := http.NewRequest(http.MethodDelete, "/user/"+strconv.Itoa(UserId), nil)
	deleteUserWriter := httptest.NewRecorder()
	deleteUserRouter.ServeHTTP(deleteUserWriter, deleteUserRequest)

	assert.Equal(t, 200, deleteUserWriter.Code)
	assert.Equal(
		t,
		true,
		strings.Contains(
			deleteUserWriter.Body.String(),
			`{"data":{"status":"Пользователь удален"},"success":true}`,
		),
	)
}
