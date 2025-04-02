const addArticle = async (newArticle) => {
    const token = localStorage.getItem('adminToken');
    const response = await fetch('http://localhost:8000/articles', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify(newArticle),
    });
    // Manejar la respuesta
  };
  