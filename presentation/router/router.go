package router

import (
	"net/http"
	"wellbi-api/presentation/api/inject"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo, i inject.Injector) {
	// Routes
	e.GET("/", hello)

	api := e.Group("api")
	api.GET("/posts", i.V1RoadmapListController())
	api.GET("/posts/:id", i.V1RoadmapDetailController())
	api.POST("/posts", i.V1RoadmapCreateController())
	api.PUT("/posts/:id", i.V1RoadmapUpdateController())
	api.DELETE("/posts/:id", i.V1RoadmapDeleteController())
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a wellbi-api!")
}
