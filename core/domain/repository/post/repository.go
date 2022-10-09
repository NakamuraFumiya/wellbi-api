package post

import (
	"echo-twitter-clone/core/domain/model"

	"github.com/labstack/echo/v4"
)

type PostRepository interface {
	Fetch(c echo.Context) ([]model.Post, error)
	FetchByID(c echo.Context) (*model.Post, error)
	Create(c echo.Context) (*model.Post, error)
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
