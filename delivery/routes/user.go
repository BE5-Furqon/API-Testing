package routes

import (
	"layerarch/delivery/controllers"
	m "layerarch/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterUserPath(e *echo.Echo, uc *controllers.UserController) {

	m.LogMiddleware(e)

	r := e.Group("")
	m.JwtMiddleware(r)
	
	r.GET("/users", uc.GetAll())
	r.GET("/users/:id", uc.Get())
	e.POST("/users/register", uc.Register())
	e.POST("/users/login", uc.Login())
	r.PUT("/users/:id", uc.Edit())
	r.DELETE("/users/:id", uc.Delete())
}