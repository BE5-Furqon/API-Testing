package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"layerarch/delivery/common"
	"layerarch/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockBookRepository struct{}

func (m mockBookRepository) GetAll() ([]entities.Book, error) {
	return []entities.Book{
		{
			Title: "Clean Code",
			Author: "Martin",
		},
	}, nil
}

func (m mockBookRepository) Get(bookId int) (entities.Book, error) {
	return entities.Book{
			Title: "Clean Code",
			Author: "Martin",
		}, nil
}

func (m mockBookRepository) Insert(entities.Book) (entities.Book, error) {
	return entities.Book{
			Title: "Clean Code",
			Author: "Martin",
	}, nil
}


func (m mockBookRepository) Edit(book entities.Book, bookId int) (entities.Book, error) {
	return entities.Book{
			Title: "Clean Code",
			Author: "Martin",
	}, nil
}

func (m mockBookRepository) Delete(bookId int) (entities.Book, error) {
	return entities.Book{
			Title: "Clean Code",
			Author: "Martin",
	}, nil
}

type mockFalseBookRepository struct{}

func (m mockFalseBookRepository) GetAll() ([]entities.Book, error) {
	return []entities.Book{
		{
			Title: "",
			Author: "",
		},
	}, errors.New("not found")
}

func (m mockFalseBookRepository) Get(bookId int) (entities.Book, error) {
	return entities.Book{
			Title: "",
			Author: "",
		}, errors.New("not found")
}

func (m mockFalseBookRepository) Insert(entities.Book) (entities.Book, error) {
	return entities.Book{
			Title: "",
			Author: "",
	}, errors.New("not found")
}

func (m mockFalseBookRepository) Edit(book entities.Book, bookId int) (entities.Book, error) {
	return entities.Book{
			Title: "",
			Author: "",
	}, errors.New("not found")
}

func (m mockFalseBookRepository) Delete(bookId int) (entities.Book, error) {
	return entities.Book{
		Title: "",
		Author: "",
	}, errors.New("not found")
}


func TestGetAllBookController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/books")

		bookController := NewBookController(mockBookRepository{})
		bookController.GetAll()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()
	
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/books")

		bookController := NewBookController(mockFalseBookRepository{})
		bookController.GetAll()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}

func TestGetBookController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		bookController := NewBookController(mockBookRepository{})
		bookController.Get()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
		assert.Equal(t, "Clean Code", response.Data.(interface{}).(map[string]interface{})["title"])
		assert.Equal(t, "Martin", response.Data.(interface{}).(map[string]interface{})["author"])
	})

	t.Run("Fail", func(t *testing.T) {
		e := echo.New()
		
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewBookController(mockFalseBookRepository{})
		userController.Get()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}

func TestInsertBookController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"title": "Clean Code",
			"author": "Martin",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/books")

		userController := NewBookController(mockBookRepository{})
		userController.Insert()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
		assert.Equal(t, "Clean Code", response.Data.(interface{}).(map[string]interface{})["title"])
		assert.Equal(t, "Martin", response.Data.(interface{}).(map[string]interface{})["author"])
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"title": "Clean Code",
			"author": "Martin",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/register")

		userController := NewBookController(mockFalseBookRepository{})
		userController.Insert()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}

func TestEditBookController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"title": "Clean Code",
			"author": "Martin",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewBookController(mockBookRepository{})
		userController.Edit()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"title": "Clean Code",
			"author": "Martin",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewBookController(mockFalseBookRepository{})
		userController.Edit()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}

func TestDeleteBookController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewBookController(mockBookRepository{})
		userController.Delete()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewBookController(mockFalseBookRepository{})
		userController.Delete()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}