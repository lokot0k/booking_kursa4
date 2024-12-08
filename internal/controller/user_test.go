package controller_test

import (
	"bytes"
	"encoding/json"
	"meeting-room-booking/internal/controller"
	"meeting-room-booking/internal/controller/mocks"
	"meeting-room-booking/internal/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsers(t *testing.T) {
	mockService := new(mocks.UserService)
	mockService.On("GetAllUsers").Return([]domain.User{}, nil)

	uc := controller.NewUserController(mockService)
	router := setupRouter()
	router.GET("/users", uc.GetAll)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterUser(t *testing.T) {
	mockService := new(mocks.UserService)
	mockService.On("CreateUser", mock.Anything).Return(nil)

	uc := controller.NewUserController(mockService)
	router := setupRouter()
	router.POST("/register", uc.Register)

	user := &domain.User{Name: "test", Password: "password"}
	jsonData, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginUser(t *testing.T) {
	mockService := new(mocks.UserService)
	mockUser := &domain.User{ID: 1, Name: "test"}
	mockService.On("Authenticate", "test", "password").Return(mockUser, nil)

	uc := controller.NewUserController(mockService)
	router := setupRouter()
	router.POST("/login", uc.Login)

	req, _ := http.NewRequest("POST", "/login", nil)
	req.Header.Set("Authorization", "Basic test:password")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
