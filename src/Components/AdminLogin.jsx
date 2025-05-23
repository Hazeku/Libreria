import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../Styles/AdminLogin.css';
import API_URL from '../api/config';

const AdminLogin = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    console.log("🔐 Enviando login a:", `${API_URL}/login`);
    e.preventDefault();
    try {
      const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
        credentials: 'include',
      });

      if (!response.ok) {
        setError('Credenciales inválidas');
        return;
      }

      const data = await response.json();
      console.log("🔐 Respuesta de login:", data);
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
