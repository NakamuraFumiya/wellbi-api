package create

import (
	"net/http"
	"wellbi-api/core/usecase/roadmap"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	RoadmapUseCase roadmap.RoadmapUseCase
}

func NewController(useCase roadmap.RoadmapUseCase) Controller {
	return Controller{useCase}
}

func (c Controller) CreateRoadmap(ctx echo.Context) error {
	post, err := c.RoadmapUseCase.CreateRoadmap(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, post)
}
