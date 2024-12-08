package controller

import (
	"meeting-room-booking/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingService interface {
	GetAll() ([]domain.Booking, error)
	GetBookingByID(id int) (*domain.Booking, error)
	CreateBooking(booking domain.Booking) error
	UpdateBooking(booking domain.Booking) error
	DeleteBooking(id int) error
}

type BookingController struct {
	bookingService BookingService
}

func NewBookingController(bookingService BookingService) *BookingController {
	return &BookingController{
		bookingService: bookingService,
	}
}

func (bc *BookingController) GetAll(c *gin.Context) {
	_, err := bc.bookingService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	bookings, err := bc.bookingService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookings,
	})
}

func (ctrl *BookingController) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	booking, err := ctrl.bookingService.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func (ctrl *BookingController) Create(c *gin.Context) {
	var booking domain.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.bookingService.CreateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Booking created successfully"})
}

func (ctrl *BookingController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var booking domain.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	booking.ID = id
	if err := ctrl.bookingService.UpdateBooking(booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully"})
}

func (ctrl *BookingController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.bookingService.DeleteBooking(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
