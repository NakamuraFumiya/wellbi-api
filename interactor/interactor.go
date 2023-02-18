package interactor

import (
	postrepository "echo-twitter-clone/core/domain/repository/post"
	postusecase "echo-twitter-clone/core/usecase/post"
	"echo-twitter-clone/infrastructure/persistence/gorm/handler"
	"echo-twitter-clone/infrastructure/persistence/gorm/repository/post"

	"echo-twitter-clone/presentation/controller"

	"gorm.io/gorm"
)

type Interactor interface {
	NewAppHandler() controller.AppHandler
	NewPostRepository() postrepository.PostRepository
	NewPostUseCase() postusecase.PostUseCase
	NewPostHandler() controller.PostHandler
}

type interactor struct {
	Conn *gorm.DB
}

func NewInteractor(Conn *gorm.DB) Interactor {
	return &interactor{Conn}
}

type appHandler struct {
	controller.PostHandler
}

func (i *interactor) NewAppHandler() controller.AppHandler {
	appHandler := &appHandler{}
	appHandler.PostHandler = i.NewPostHandler()
	return appHandler
}

func (i *interactor) NewPostRepository() postrepository.PostRepository {
	return post.NewPostRepository(&handler.Handler{})
}

func (i *interactor) NewPostUseCase() postusecase.PostUseCase {
	return postusecase.NewPostUseCase(i.NewPostRepository())
}

func (i *interactor) NewPostHandler() controller.PostHandler {
	return controller.NewPostHandler(i.NewPostUseCase())
}
