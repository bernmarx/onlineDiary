package adapter

import "github.com/labstack/echo/v5"

type AdapterFunc[T any, Y any] func(echo.Context, T) (Y, error)
