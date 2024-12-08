import axios from 'axios';

const API_URL = 'http://localhost:4000';

const getAuthHeader = () => {
  const user = JSON.parse(localStorage.getItem('user'));
  return user ? { Authorization: `Basic ${user.name}:${user.password}` } : {};
};

export const loginUser = async (username, password) => {
  const credentials = `${username}:${password}`;
  const authToken = `Basic ${credentials}`;

  return axios.post(`${API_URL}/v1/login`, {}, {
    headers: { Authorization: authToken },
  });
};

export const registerUser = async (data) => {
  return axios.post(`${API_URL}/v1/register`, data);
};

export const logoutUser = () => {
  localStorage.removeItem('user');
};

export const getAllBookings = async () => {
  return axios.get(`${API_URL}/v1/bookings`, { headers: getAuthHeader() });
};

export const createBooking = async (data) => {
  return axios.post(`${API_URL}/v1/booking`, data, { headers: getAuthHeader() });
};

export const deleteBooking = async (id) => {
  return axios.delete(`${API_URL}/v1/booking/${id}`, {
    headers: getAuthHeader(),
  });
};