package postFormData

import (
	"gin_tonic/internal/repository/userRepository"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Endpoint = func(context *gin.Context) {
	//migration := migrationRepository.FindAllMigration()
	//fmt.Printf("%#v\n%#v", migration[0], migration[1])
	//name := context.PostForm("name")
	//context.JSON(http.StatusOK, gin.H{
	//	"status":     "Hello, " + name + ". Request postFormData success",
	//	"migration1": migration[0],
	//	"migration2": migration[1],
	//})

	users := userRepository.FindAllUser()
	name := context.PostForm("name")
	context.JSON(http.StatusOK, gin.H{
		"status": "Hello, " + name + ". Request postFormData success",
		"user":   users[0],
	})
}
