package controllers

import (
	"layerarch/delivery/common"
	"layerarch/entities"
	"layerarch/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	Repo repository.Book
}

func NewBookController(book repository.Book) *BookController {
	return &BookController{Repo: book}
}

func (bc BookController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
	
		books, err := bc.Repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(books))
	}
}

func (bc BookController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		rId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusBadRequest))
		}

		book, err := bc.Repo.Get(rId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(book))
	}
}

func (bc BookController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		b := entities.Book{}
		c.Bind(&b)

		book, err := bc.Repo.Insert(b)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK,  common.SuccessResponse(book))
	}
}

func (bc BookController) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		rId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusBadRequest))
		}

		b := entities.Book{}
		c.Bind(&b)

		book, err := bc.Repo.Edit(b, rId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(book))
	}
}

func (bc BookController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		rId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error(), http.StatusBadRequest))
		}

		book, err := bc.Repo.Delete(rId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.ErrorResponse(err.Error(), http.StatusNotFound))
		}

		return c.JSON(http.StatusOK, common.SuccessResponse(book))
	}
}