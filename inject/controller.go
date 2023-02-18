package inject

import (
	postusecase "echo-twitter-clone/core/usecase/post"
	controller2 "echo-twitter-clone/presentation/api/controller"
	"echo-twitter-clone/presentation/api/controller/v1/post"
)

type appController struct {
	post.PostController
}

func (i *Injector) NewAppController() controller2.AppController {
	appController := &appController{}
	appController.PostController = i.NewPostController()
	return appController
}

func (i *Injector) NewPostController() post.PostController {
	return post.NewPostController(
		postusecase.NewPostUseCase(i.NewPostRepository()),
	)
}
