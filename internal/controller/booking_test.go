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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	return router
}

func TestGetAllBookings(t *testing.T) {
	mockService := new(mocks.BookingService)
	mockService.On("GetAll").Return([]domain.Booking{}, nil)

	bc := controller.NewBookingController(mockService)
	router := setupRouter()
	router.GET("/bookings", bc.GetAll)

	req, _ := http.NewRequest("GET", "/bookings", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBookingByID(t *testing.T) {
	mockService := new(mocks.BookingService)
	mockBooking := &domain.Booking{ID: 1, RoomName: "Room 1"}
	mockService.On("GetBookingByID", 1).Return(mockBooking, nil)

	bc := controller.NewBookingController(mockService)
	router := setupRouter()
	router.GET("/booking/:id", bc.GetByID)

	req, _ := http.NewRequest("GET", "/booking/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateBooking(t *testing.T) {
	mockService := new(mocks.BookingService)
	mockService.On("CreateBooking", mock.Anything).Return(nil)

	bc := controller.NewBookingController(mockService)
	router := setupRouter()
	router.POST("/booking", bc.Create)

	booking := &domain.Booking{RoomName: "Room 1"}
	jsonData, _ := json.Marshal(booking)
	req, _ := http.NewRequest("POST", "/booking", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateBooking(t *testing.T) {
	mockService := new(mocks.BookingService)
	mockService.On("UpdateBooking", mock.Anything).Return(nil)

	bc := controller.NewBookingController(mockService)
	router := setupRouter()
	router.PUT("/booking/:id", bc.Update)

	booking := &domain.Booking{RoomName: "Updated Room"}
	jsonData, _ := json.Marshal(booking)
	req, _ := http.NewRequest("PUT", "/booking/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteBooking(t *testing.T) {
	mockService := new(mocks.BookingService)
	mockService.On("DeleteBooking", 1).Return(nil)

	bc := controller.NewBookingController(mockService)
	router := setupRouter()
	router.DELETE("/booking/:id", bc.Delete)

	req, _ := http.NewRequest("DELETE", "/booking/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
