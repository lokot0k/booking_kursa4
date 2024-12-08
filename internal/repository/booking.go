package repository

import (
	"database/sql"
	"meeting-room-booking/internal/domain"
)

type BookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) GetAll() ([]domain.Booking, error) {

	rows, err := r.db.Query("SELECT bookings.id, room_name, start_time, end_time, user_id, username FROM bookings JOIN users ON user_id = users.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		var booking domain.Booking
		if err := rows.Scan(&booking.ID, &booking.RoomName, &booking.StartTime, &booking.EndTime, &booking.UserID, &booking.Username); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (r *BookingRepository) GetByID(id int) (*domain.Booking, error) {

	row := r.db.QueryRow("SELECT bookings.id, room_name, start_time, end_time, user_id, username FROM bookings JOIN users ON user_id = users.id WHERE bookings.id = $1", id)

	var booking domain.Booking
	if err := row.Scan(&booking.ID, &booking.RoomName, &booking.StartTime, &booking.EndTime, &booking.UserID, &booking.Username); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *BookingRepository) Create(booking domain.Booking) error {

	_, err := r.db.Exec("INSERT INTO bookings (room_name, start_time, end_time, user_id) VALUES ($1, $2, $3, $4)", booking.RoomName, booking.StartTime, booking.EndTime, booking.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) Update(booking domain.Booking) error {

	_, err := r.db.Exec("UPDATE bookings SET room_name = $1, start_time = $2, end_time = $3, user_id = $4 WHERE id = $5", booking.RoomName, booking.StartTime, booking.EndTime, booking.UserID, booking.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) Delete(id int) error {

	_, err := r.db.Exec("DELETE FROM bookings WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
