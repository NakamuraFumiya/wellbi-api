package controller

import (
	"echo-twitter-clone/core/usecase/post"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostHandler interface {
	GetPosts(c echo.Context) error
	GetPost(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type postHandler struct {
	PostUseCase post.PostUseCase
}

func NewPostHandler(u post.PostUseCase) postHandler {
	return postHandler{u}
}

func (h postHandler) GetPosts(c echo.Context) error {
	posts, err := h.PostUseCase.GetPosts(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "posts does not exist")
	}
	return c.JSON(http.StatusOK, posts)
}

func (h postHandler) GetPost(c echo.Context) error {
	post, err := h.PostUseCase.GetPost(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "post does not exist")
	}
	return c.JSON(http.StatusOK, post)
}

func (h postHandler) CreatePost(c echo.Context) error {
	post, err := h.PostUseCase.CreatePost(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, post)
}

func (h postHandler) UpdatePost(c echo.Context) error {
	err := h.PostUseCase.UpdatePost(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "post can not update")
	}
	return nil
}

func (h postHandler) DeletePost(c echo.Context) error {
	err := h.PostUseCase.DeletePost(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "post can not delete")
	}
	return nil
}
