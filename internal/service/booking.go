package service

import (
	"errors"
	"meeting-room-booking/internal/domain"
)

var (
	ErrOverlap = errors.New("the booking time overlaps with an existing reservation")
	ErrOrder   = errors.New("invalid booking time: start time must be before end time")
)

type BookingRepository interface {
	GetAll() ([]domain.Booking, error)
	GetByID(id int) (*domain.Booking, error)
	Create(booking domain.Booking) error
	Update(booking domain.Booking) error
	Delete(id int) error
}

type BookingService struct {
	bookingRepo BookingRepository
}

func NewBookingService(bookingRepo BookingRepository) *BookingService {
	return &BookingService{
		bookingRepo: bookingRepo,
	}
}

func (u *BookingService) GetAll() ([]domain.Booking, error) {
	return u.bookingRepo.GetAll()
}

func (u *BookingService) GetBookingByID(id int) (*domain.Booking, error) {
	return u.bookingRepo.GetByID(id)
}

func (u *BookingService) CreateBooking(booking domain.Booking) error {
	if booking.StartTime.After(booking.EndTime) || booking.StartTime.Equal(booking.EndTime) {
		return ErrOrder
	}

	bookings, err := u.bookingRepo.GetAll()
	if err != nil {
		return err
	}

	if err := IsBookingOverlapping(bookings, booking); err != nil {
		return err
	}

	return u.bookingRepo.Create(booking)
}

func (u *BookingService) UpdateBooking(booking domain.Booking) error {
	return u.bookingRepo.Update(booking)
}

func (u *BookingService) DeleteBooking(id int) error {
	return u.bookingRepo.Delete(id)
}

func IsBookingOverlapping(existingBookings []domain.Booking, newBooking domain.Booking) error {
	for _, booking := range existingBookings {
		if newBooking.RoomName == booking.RoomName &&
			newBooking.StartTime.Before(booking.EndTime) &&
			newBooking.EndTime.After(booking.StartTime) {
			return ErrOverlap
		}
	}
	return nil
}
