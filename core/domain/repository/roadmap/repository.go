package roadmap

import (
	"echo-twitter-clone/core/domain/model"

	"github.com/labstack/echo/v4"
)

type RoadmapRepository interface {
	Fetch(c echo.Context) ([]model.Roadmap, error)
	FetchByID(c echo.Context) (*model.Roadmap, error)
	Create(c echo.Context) (*model.Roadmap, error)
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
