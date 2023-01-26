package registrationUser_test

import (
	"bytes"
	"gin_tonic/internal/controllers/user/registrationUser"
	"gin_tonic/internal/enums/role"
	"gin_tonic/internal/support/localContext"
	"gin_tonic/testHelpers/createUser"
	"gin_tonic/testHelpers/deleteUser"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
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

// Flow: TestRegistrationUser -> DeleteUser
func TestRegistrationUser(t *testing.T) {
	router := gin.Default()
	router.POST("/user/registration", func(c *gin.Context) {
		registrationUser.Endpoint(c)
	})

	writer := httptest.NewRecorder()
	body := bytes.NewReader([]byte(
		`{"name": "` + createUser.USERNAME + `",` +
			`"email": "` + createUser.EMAIL + `",` +
			`"role_id": ` + strconv.Itoa(role.SUPPORT) + `,` +
			`"phone": "89998887766",` +
			`"venue_id": ` + strconv.Itoa(rand.Intn(100)) + `,` +
			` "password": "12345678"}`,
	))
	request, _ := http.NewRequest(http.MethodPost, "/user/registration", body)
	router.ServeHTTP(writer, request)

	assert.Equal(t, 201, writer.Code)
	assert.Equal(
		t,
		true,
		strings.Contains(writer.Body.String(), `{"data":{"status":"Пользователь создан"},"success":true}`),
	)

	deleteUser.DeleteUser(context) // tearDown
}
