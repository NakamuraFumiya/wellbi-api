package inject

import (
	roadmaprepository "echo-twitter-clone/core/domain/repository/roadmap"
	"echo-twitter-clone/infrastructure/persistence/gorm/handler"
	"echo-twitter-clone/infrastructure/persistence/gorm/repository/roadmap"
)

func (i *Injector) NewRoadmapRepository() roadmaprepository.RoadmapRepository {
	return roadmap.NewRoadmapRepository(&handler.Handler{})
}
