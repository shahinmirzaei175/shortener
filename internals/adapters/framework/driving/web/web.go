package web

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"shortener/internals/config"
	"shortener/internals/ports"
	"shortener/internals/route"
	Utility "shortener/internals/utilities"
)

type webAdapter struct {
}

func (w webAdapter) Run() {
	server := echo.New()
	server.Validator = &Utility.CustomValidator{Validator: validator.New()}
	err := route.SetRoute(server)
	if err != nil {
		panic(err)
	}
	err = server.Start(":" + config.AppConfig.ServerPort)
	if err != nil {
		panic(err)
	}
}

func WebAdapter() ports.Web {
	return webAdapter{}
}
