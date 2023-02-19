package post

import (
	"echo-twitter-clone/core/domain/model"
	"echo-twitter-clone/core/domain/repository/post"
	"echo-twitter-clone/infrastructure/persistence/gorm/handler"
	"errors"
	"github.com/labstack/echo/v4"
)

type PostRepository struct {
	handler *handler.Handler
}

func NewPostRepository(handler *handler.Handler) post.PostRepository {
	return &PostRepository{handler}
}

func (r *PostRepository) Fetch(c echo.Context) ([]model.Post, error) {
	var posts []model.Post

	if err := r.handler.DB().Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) FetchByID(c echo.Context) (*model.Post, error) {
	var post model.Post

	if err := r.handler.DB().First(&post, c.Param("id")).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) Create(c echo.Context) (*model.Post, error) {
	message := c.FormValue("message")
	if message == "" {
		return nil, errors.New("invalid to or message fields")
	}
	post := model.Post{
		Message: message,
	}
	if err := r.handler.DB().Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(c echo.Context) error {
	if err := r.handler.DB().Model(&model.Post{}).
		Where("id = ?", c.Param("id")).
		Update("message", c.FormValue("message")).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Delete(c echo.Context) error {
	if err := r.handler.DB().Delete(&model.Post{}, c.Param("id")).Error; err != nil {
		return err
	}
	return nil
}
