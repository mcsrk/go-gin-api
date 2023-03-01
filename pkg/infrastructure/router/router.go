package router

import (
	"github.com/go-gin-api/pkg/adapter/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/albums", func(context echo.Context) error { return c.Album.GetAlbums(context) })

	return e
}
