package internal

import (
	"Tortuga/internal/Infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type App struct {
	router *gin.Engine
}

func BuildApp() (App, error) {
	err := Infrastructure.ReadConfig()

	if err != nil {
		return App{}, err
	}

	r := Infrastructure.CreateRouter()

	_, err = Infrastructure.ConnectToDatabase()

	if err != nil {
		return App{}, nil
	}

	return App{r}, nil
}

func (app *App) StartHttp() error {
	return app.router.Run(fmt.Sprintf(":%s", viper.GetString("HTTP_PORT")))
}
