import React, { useState } from 'react';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import BookingPage from './pages/BookingPage';
import './App.css';

const App = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(!!localStorage.getItem('user'));

  const handleLogin = () => {
    setIsLoggedIn(true);
  };

  return (
      <div className="container">
        {!isLoggedIn ? (
            <>
              <LoginPage onLogin={handleLogin} />
              <RegisterPage />
            </>
        ) : (
            <BookingPage />
        )}
      </div>
  );
};

export default App;
