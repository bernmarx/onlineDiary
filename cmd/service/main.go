package main

import (
	"log"
	"net/http"

	"github.com/bernmarx/onlineDiary/internal/handlers/api_1_get_children_list_for_parent"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// alarmedruby@gmail.com
// Dh54Fuc7MPb6g9d

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/api/1/getChildrenListForParent",
			Handler: api_1_get_children_list_for_parent.NewHandler(app.Dao()),
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
