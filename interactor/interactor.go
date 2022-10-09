package interactor

import (
	postusecase "echo-twitter-clone/core/application/usecase/post"
	postrepository "echo-twitter-clone/core/domain/repository/post"
	"echo-twitter-clone/infrastructure/persistence/gorm/repository/post"
	handler "echo-twitter-clone/presentation/handler"

	"gorm.io/gorm"
)

type Interactor interface {
	NewAppHandler() handler.AppHandler
	NewPostRepository() postrepository.PostRepository
	NewPostUseCase() postusecase.PostUseCase
	NewPostHandler() handler.PostHandler
}

type interactor struct {
	Conn *gorm.DB
}

func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	handler.PostHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.PostHandler = i.NewPostHandler()
	return appHandler
}

func (i *interactor) NewPostRepository() postrepository.PostRepository {
	return post.NewPostRepository(i.Conn)
}

func (i *interactor) NewPostUseCase() postusecase.PostUseCase {
	return postusecase.NewPostUseCase(i.NewPostRepository())
}

func (i *interactor) NewPostHandler() handler.PostHandler {
	return handler.NewPostHandler(i.NewPostUseCase())
}
