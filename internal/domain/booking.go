package domain

import "time"

type Booking struct {
	ID        int       `json:"id"`
	RoomName  string    `json:"room_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
}
