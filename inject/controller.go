package inject

import (
	postusecase "echo-twitter-clone/core/usecase/post"
	"echo-twitter-clone/presentation/controller"
)

type appController struct {
	controller.PostController
}

func (i *Injector) NewAppController() controller.AppController {
	appController := &appController{}
	appController.PostController = i.NewPostController()
	return appController
}

func (i *Injector) NewPostController() controller.PostController {
	return controller.NewPostController(
		postusecase.NewPostUseCase(i.NewPostRepository()),
	)
}
