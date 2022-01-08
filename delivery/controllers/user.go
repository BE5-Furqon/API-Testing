package controllers

import (
	"layerarch/delivery/common"
	"layerarch/entities"
	"layerarch/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo repository.User
}

func NewUserController(user repository.User) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
	
		users, err := uc.Repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(users))
	}
}

func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		rId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusBadRequest))
		}

		user, err := uc.Repo.Get(rId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(user))
	}
}

func (uc UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := entities.User{}
		c.Bind(&u)

		user, err := uc.Repo.Register(u)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(user))
	}
}

func (uc UserController) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		rId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusBadRequest))
		}

		u := entities.User{}
		c.Bind(&u)

		user, err := uc.Repo.Edit(u, rId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(user))
	}
}

func (uc UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		rId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusBadRequest))
		}

		user, err := uc.Repo.Delete(rId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(user))
	}
}

func (uc UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var u entities.User
		c.Bind(&u)

		user, err := uc.Repo.Login(u.Email, u.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusBadRequest))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(user))
	}
}