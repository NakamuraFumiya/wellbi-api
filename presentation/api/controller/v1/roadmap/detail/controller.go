package detail

import (
	"echo-twitter-clone/core/usecase/roadmap"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	RoadmapUseCase roadmap.RoadmapUseCase
}

func NewController(useCase roadmap.RoadmapUseCase) Controller {
	return Controller{useCase}
}

func (c Controller) GetRoadmap(ctx echo.Context) error {
	post, err := c.RoadmapUseCase.GetRoadmap(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "roadmap does not exist")
	}
	return ctx.JSON(http.StatusOK, post)
}
