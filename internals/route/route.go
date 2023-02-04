package route

import (
	"github.com/labstack/echo/v4"
	"shortener/internals/core/application/handlers"
)

func SetRoute(e *echo.Echo) error {
	URLRoute(e)
	RedirectRoute(e)
	return nil
}

func URLRoute(e *echo.Echo) {
	urlhandler := handlers.UrlHandler()

	g := e.Group("urls")
	g.GET("", urlhandler.Index)
	g.POST("", urlhandler.Store)
	g.GET("/:id", urlhandler.Show)
}

func RedirectRoute(e *echo.Echo) {
	urlhandler := handlers.UrlHandler()

	g := e.Group("s")
	g.GET("/:short", urlhandler.Redirect)
}
