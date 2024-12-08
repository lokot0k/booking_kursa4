import React, { useState } from 'react';
import '../App.css';
import {registerUser } from '../api.js'
const RegisterPage = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleRegister = async (event) => {
        event.preventDefault();
        try {
            await registerUser({ name: username, password: password });
            alert('Registration successful. You can now log in.');
        } catch (error) {
            alert('Registration failed. Please try again.');
        }
        setUsername('');
        setPassword('');
    };

    return (
        <div className="register-form">
            <h2>Register</h2>
            <input
                type="text"
                placeholder="Username"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
            />
            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <button onClick={handleRegister}>Register</button>
        </div>
    );
};

export default RegisterPage;
