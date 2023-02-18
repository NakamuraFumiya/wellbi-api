package post

import (
	"echo-twitter-clone/core/usecase/post"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostController interface {
	GetPosts(c echo.Context) error
	GetPost(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type postController struct {
	PostUseCase post.PostUseCase
}

func NewPostController(u post.PostUseCase) postController {
	return postController{u}
}

func (c postController) GetPosts(ctx echo.Context) error {
	posts, err := c.PostUseCase.GetPosts(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "posts does not exist")
	}
	return ctx.JSON(http.StatusOK, posts)
}

func (c postController) GetPost(ctx echo.Context) error {
	post, err := c.PostUseCase.GetPost(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "post does not exist")
	}
	return ctx.JSON(http.StatusOK, post)
}

func (c postController) CreatePost(ctx echo.Context) error {
	post, err := c.PostUseCase.CreatePost(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, post)
}

func (c postController) UpdatePost(ctx echo.Context) error {
	err := c.PostUseCase.UpdatePost(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "post can not update")
	}
	return nil
}

func (h postController) DeletePost(c echo.Context) error {
	err := h.PostUseCase.DeletePost(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "post can not delete")
	}
	return nil
}
