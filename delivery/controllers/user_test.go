package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"layerarch/constant"
	"layerarch/delivery/common"
	"layerarch/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct{}

func (m mockUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{
			Name:           "Furqon",
			Email:          "furqonzt99@gmail.com",
			Password:       "1234qwer",
		},
	}, nil
}

func (m mockUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{
			Name:           "Furqon",
			Email:          "furqonzt99@gmail.com",
			Password:       "1234qwer",
		}, nil
}

func (m mockUserRepository) Register(entities.User) (entities.User, error) {
	return entities.User{
		Name:           "Furqon",
		Email:          "furqonzt99@gmail.com",
		Password:       "1234qwer",
	}, nil
}

func (m mockUserRepository) Login(email string, password string) (interface{}, error) {
	return entities.User{
		Name:           "Furqon",
		Email:          "furqonzt99@gmail.com",
		Password:       "1234qwer",
	}, nil
}

func (m mockUserRepository) Edit(user entities.User, userId int) (entities.User, error) {
	return entities.User{
		Name:           "Furqon",
		Email:          "furqonzt99@gmail.com",
		Password:       "1234qwer",
	}, nil
}

func (m mockUserRepository) Delete(userId int) (entities.User, error) {
	return entities.User{
		Name:           "Furqon",
		Email:          "furqonzt99@gmail.com",
		Password:       "1234qwer",
	}, nil
}

type mockFalseUserRepository struct{}

func (m mockFalseUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{
			Name:           "",
			Email:          "",
			Password:       "",
		},
	}, errors.New("not found")
}

func (m mockFalseUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{
			Name:           "",
			Email:          "",
			Password:       "",
		}, errors.New("not found")
}

func (m mockFalseUserRepository) Register(entities.User) (entities.User, error) {
	return entities.User{
		Name:           "",
		Email:          "",
		Password:       "",
	}, errors.New("bad request")
}

func (m mockFalseUserRepository) Login(email string, password string) (interface{}, error) {
	return entities.User{
		Name:           "",
		Email:          "",
		Password:       "",
	}, errors.New("bad request")
}

func (m mockFalseUserRepository) Edit(user entities.User, userId int) (entities.User, error) {
	return entities.User{
		Name:           "",
		Email:          "",
		Password:       "",
	}, errors.New("not found")
}

func (m mockFalseUserRepository) Delete(userId int) (entities.User, error) {
	return entities.User{
		Name:           "",
		Email:          "",
		Password:       "",
	}, errors.New("not found")
}

var jwtToken string

func TestRegisterUserController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"name": "Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/register")

		userController := NewUserController(mockUserRepository{})
		userController.Register()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
		assert.Equal(t, "Furqon", response.Data.(interface{}).(map[string]interface{})["name"])
		assert.Equal(t, "furqonzt99@gmail.com", response.Data.(interface{}).(map[string]interface{})["email"])
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"name": "Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/register")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Register()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "bad request", response.Message)
	})
}

func TestLoginUserController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/login")

		userController := NewUserController(mockUserRepository{})
		userController.Login()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		jwtToken = fmt.Sprintf("%v", response.Data)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/login")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Login()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, "bad request", response.Message)
	})
}

func TestGetAllUserController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))
	
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/users")

		userController := NewUserController(mockUserRepository{})
		userController.GetAll()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))
	
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/users")

		userController := NewUserController(mockFalseUserRepository{})
		userController.GetAll()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}

func TestGetUserController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))
		
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewUserController(mockUserRepository{})
		userController.Get()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
		assert.Equal(t, "Furqon", response.Data.(interface{}).(map[string]interface{})["name"])
		assert.Equal(t, "furqonzt99@gmail.com", response.Data.(interface{}).(map[string]interface{})["email"])
	})

	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))
		
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Get()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}

func TestEditUserController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))

		requestBody, _ := json.Marshal(map[string]string{
			"name": "Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewUserController(mockUserRepository{})
		userController.Edit()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))

		requestBody, _ := json.Marshal(map[string]string{
			"name": "Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Edit()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}

func TestDeleteUserController(t *testing.T)  {

	t.Run("Success", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewUserController(mockUserRepository{})
		userController.Delete()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "success operation", response.Message)
	})
	
	t.Run("Fail", func(t *testing.T) {
		e := echo.New()

		config := middleware.JWTConfig{
			Claims:     &entities.User{},
			SigningKey: []byte(constant.JWT_SECRET),
		}

		e.Use(middleware.JWTWithConfig(config))

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("%v", jwtToken))
		res := httptest.NewRecorder()
		
		c := e.NewContext(req, res)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		userController := NewUserController(mockFalseUserRepository{})
		userController.Delete()(c)

		var response common.Response

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, "not found", response.Message)
	})
}