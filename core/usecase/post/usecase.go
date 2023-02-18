package post

import (
	"echo-twitter-clone/core/domain/model"
	"echo-twitter-clone/core/domain/repository/post"

	"github.com/labstack/echo/v4"
)

type PostUseCase interface {
	GetPosts(c echo.Context) ([]model.Post, error)
	GetPost(c echo.Context) (*model.Post, error)
	CreatePost(c echo.Context) (*model.Post, error)
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}

type postUseCase struct {
	post.PostRepository
}

func NewPostUseCase(r post.PostRepository) PostUseCase {
	return &postUseCase{r}
}

func (u *postUseCase) GetPosts(c echo.Context) ([]model.Post, error) {
	return u.PostRepository.Fetch(c)
}

func (u *postUseCase) GetPost(c echo.Context) (*model.Post, error) {
	return u.PostRepository.FetchByID(c)
}

func (u *postUseCase) CreatePost(c echo.Context) (*model.Post, error) {
	return u.PostRepository.Create(c)
}

func (u *postUseCase) UpdatePost(c echo.Context) error {
	return u.PostRepository.Update(c)
}

func (u *postUseCase) DeletePost(c echo.Context) error {
	return u.PostRepository.Delete(c)
}
