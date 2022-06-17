package Infrastructure

import (
	auth "Tortuga/internal/auth/login/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateRouter(container *Container) *gin.Engine {
	r := gin.Default()

	registerRoutes(r, container)

	return r
}

func registerRoutes(r *gin.Engine, container *Container) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	login(r, container)
}

func login(r *gin.Engine, container *Container) {
	r.GET("/api/auth/login", func(context *gin.Context) {
		request := auth.LoginRequest{}
		err := context.BindJSON(&request)
		if nil != err {
			fmt.Errorf(err.Error())
		}

		token, err = auth.Login(request, container.PasswordValidator(), container.EmailValidator(), container)

		handleError(err)
	})
}

func handleError(err error) {
	if nil != err {
		fmt.Errorf(err.Error())
	}
}
