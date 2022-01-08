package middlewares

import (
	"layerarch/constant"
	"layerarch/entities"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JwtMiddleware(e *echo.Group) {
	config := middleware.JWTConfig{
		Claims:     &entities.User{},
		SigningKey: []byte(constant.JWT_SECRET),
	}

	e.Use(middleware.JWTWithConfig(config))
}