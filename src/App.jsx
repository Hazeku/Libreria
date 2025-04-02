// App.jsx
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import PublicHome from './Components/PublicHome';
import AdminLogin from './Components/AdminLogin';
import PrivateRoute from './Components/PrivateRoute';
import AdminDashboard from './Components/AdminDashboard';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Ruta para el login del administrador */}
        <Route path="/admin-login" element={<AdminLogin />} />
        
        {/* Ruta protegida para el administrador */}
        <Route 
          path="/admin" 
          element={
            <PrivateRoute>
              <AdminDashboard />
            </PrivateRoute>
          } 
        />

        {/* Ruta pública */}
        <Route path="/*" element={<PublicHome />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
