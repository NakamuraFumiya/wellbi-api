package inject

import (
	roadmaprepository "wellbi-api/core/domain/repository/roadmap"
	"wellbi-api/infrastructure/persistence/gorm/handler"
	"wellbi-api/infrastructure/persistence/gorm/repository/roadmap"
)

func (i *Injector) NewRoadmapRepository() roadmaprepository.RoadmapRepository {
	return roadmap.NewRoadmapRepository(&handler.Handler{})
}
