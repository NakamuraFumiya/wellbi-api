package inject

import (
	postusecase "echo-twitter-clone/core/usecase/post"
	v1postcreate "echo-twitter-clone/presentation/api/controller/v1/post/create"
	v1postdelete "echo-twitter-clone/presentation/api/controller/v1/post/delete"
	v1postdetail "echo-twitter-clone/presentation/api/controller/v1/post/detail"
	v1postlist "echo-twitter-clone/presentation/api/controller/v1/post/list"
	v1postupdate "echo-twitter-clone/presentation/api/controller/v1/post/update"
	"github.com/labstack/echo/v4"
)

// V1Controllerへの依存注入
func (i *Injector) V1PostListController() echo.HandlerFunc {
	return NewHandlerFunc(v1postlist.NewController(
		postusecase.NewPostUseCase(i.NewPostRepository()),
	).GetPosts)
}

func (i *Injector) V1PostDetailController() echo.HandlerFunc {
	return NewHandlerFunc(v1postdetail.NewController(
		postusecase.NewPostUseCase(i.NewPostRepository()),
	).GetPost)
}

func (i *Injector) V1PostCreateController() echo.HandlerFunc {
	return NewHandlerFunc(v1postcreate.NewController(
		postusecase.NewPostUseCase(i.NewPostRepository()),
	).CreatePost)
}

func (i *Injector) V1PostUpdateController() echo.HandlerFunc {
	return NewHandlerFunc(v1postupdate.NewController(
		postusecase.NewPostUseCase(i.NewPostRepository()),
	).UpdatePost)
}

func (i *Injector) V1PostDeleteController() echo.HandlerFunc {
	return NewHandlerFunc(v1postdelete.NewController(
		postusecase.NewPostUseCase(i.NewPostRepository()),
	).DeletePost)
}
