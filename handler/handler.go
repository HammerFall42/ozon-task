package handler

import (
	"log"
	"net/http"

	"github.com/HammerFall42/ozon-task/repository"

	"github.com/labstack/echo"
)

type AddInput struct {
	Url	string	`json:"url"`
}

type GetInput struct {
	Shortened	string	`json:"shortened"`
}

type Handler struct {
	repo repository.Repo
}

func NewHandler(r repository.Repo) *Handler {
	return &Handler{repo: r}
}

func (h *Handler) InitHandler(api *echo.Group) {
	v1 := api.Group("/v1")
	{
		v1.POST("/url", h.AddUrl)
		v1.GET("/url", h.GetUrl)
	}
}

func (h *Handler) AddUrl(ctx echo.Context) error {
	var input AddInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	shortened, err := h.repo.CallAddNewUrl(input.Url)

	if err != nil {
		log.Println("failed to add new url")
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Println("new url added")

	return ctx.JSON(http.StatusOK, map[string]string{"shortened": shortened})
}

func (h *Handler) GetUrl(ctx echo.Context) error {
	var input GetInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	url, err := h.repo.CallGetUrl(input.Shortened)

	if err != nil {
		log.Println("failed to get url")
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	log.Println("url found")

	return ctx.JSON(http.StatusOK, map[string]string{"url": url})
}