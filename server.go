package main

import (
	"echo-twitter-clone/model"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize DB
	db := dbInit()
	db.AutoMigrate(&model.Post{})

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is a Twitter clone!")
	})

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}

// handler
func dbInit() *gorm.DB {
	dsn := "twitter_user:password@tcp(127.0.0.1:3306)/twitter?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
