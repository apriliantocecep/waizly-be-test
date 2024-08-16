package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"waizly/internal/config"
	"waizly/internal/entity"
	"waizly/internal/model"
)

func TestRegister(t *testing.T) {
	ClearAll()

	requestBody := model.RegisterUserRequest{
		Name:     "Cecep Aprilianto",
		Email:    "cecepaprilianto@gmail.com",
		Password: "makanterus",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(bodyJson))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[model.UserResponse, any]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, requestBody.Name, responseBody.Data.User.Name)
	assert.Equal(t, requestBody.Email, responseBody.Data.User.Email)
	assert.NotNil(t, responseBody.Data.User.CreatedAt)
	assert.NotNil(t, responseBody.Data.User.UpdatedAt)
	assert.NotEmpty(t, responseBody.Data.Token)
}

func TestRegisterError(t *testing.T) {
	ClearAll()

	requestBody := model.RegisterUserRequest{
		Name:     "",
		Email:    "",
		Password: "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(bodyJson))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[any, []model.ErrorDetail]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, 3, len(responseBody.Details))
	assert.Equal(t, "The Name is required", responseBody.Details[0].Message)
	assert.Equal(t, "The Email is required", responseBody.Details[1].Message)
	assert.Equal(t, "The Password is required", responseBody.Details[2].Message)
}

func TestRegisterDuplicate(t *testing.T) {
	ClearAll()
	TestRegister(t) // register success

	requestBody := model.RegisterUserRequest{
		Name:     "Cecep Aprilianto",
		Email:    "cecepaprilianto@gmail.com",
		Password: "makanterus",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(bodyJson))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[any, []model.ErrorDetail]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Equal(t, 1, len(responseBody.Details))
	assert.Equal(t, "email already exists", responseBody.Details[0].Message)
}

func TestLogin(t *testing.T) {
	TestRegister(t) // register success

	requestBody := model.LoginUserRequest{
		Email:    "cecepaprilianto@gmail.com",
		Password: "makanterus",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(bodyJson))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[model.UserResponse, any]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, requestBody.Email, responseBody.Data.User.Email)
	assert.NotNil(t, responseBody.Data.User.CreatedAt)
	assert.NotNil(t, responseBody.Data.User.UpdatedAt)
	assert.NotEmpty(t, responseBody.Data.Token)
}

func TestLoginWrongEmail(t *testing.T) {
	ClearAll()
	TestRegister(t) // register success

	requestBody := model.LoginUserRequest{
		Email:    "salah",
		Password: "makanterus",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(bodyJson))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[any, []model.ErrorDetail]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, 1, len(responseBody.Details))
	assert.Equal(t, "The Email must be an email", responseBody.Details[0].Message)
}

func TestLoginWrongPassword(t *testing.T) {
	ClearAll()
	TestRegister(t) // register success

	requestBody := model.LoginUserRequest{
		Email:    "cecepaprilianto@gmail.com",
		Password: "salah",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request, err := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(bodyJson))
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[any, []model.ErrorDetail]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, 1, len(responseBody.Details))
	assert.Equal(t, "invalid password", responseBody.Details[0].Message)
}

func TestGetCurrentUser(t *testing.T) {
	ClearAll()
	TestLogin(t) // login success

	user := new(entity.User)
	err := db.Where("email = ?", "cecepaprilianto@gmail.com").First(&user).Error
	assert.Nil(t, err)

	jwt := config.NewJwtWrapper(viperConfig)
	token, err := jwt.GenerateToken(*user)
	assert.Nil(t, err)

	request, err := http.NewRequest(http.MethodGet, "/api/user/me", nil)
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[model.UserResponse, any]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, responseBody.Data.User.ID)
	assert.Equal(t, "Cecep Aprilianto", responseBody.Data.User.Name)
	assert.Equal(t, "cecepaprilianto@gmail.com", responseBody.Data.User.Email)
	assert.NotNil(t, responseBody.Data.User.CreatedAt)
	assert.NotNil(t, responseBody.Data.User.UpdatedAt)
}

func TestGetCurrentUserUnauthorized(t *testing.T) {
	ClearAll()
	TestLogin(t) // login success

	request, err := http.NewRequest(http.MethodGet, "/api/user/me", nil)
	assert.Nil(t, err)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "wrong")

	w := httptest.NewRecorder()
	app.ServeHTTP(w, request)

	var responseBody model.ApiResponse[any, string]
	errUnmarshal := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Nil(t, errUnmarshal)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "Unauthorized", responseBody.Details)
}
