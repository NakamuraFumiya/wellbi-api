package main

import (
	"wellbi-api/core/domain/model"
	"wellbi-api/infrastructure/persistence/gorm/handler"
	"wellbi-api/presentation/api/inject"
	middlewarepackage "wellbi-api/presentation/middleware"
	"wellbi-api/presentation/router"

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
