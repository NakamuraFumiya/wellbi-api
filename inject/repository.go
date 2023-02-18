package inject

import (
	postrepository "echo-twitter-clone/core/domain/repository/post"
	"echo-twitter-clone/infrastructure/persistence/gorm/handler"
	"echo-twitter-clone/infrastructure/persistence/gorm/repository/post"
)

func (i *Injector) NewPostRepository() postrepository.PostRepository {
	return post.NewPostRepository(&handler.Handler{})
}
