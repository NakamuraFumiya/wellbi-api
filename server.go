package main

import (
	"echo-twitter-clone/config"
	"echo-twitter-clone/core/domain/model"
	"echo-twitter-clone/interactor"
	"echo-twitter-clone/presentation/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connecting to a Database
	db := config.ConnectDB()
	// Migrate the schema
	db.AutoMigrate(&model.Post{})

	// Routes
	i := interactor.NewInteractor(db)
	h := i.NewAppHandler()
	router.SetRouter(e, h)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
