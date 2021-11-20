package app

import (
	"../handler"
	"../repository"
	"github.com/labstack/echo"
)

func Run(mode byte, path string) error {
	r, err := repository.InitRepo(mode, path)
	if err != nil {
		return err
	}

	router := echo.New()
	group := router.Group("/api")
	h := handler.NewHandler(r)
	h.InitHandler(group)

	if err := router.Start("localhost:9000"); err != nil {
		return err
	}

	return nil
}