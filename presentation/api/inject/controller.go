package inject

import (
	roadmapusecase "echo-twitter-clone/core/usecase/roadmap"
	v1roadmapcreate "echo-twitter-clone/presentation/api/controller/v1/roadmap/create"
	v1roadmapdelete "echo-twitter-clone/presentation/api/controller/v1/roadmap/delete"
	v1roadmapdetail "echo-twitter-clone/presentation/api/controller/v1/roadmap/detail"
	v1roadmaplist "echo-twitter-clone/presentation/api/controller/v1/roadmap/list"
	v1roadmapupdate "echo-twitter-clone/presentation/api/controller/v1/roadmap/update"
	"github.com/labstack/echo/v4"
)

// V1Controllerへの依存注入
func (i *Injector) V1RoadmapListController() echo.HandlerFunc {
	return NewHandlerFunc(v1roadmaplist.NewController(
		roadmapusecase.NewRoadmapUseCase(i.NewRoadmapRepository()),
	).GetRoadmaps)
}

func (i *Injector) V1RoadmapDetailController() echo.HandlerFunc {
	return NewHandlerFunc(v1roadmapdetail.NewController(
		roadmapusecase.NewRoadmapUseCase(i.NewRoadmapRepository()),
	).GetRoadmap)
}

func (i *Injector) V1RoadmapCreateController() echo.HandlerFunc {
	return NewHandlerFunc(v1roadmapcreate.NewController(
		roadmapusecase.NewRoadmapUseCase(i.NewRoadmapRepository()),
	).CreateRoadmap)
}

func (i *Injector) V1RoadmapUpdateController() echo.HandlerFunc {
	return NewHandlerFunc(v1roadmapupdate.NewController(
		roadmapusecase.NewRoadmapUseCase(i.NewRoadmapRepository()),
	).UpdateRoadmap)
}

func (i *Injector) V1RoadmapDeleteController() echo.HandlerFunc {
	return NewHandlerFunc(v1roadmapdelete.NewController(
		roadmapusecase.NewRoadmapUseCase(i.NewRoadmapRepository()),
	).DeleteRoadmap)
}
