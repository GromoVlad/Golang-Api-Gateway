package someAction

import (
	"gin_tonic/internal/controllers/someAction"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSomeAction(t *testing.T) {
	router := gin.Default()
	router.GET("/some-action", func(c *gin.Context) {
		someAction.Endpoint(c)
	})

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/some-action", nil)
	router.ServeHTTP(writer, request)

	assert.Equal(t, 200, writer.Code)
	assert.Equal(t, `{"success":true}`, writer.Body.String())
}
