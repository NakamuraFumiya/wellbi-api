package delete

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

func (h Controller) DeletePost(c echo.Context) error {
	err := h.PostUseCase.DeletePost(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "post can not delete")
	}
	return nil
}
