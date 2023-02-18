package list

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

func (c Controller) GetPosts(ctx echo.Context) error {
	posts, err := c.PostUseCase.GetPosts(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "posts does not exist")
	}
	return ctx.JSON(http.StatusOK, posts)
}
