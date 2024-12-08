import React, { useEffect, useState } from 'react';
import { getAllBookings, createBooking, logoutUser, deleteBooking } from '../api';
import '../App.css';
import '../api.js';
const BookingPage = () => {
    const [bookings, setBookings] = useState([]);
    const [roomName, setRoomName] = useState('');
    const [startTime, setStartTime] = useState('');
    const [endTime, setEndTime] = useState('');

    const convertToISO = (datetimeLocal) => {
        const date = new Date(datetimeLocal);
        return date.toISOString();
    };

    const fetchBookings = async () => {
        const response = await getAllBookings();
        setBookings(response.data.data);
    };

    const handleCreateBooking = async () => {
        try {
            let user = JSON.parse(localStorage.getItem('user'));
            await createBooking({
                room_name: roomName,
                start_time: convertToISO(startTime),
                end_time: convertToISO(endTime),
                user_id: user.id,
            });
            fetchBookings();
            alert('Booking created successfully');
        } catch (error) {
            alert('Failed to create booking');
        }
    };

    const handleLogout = () => {
        logoutUser();
        window.location.reload();
    };

    const handleDeleteBooking = async (id) => {
        try {
            await deleteBooking(id);
            alert('Booking deleted successfully');
            fetchBookings();
        } catch (error) {
            alert('Failed to delete booking');
        }
    };

    useEffect(() => {
        fetchBookings();
    }, []);

    return (
        <div className="booking-page">
            <h2>Booking Page</h2>
            <button onClick={handleLogout}>Logout</button>
            <h3>Create Booking</h3>
            <div className="booking-form">
                <input
                    type="text"
                    placeholder="Room Name"
                    value={roomName}
                    onChange={(e) => setRoomName(e.target.value)}
                />
                <input
                    type="datetime-local"
                    placeholder="Start Time"
                    value={startTime}
                    onChange={(e) => setStartTime(e.target.value)}
                />
                <input
                    type="datetime-local"
                    placeholder="End Time"
                    value={endTime}
                    onChange={(e) => setEndTime(e.target.value)}
                />
                <button onClick={handleCreateBooking}>Create Booking</button>
            </div>
            <h3>Bookings</h3>
            <ul className="bookings-list">
                {bookings && bookings.length > 0 ? (
                    bookings.map((booking) => (
                        <li key={booking.id}>
                            <p>
                                {booking.room_name} — {new Date(booking.start_time).toLocaleString()} to {new Date(booking.end_time).toLocaleString()} by {booking.username}
                            </p>
                            <button onClick={() => handleDeleteBooking(booking.id)}>Delete</button>
                        </li>
                    ))
                ) : (
                    <p>No bookings available</p>
                )}
            </ul>
        </div>
    );
};

export default BookingPage;
