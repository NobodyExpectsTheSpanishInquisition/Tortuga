package internal

import (
	. "Tortuga/internal/Infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	router    *gin.Engine
	container *Container
}

func BuildApp() (App, error) {
	err := ReadConfig()

	if err != nil {
		return App{}, err
	}

	container := &Container{}
	container.Init()

	r := CreateRouter(container)

	_, err = ConnectToDatabase()

	if err != nil {
		return App{}, nil
	}

	return App{r, container}, nil
}

func (app *App) StartHttp() error {
	return app.router.Run(fmt.Sprintf(":%s", viper.GetString("HTTP_PORT")))
}
