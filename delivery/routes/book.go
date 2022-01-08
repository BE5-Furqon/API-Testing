package routes

import (
	"layerarch/delivery/controllers"
	m "layerarch/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterBookPath(e *echo.Echo, bc *controllers.BookController)  {
	m.LogMiddleware(e)

	r := e.Group("")
	m.JwtMiddleware(r)
	
	e.GET("/books", bc.GetAll())
	e.GET("/books/:id", bc.Get())
	r.POST("/books", bc.Insert())
	r.PUT("/books/:id", bc.Edit())
	r.DELETE("/books/:id", bc.Delete())
}