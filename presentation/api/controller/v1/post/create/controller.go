package create

import (
	"echo-twitter-clone/core/usecase/post"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	PostUseCase post.PostUseCase
}

func NewController(useCase post.PostUseCase) Controller {
	return Controller{useCase}
}

func (c Controller) CreatePost(ctx echo.Context) error {
	post, err := c.PostUseCase.CreatePost(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, post)
}
