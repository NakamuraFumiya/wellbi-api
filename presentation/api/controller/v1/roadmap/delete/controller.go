package delete

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

func (h Controller) DeleteRoadmap(c echo.Context) error {
	err := h.RoadmapUseCase.DeleteRoadmap(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "roadmap can not delete")
	}
	return nil
}
