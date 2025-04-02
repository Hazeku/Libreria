import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../Styles/AdminLogin.css';

const AdminLogin = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8000/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        setError('Credenciales inválidas');
        return;
      }

      const data = await response.json();
      localStorage.setItem('adminToken', data.token);
      navigate('/admin');
    } catch (err) {
      setError('Error al conectar con el servidor');
    }
  };

  return (
    <div className="admin-login-container">
      <div className="admin-login-box">
        <h2>Login de Administrador</h2>
        <form onSubmit={handleSubmit} className="admin-login-form">
          <input
            type="text"
            className="admin-login-input"
            placeholder="Usuario"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <input
            type="password"
            className="admin-login-input"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <button type="submit" className="admin-login-button">Ingresar</button>
        </form>
        {error && <p className="admin-login-error">{error}</p>}
      </div>
    </div>
  );
};

export default AdminLogin;
