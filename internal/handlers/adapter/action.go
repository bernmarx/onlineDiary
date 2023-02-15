package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v5"
)

func Action[T any, Y any](f AdapterFunc[T, Y]) echo.HandlerFunc {
	return func(c echo.Context) error {
		v, err := decode[T](c.Request())
		if err != nil {
			return err
		}

		res, err := f(c, v)
		if err != nil {
			return err
		}

		bytes, err := json.Marshal(res)
		if err != nil {
			return err
		}

		c.Response().Writer.WriteHeader(http.StatusOK)
		_, err = c.Response().Writer.Write(bytes)

		return err
	}
}

func decode[T any](r *http.Request) (T, error) {
	var res T

	err := json.NewDecoder(r.Body).Decode(&res)

	return res, err
}
