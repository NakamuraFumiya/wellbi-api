package post

import (
	"echo-twitter-clone/config"
	"echo-twitter-clone/core/domain/model"
	"echo-twitter-clone/core/domain/repository/post"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(Conn *gorm.DB) post.PostRepository {
	return &PostRepository{Conn}
}

func (r *PostRepository) Fetch(c echo.Context) error {
	var posts []model.Post
	db := config.ConnectDB()
	db.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func (r *PostRepository) FetchByID(c echo.Context) error {
	var post model.Post
	db := config.ConnectDB()
	db.First(&post, c.Param("id"))
	if post.ID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "post does not exist"}
	}
	return c.JSON(http.StatusOK, post)
}

func (r *PostRepository) Create(c echo.Context) error {
	db := config.ConnectDB()
	message := c.FormValue("message")
	if message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid to or message fields"}
	}
	post := model.Post{
		Message: message,
	}
	db.Create(&post)
	return c.JSON(http.StatusOK, post)
}

func (r *PostRepository) Update(c echo.Context) error {
	db := config.ConnectDB()
	db.Model(&model.Post{}).
		Where("id = ?", c.Param("id")).
		Update("message", c.FormValue("message"))
	return c.String(http.StatusOK, "updated post id: "+c.Param("id")+"\n")
}

func (r *PostRepository) Delete(c echo.Context) error {
	db := config.ConnectDB()
	db.Delete(&model.Post{}, c.Param("id"))
	return c.String(http.StatusOK, "deleted post id: "+c.Param("id")+"\n")
}
