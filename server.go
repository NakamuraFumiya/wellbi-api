package main

import (
	"echo-twitter-clone/core/domain/model"
	"echo-twitter-clone/infrastructure/persistence/gorm/handler"
	"echo-twitter-clone/presentation/api/inject"
	middlewarepackage "echo-twitter-clone/presentation/middleware"
	"echo-twitter-clone/presentation/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	middlewarepackage.NewMiddleware(e)
	e.Use(middleware.CORS())

	// Connecting to a Database
	db := handler.ConnectDB()
	// Migrate the schema
	db.AutoMigrate(&model.Roadmap{})

	// Routes
	i := inject.Injector{}
	router.SetRouter(e, i)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
