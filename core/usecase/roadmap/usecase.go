package roadmap

import (
	"wellbi-api/core/domain/model"
	"wellbi-api/core/domain/repository/roadmap"

	"github.com/labstack/echo/v4"
)

type RoadmapUseCase interface {
	GetRoadmaps(c echo.Context) ([]model.Roadmap, error)
	GetRoadmap(c echo.Context) (*model.Roadmap, error)
	CreateRoadmap(c echo.Context) (*model.Roadmap, error)
	UpdateRoadmap(c echo.Context) error
	DeleteRoadmap(c echo.Context) error
}

type roadmapUseCase struct {
	roadmap.RoadmapRepository
}

func NewRoadmapUseCase(r roadmap.RoadmapRepository) RoadmapUseCase {
	return &roadmapUseCase{r}
}

func (u *roadmapUseCase) GetRoadmaps(c echo.Context) ([]model.Roadmap, error) {
	return u.RoadmapRepository.Fetch(c)
}

func (u *roadmapUseCase) GetRoadmap(c echo.Context) (*model.Roadmap, error) {
	return u.RoadmapRepository.FetchByID(c)
}

func (u *roadmapUseCase) CreateRoadmap(c echo.Context) (*model.Roadmap, error) {
	return u.RoadmapRepository.Create(c)
}

func (u *roadmapUseCase) UpdateRoadmap(c echo.Context) error {
	return u.RoadmapRepository.Update(c)
}

func (u *roadmapUseCase) DeleteRoadmap(c echo.Context) error {
	return u.RoadmapRepository.Delete(c)
}
