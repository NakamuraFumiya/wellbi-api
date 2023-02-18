package inject

import "github.com/labstack/echo/v4"

type Injector struct{}

type callFunc func(ctx echo.Context) error

// TODO: この辺りまだよく分かってないから調べる
func NewHandlerFunc(h callFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return h(ctx)
	}
}
