package router

import (
	"echo-twitter-clone/presentation/api/inject"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo, i inject.Injector) {
	// Routes
	e.GET("/", hello)

	api := e.Group("/api")
	api.GET("/posts", i.V1PostListController())
	api.GET("/posts/:id", i.V1PostDetailController())
	api.POST("/posts", i.V1PostCreateController())
	api.PUT("/posts/:id", i.V1PostUpdateController())
	api.DELETE("posts/:id", i.V1PostDeleteController())
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a Twitter clone!")
}
