package main

import (
	"fmt"
	"layerarch/config"
	"layerarch/delivery/controllers"
	"layerarch/delivery/routes"
	"layerarch/repository"
	"layerarch/util"

	"github.com/labstack/echo/v4"
)

func main()  {
	config := config.GetConfig()

	db := util.InitDB(config)
	util.InitialMigrate(db)

	userRepo := repository.NewUserRepo(db)
	userController := controllers.NewUserController(userRepo)

	bookRepo := repository.NewBookRepo(db)
	bookController := controllers.NewBookController(bookRepo)

	e := echo.New()

	routes.RegisterUserPath(e, userController)
	routes.RegisterBookPath(e, bookController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Port)))
}