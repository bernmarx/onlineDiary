package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/labstack/echo/v5"
)

func Query[T any, Y any](f AdapterFunc[T, Y]) echo.HandlerFunc {
	return func(c echo.Context) error {
		q, err := decodeQuery[T](c.Request())
		if err != nil {
			return err
		}

		res, err := f(c, q)
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

func decodeQuery[T any](r *http.Request) (T, error) {
	var res T

	err := schema.NewDecoder().Decode(&res, r.URL.Query())

	return res, err
}
