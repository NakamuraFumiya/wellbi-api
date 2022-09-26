package post

import "github.com/labstack/echo/v4"

type PostRepository interface {
	Fetch(c echo.Context) error
	FetchByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
