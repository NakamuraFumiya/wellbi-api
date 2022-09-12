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
	dsn := "twitter_user:password@tcp(127.0.0.1:3306)/twitter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&model.Post{})

	// // Create
	// db.Create(&model.Post{From: "Fumiya Nakamura", Message: "Hello!"})

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is a Twitter clone!")
	})

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
