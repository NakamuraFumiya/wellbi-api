package list

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

func (c Controller) GetRoadmaps(ctx echo.Context) error {
	posts, err := c.RoadmapUseCase.GetRoadmaps(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "posts does not exist")
	}
	return ctx.JSON(http.StatusOK, posts)
}
