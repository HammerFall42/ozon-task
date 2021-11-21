package app

import (
	"ozon-task/handler"
	"ozon-task/repository"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func Run(mode byte) error {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	r, err := repository.InitRepo(mode)
	if err != nil {
		return err
	}

	router := echo.New()
	group := router.Group("/api")
	h := handler.NewHandler(r)
	h.InitHandler(group)

	if err := router.Start("localhost" + viper.GetString("port")); err != nil {
		return err
	}

	return nil
}