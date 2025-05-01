import API_URL from '../api/config';

import React from 'react';

const AdminDashboard = () => {
  const addArticle = async (newArticle) => {
    const token = localStorage.getItem('adminToken');
    const response = await fetch(`${API_URL}/articles`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(newArticle),
    });
    // Podés manejar la respuesta aquí, si querés mostrar algo
    const data = await response.json();
    console.log(data);
  };

  return (
    <div>
      <h1>Panel de administrador</h1>
      {/* Acá podés agregar formularios o botones que llamen a addArticle */}
      <button onClick={() => addArticle({ title: 'Nuevo artículo' })}>
        Agregar artículo de prueba
      </button>
    </div>
  );
};

export default AdminDashboard;