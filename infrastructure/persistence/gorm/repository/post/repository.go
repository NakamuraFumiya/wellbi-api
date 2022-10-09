package post

import (
	"echo-twitter-clone/config"
	"echo-twitter-clone/core/domain/model"
	"echo-twitter-clone/core/domain/repository/post"
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(Conn *gorm.DB) post.PostRepository {
	return &PostRepository{Conn}
}

func (r *PostRepository) Fetch(c echo.Context) ([]model.Post, error) {
	var posts []model.Post
	db := config.ConnectDB()

	if err := db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) FetchByID(c echo.Context) (*model.Post, error) {
	var post model.Post
	db := config.ConnectDB()

	if err := db.First(&post, c.Param("id")).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) Create(c echo.Context) (*model.Post, error) {
	db := config.ConnectDB()
	message := c.FormValue("message")
	if message == "" {
		return nil, errors.New("invalid to or message fields")
	}
	post := model.Post{
		Message: message,
	}
	if err := db.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Update(c echo.Context) error {
	db := config.ConnectDB()
	if err := db.Model(&model.Post{}).
		Where("id = ?", c.Param("id")).
		Update("message", c.FormValue("message")).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Delete(c echo.Context) error {
	db := config.ConnectDB()
	if err := db.Delete(&model.Post{}, c.Param("id")).Error; err != nil {
		return err
	}
	return nil
}
