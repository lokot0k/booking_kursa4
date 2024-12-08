package service_test

import (
	"testing"
	"time"

	"meeting-room-booking/internal/domain"
	"meeting-room-booking/internal/service"
	"meeting-room-booking/internal/service/mocks"

	"github.com/stretchr/testify/require"
)

func TestBookingService_CreateBooking_Success(t *testing.T) {
	mockRepo := new(mocks.BookingRepository)

	newBooking := domain.Booking{
		RoomName:  "Room1",
		StartTime: time.Date(2024, 11, 22, 10, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 11, 22, 12, 0, 0, 0, time.UTC),
	}
	existingBookings := []domain.Booking{}

	mockRepo.On("GetAll").Return(existingBookings, nil)
	mockRepo.On("Create", newBooking).Return(nil)

	svc := service.NewBookingService(mockRepo)

	err := svc.CreateBooking(newBooking)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestBookingService_CreateBooking_OverlapError(t *testing.T) {
	mockRepo := new(mocks.BookingRepository)

	newBooking := domain.Booking{
		RoomName:  "Room1",
		StartTime: time.Date(2024, 11, 22, 11, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 11, 22, 13, 0, 0, 0, time.UTC),
	}
	existingBookings := []domain.Booking{
		{
			RoomName:  "Room1",
			StartTime: time.Date(2024, 11, 22, 10, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 11, 22, 12, 0, 0, 0, time.UTC),
		},
	}

	mockRepo.On("GetAll").Return(existingBookings, nil)

	svc := service.NewBookingService(mockRepo)

	err := svc.CreateBooking(newBooking)
	require.Error(t, err)
	require.Equal(t, service.ErrOverlap, err)

	mockRepo.AssertExpectations(t)
}

func TestBookingService_CreateBooking_OrderError(t *testing.T) {
	mockRepo := new(mocks.BookingRepository)

	newBooking := domain.Booking{
		RoomName:  "Room1",
		StartTime: time.Date(2024, 11, 22, 14, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 11, 22, 12, 0, 0, 0, time.UTC),
	}

	svc := service.NewBookingService(mockRepo)

	err := svc.CreateBooking(newBooking)
	require.Error(t, err)
	require.Equal(t, service.ErrOrder, err)

	mockRepo.AssertNotCalled(t, "GetAll")
	mockRepo.AssertNotCalled(t, "Create")
}

func TestBookingService_GetAll_Success(t *testing.T) {
	mockRepo := new(mocks.BookingRepository)

	existingBookings := []domain.Booking{
		{
			ID:        1,
			RoomName:  "Room1",
			StartTime: time.Date(2024, 11, 22, 10, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 11, 22, 12, 0, 0, 0, time.UTC),
		},
	}

	mockRepo.On("GetAll").Return(existingBookings, nil)

	svc := service.NewBookingService(mockRepo)

	bookings, err := svc.GetAll()
	require.NoError(t, err)
	require.Equal(t, existingBookings, bookings)

	mockRepo.AssertExpectations(t)
}

func TestBookingService_DeleteBooking_Success(t *testing.T) {
	mockRepo := new(mocks.BookingRepository)

	mockRepo.On("Delete", 1).Return(nil)

	svc := service.NewBookingService(mockRepo)

	err := svc.DeleteBooking(1)
	require.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
