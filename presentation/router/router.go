package router

import (
	"echo-twitter-clone/presentation/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo, h controller.AppHandler) {
	// Routes
	e.GET("/", hello)

	api := e.Group("/api")
	api.GET("/posts", h.GetPosts)
	api.POST("/posts", h.CreatePost)
	api.GET("/posts/:id", h.GetPost)
	api.PUT("/posts/:id", h.UpdatePost)
	api.DELETE("posts/:id", h.DeletePost)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a Twitter clone!")
}
