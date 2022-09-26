package router

import (
	"echo-twitter-clone/config"
	"echo-twitter-clone/core/domain/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouter(e *echo.Echo) {
	// Routes
	e.GET("/", hello)
	e.POST("/posts", createPost)
	e.GET("/posts", readPostAll)
	e.GET("/posts/:id", readPostDetail)
	e.PUT("/posts/:id", updatePost)
	e.DELETE("posts/:id", deletePost)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is a Twitter clone!")
}

func createPost(c echo.Context) error {
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

func readPostAll(c echo.Context) error {
	var posts []model.Post
	db := config.ConnectDB()
	db.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func readPostDetail(c echo.Context) error {
	var post model.Post
	db := config.ConnectDB()
	db.First(&post, c.Param("id"))
	if post.ID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "post does not exist"}
	}
	return c.JSON(http.StatusOK, post)
}

func updatePost(c echo.Context) error {
	db := config.ConnectDB()
	db.Model(&model.Post{}).
		Where("id = ?", c.Param("id")).
		Update("message", c.FormValue("message"))
	return c.String(http.StatusOK, "updated post id: "+c.Param("id")+"\n")
}

func deletePost(c echo.Context) error {
	db := config.ConnectDB()
	db.Delete(&model.Post{}, c.Param("id"))
	return c.String(http.StatusOK, "deleted post id: "+c.Param("id")+"\n")
}
