// src/Components/PublicHome.jsx
import React, { useEffect, useState, useRef } from 'react';
import '../Styles/App.css';
import '../Styles/Navbar.css';
import '../Styles/CategoryList.css';
import '../Styles/ItemList.css';
import '../Styles/Footer.css';
import '../Styles/Carousel.css';
import '../Styles/Modal.css';
import '../Styles/Servicios.css';
import CategoryList from './CategoryList';
import ItemList from './ItemList';
import Navbar from './Navbar';
import Servicios from './Servicios';
import Footer from './Footer';
import Carousel from './Carousel';
import Modal from './Modal';
import Titulo from './Titulo';
import AOS from 'aos';
import 'aos/dist/aos.css';
import 'bootstrap-icons/font/bootstrap-icons.css';

import API_URL from '../api/config';

const PublicHome = () => {
  const [articles, setArticles] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState(null);
  const [selectedItem, setSelectedItem] = useState(null);
  const [cartItems, setCartItems] = useState([]);
  const categoryListRef = useRef(null);

  useEffect(() => {
    AOS.init({ duration: 1000, offset: 100, delay: 100, easing: 'ease-in-out', once: false });
    fetchArticles();
  }, []);

  const fetchArticles = async () => {
    try {
      const res = await fetch(`${API_URL}/articles`);
      const data = await res.json();
      if (Array.isArray(data.articles)) {
        // ✅ Agregamos la URL completa para las imágenes
        const articlesConImagenes = data.articles.map(article => ({
          ...article,
          image: `${API_URL}${article.image}` // ejemplo: http://localhost:8080/Images/123.webp
        }));
        setArticles(articlesConImagenes);
      } else {
        setArticles([]);
      }
    } catch (error) {
      console.error('Error al cargar artículos:', error);
    }
  };

  const categories = [
    "Instrumentos escolares", "Suministros escolares", "Libros", "Utilidades",
    "Biblioratos", "Carpetas", "Material de Arte, Manualidades, Decoraciones"
  ];

  const handleCategorySelect = (category) => {
    setSelectedCategory(category);
    scrollToCategoryList();
  };

  const handleItemClick = (item) => {
    setSelectedItem(item);
  };

  const handleAddToCart = (item) => {
    const existingItem = cartItems.find(cartItem => cartItem.id === item.id);
    if (existingItem) {
      setCartItems(cartItems.map(cartItem =>
        cartItem.id === item.id ? { ...cartItem, quantity: cartItem.quantity + 1 } : cartItem
      ));
    } else {
      setCartItems([...cartItems, { ...item, quantity: 1 }]);
    }
    setSelectedItem(null);
  };

  const removeFromCart = (id) => {
    setCartItems(cartItems.filter(item => item.id !== id));
  };

  const closeModal = () => setSelectedItem(null);

  const filteredArticles = selectedCategory
    ? articles.filter(article => article.category === selectedCategory)
    : articles;

  const scrollToCategoryList = () => {
    categoryListRef.current.scrollIntoView({ behavior: 'smooth' });
  };

  return (
    <div className="App">
      <Titulo />
      <div className="navbar-container">
        <Navbar
          categories={categories}
          onSelectCategory={handleCategorySelect}
          cartItems={cartItems}
          removeFromCart={removeFromCart}
        />
      </div>
      <Carousel />
      <Servicios />
      <div className="container" ref={categoryListRef}>
        <CategoryList categories={categories} onSelectCategory={handleCategorySelect} />
        <ItemList items={filteredArticles} onItemClick={handleItemClick} />
      </div>
      <Footer />
      {selectedItem && (
        <Modal
          item={selectedItem}
          onAddToCart={handleAddToCart}
          onClose={closeModal}
        />
      )}
    </div>
  );
};

export default PublicHome;
