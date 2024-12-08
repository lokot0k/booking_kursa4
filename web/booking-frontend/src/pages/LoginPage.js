import React, { useState } from 'react';
import '../App.css';
import {loginUser} from '../api.js'
const LoginPage = ({ onLogin }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = async (event) => {
        event.preventDefault();
        try {
            let response = await loginUser(username, password);
            if (response.status === 200) {
                localStorage.setItem('user', JSON.stringify(response.data.user));
                onLogin();
            }
        } catch (error) {
            alert('Login failed. Please try again.');
        }
        setUsername('');
        setPassword('');
    };

    return (
        <div className="login-form">
            <h2>Login</h2>
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
            <button onClick={handleLogin}>Login</button>
        </div>
    );
};
export default LoginPage;
