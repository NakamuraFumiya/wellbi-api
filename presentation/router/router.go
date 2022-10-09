package router

import (
	"echo-twitter-clone/presentation/handler"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo, h handler.AppHandler) {
	// Routes
	e.GET("/", hello)
	e.POST("/posts", h.CreatePost)
	e.GET("/posts", h.GetPosts)
	e.GET("/posts/:id", h.GetPost)
	e.PUT("/posts/:id", h.UpdatePost)
	e.DELETE("posts/:id", h.DeletePost)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a Twitter clone!")
}
