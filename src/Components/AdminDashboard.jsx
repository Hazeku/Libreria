import React, { useEffect, useState, useCallback } from 'react';
import Modal from 'react-modal';
import API_URL from '../api/config';
import '../Styles/AdminDashboard.css';
import AOS from 'aos';
import 'aos/dist/aos.css';

Modal.setAppElement('#root');

const AdminDashboard = () => {
  const [articles, setArticles] = useState([]);
  const [categories, setCategories] = useState([]);
  const [modalIsOpen, setModalIsOpen] = useState(false);
  const [editingArticle, setEditingArticle] = useState(null);
  const [imagePreview, setImagePreview] = useState(null);
  const [imageFile, setImageFile] = useState(null);

  const [formData, setFormData] = useState({
    title: '',
    description: '',
    price: '',
    image: '',
    category: '',
  });

  const token = localStorage.getItem('adminToken');

  // ✅ PRIMERO definimos estas funciones antes del useEffect
  const fetchArticles = useCallback(async () => {
    try {
      const res = await fetch(`${API_URL}/articles`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      const data = await res.json();
      if (Array.isArray(data)) {
        setArticles(data);
      } else if (Array.isArray(data.articles)) {
        setArticles(data.articles);
      } else {
        setArticles([]);
      }
    } catch (error) {
      console.error('Error al obtener artículos:', error);
    }
  }, [token]);

  const fetchCategories = useCallback(async () => {
    try {
      const res = await fetch(`${API_URL}/categories`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      const data = await res.json();
      setCategories(data);
    } catch (error) {
      console.error('Error al obtener categorías:', error);
    }
  }, [token]);

  // ✅ Ahora sí, se usan abajo
  useEffect(() => {
    AOS.init();
    fetchArticles();
    fetchCategories();
    return () => AOS.refresh();
  }, [fetchArticles, fetchCategories]);

  useEffect(() => {
    AOS.refresh();
  }, [articles]);

  const openModal = (article = null) => {
    if (article) {
      setFormData({
        title: article.title || '',
        description: article.description || '',
        price: article.price || '',
        image: article.image || '',
        category: article.category || '',
      });
      setEditingArticle(article.id);
    } else {
      setFormData({
        title: '',
        description: '',
        price: '',
        image: '',
        category: '',
      });
      setEditingArticle(null);
    }
    setModalIsOpen(true);
  };

  const closeModal = () => {
    setModalIsOpen(false);
    setEditingArticle(null);
    setImageFile(null);
    setImagePreview(null);
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const saveArticle = async () => {
    const method = editingArticle ? 'PUT' : 'POST';
    const url = editingArticle
      ? `${API_URL}/articles/${editingArticle}`
      : `${API_URL}/articles`;

    const form = new FormData();
    form.append('title', formData.title);
    form.append('description', formData.description);
    form.append('price', formData.price);
    form.append('category', formData.category);
    if (imageFile) {
      form.append('image', imageFile);
    }

    await fetch(url, {
      method,
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: form,
    });

    fetchArticles();
    closeModal();
  };

  const deleteArticle = async (id) => {
    if (window.confirm('¿Estás seguro de eliminar este artículo?')) {
      await fetch(`${API_URL}/articles/${id}`, {
        method: 'DELETE',
        headers: { Authorization: `Bearer ${token}` },
      });
      fetchArticles();
    }
  };

  return (
    <div className="item-list">
      <h1>Panel de administrador</h1>
      {/* 🔧 Contenedor aislado para el botón */}
    <div className="admin-add-container">
      <button onClick={() => openModal()} className="admin-add-button">
        ➕ Agregar artículo
      </button>
    </div>

      <ul>
        {articles.map((article) => (
          <li key={article.id} data-aos="fade-up">
            <h3>{article.title}</h3>
            <img
              src={`${API_URL}${article.image}`}
              alt={article.title}
              onError={(e) => {
                e.target.src = '/Images/placeholder.png';
              }}
            />
            <p>{article.description}</p>
            <p><strong>${article.price}</strong></p>
            <div className="admin-buttons">
              <button
                onClick={() => openModal(article)}
                className="admin-button-edit"
              >
                ✏️ Editar
              </button>
              <button
                onClick={() => deleteArticle(article.id)}
                className="admin-button-delete"
              >
                🗑️ Eliminar
              </button>
            </div>
          </li>
        ))}
      </ul>

      <Modal
        isOpen={modalIsOpen}
        onRequestClose={closeModal}
        className="modal"
        overlayClassName="modal-overlay"
      >
        <h2 className="text-xl font-bold mb-4">
          {editingArticle ? 'Editar artículo' : 'Nuevo artículo'}
        </h2>
        <input
          type="text"
          name="title"
          placeholder="Título"
          value={formData.title}
          onChange={handleInputChange}
          className="modal-input"
        />
        <select
          name="category"
          value={formData.category}
          onChange={handleInputChange}
          className="modal-input"
        >
          <option value="">Selecciona una categoría</option>
          {categories.map((category, index) => (
  <option key={index} value={category}>
    {category}
  </option>
))}
        </select>
        <input
          type="file"
          accept="image/*"
          onChange={(e) => {
            const file = e.target.files[0];
            setImageFile(file);
            if (file) {
              const reader = new FileReader();
              reader.onloadend = () => {
                setImagePreview(reader.result);
              };
              reader.readAsDataURL(file);
            }
          }}
          className="modal-input"
        />
        {imagePreview ? (
          <img src={imagePreview} alt="Preview" className="modal-preview-image" />
        ) : (
          formData.image && (
            <img
              src={`${API_URL}${formData.image}`}
              alt="Imagen actual"
              className="modal-preview-image"
            />
          )
        )}
        <input
          type="number"
          name="price"
          placeholder="Precio"
          value={formData.price}
          onChange={handleInputChange}
          className="modal-input"
        />
        <textarea
          name="description"
          placeholder="Descripción"
          value={formData.description}
          onChange={handleInputChange}
          className="modal-textarea"
        ></textarea>
        <div className="modal-actions">
          <button onClick={saveArticle} className="modal-save">
            💾 Guardar
          </button>
          <button onClick={closeModal} className="modal-cancel">
            Cancelar
          </button>
        </div>
      </Modal>
    </div>
  );
};

export default AdminDashboard;
