package main

import (
	"echo-twitter-clone/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connecting to a Database
	db := connectDB()
	// Migrate the schema
	db.AutoMigrate(&model.Post{})

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is a Twitter clone!")
	})
	e.POST("/posts", createPost)
	e.GET("/posts", readPostAll)
	e.GET("/posts/:id", readPostDetail)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

func connectDB() *gorm.DB {
	dsn := "twitter_user:password@tcp(127.0.0.1:3306)/twitter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func createPost(c echo.Context) error {
	db := connectDB()
	message := c.FormValue("message")
	if message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid to or message fields"}
	}
	db.Create(&model.Post{Message: message})
	return nil
}

func readPostAll(c echo.Context) error {
	var posts []model.Post
	db := connectDB()
	db.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func readPostDetail(c echo.Context) error {
	var post model.Post
	id := c.Param("id")
	db := connectDB()
	db.First(&post, id)
	if post.ID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "post does not exist"}
	}
	return c.JSON(http.StatusOK, post)
}
