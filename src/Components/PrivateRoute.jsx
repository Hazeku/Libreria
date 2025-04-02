// PrivateRoute.jsx
import React from 'react';
import { Navigate } from 'react-router-dom';

const PrivateRoute = ({ children }) => {
  const token = localStorage.getItem('adminToken');
  
  if (!token) {
    // Si no hay token, redirige al login de admin
    return <Navigate to="/admin-login" replace />;
  }
  
  // Si hay token, renderiza el componente protegido
  return children;
};

export default PrivateRoute;
