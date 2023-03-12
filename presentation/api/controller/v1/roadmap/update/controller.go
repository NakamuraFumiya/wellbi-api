package update

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

func (c Controller) UpdateRoadmap(ctx echo.Context) error {
	err := c.RoadmapUseCase.UpdateRoadmap(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "roadmap can not update")
	}
	return nil
}
