package update

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

func (c Controller) UpdatePost(ctx echo.Context) error {
	err := c.PostUseCase.UpdatePost(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "post can not update")
	}
	return nil
}
