package main

import (
	"echo-twitter-clone/model"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
	e.PUT("/posts/:id", updatePost)
	e.DELETE("posts/:id", deletePost)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

func connectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("APP_USER"), os.Getenv("APP_PASSWORD"), os.Getenv("APP_DB_PORT"), os.Getenv("APP_DATABASE"))
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
	post := model.Post{
		Message: message,
	}
	db.Create(&post)
	return c.JSON(http.StatusOK, post)
}

func readPostAll(c echo.Context) error {
	var posts []model.Post
	db := connectDB()
	db.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func readPostDetail(c echo.Context) error {
	var post model.Post
	db := connectDB()
	db.First(&post, c.Param("id"))
	if post.ID == 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "post does not exist"}
	}
	return c.JSON(http.StatusOK, post)
}

func updatePost(c echo.Context) error {
	db := connectDB()
	db.Model(&model.Post{}).
		Where("id = ?", c.Param("id")).
		Update("message", c.FormValue("message"))
	return c.String(http.StatusOK, "updated post id: "+c.Param("id")+"\n")
}

func deletePost(c echo.Context) error {
	db := connectDB()
	db.Delete(&model.Post{}, c.Param("id"))
	return c.String(http.StatusOK, "deleted post id: "+c.Param("id")+"\n")
}
