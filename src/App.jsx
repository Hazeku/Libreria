import React, { useEffect, useState, useRef } from 'react';
import './Styles/App.css';
import './Styles/Navbar.css';
import './Styles/CategoryList.css';
import './Styles/ItemList.css';
import './Styles/Footer.css';
import './Styles/Carousel.css';
import './Styles/Modal.css';
import './Styles/Servicios.css';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import CategoryList from './Components/CategoryList';
import ItemList from './Components/ItemList';
import Navbar from './Components/Navbar';
import Servicios from './Components/Servicios';
import Articles from './Data/Articles';
import Footer from './Components/Footer';
import Carousel from './Components/Carousel';
import Modal from './Components/Modal';
import AOS from 'aos';
import 'aos/dist/aos.css';
import Titulo from './Components/Titulo';

function App() {
  useEffect(() => {
    AOS.init({
      duration: 1000,
      offset: 100,
      delay: 100,
      easing: 'ease-in-out',
      once: false,
    });
  }, []);

  const categories = [
    "Instrumentos escolares",
    "Suministros escolares",
    "Libros",
    "Utilidades",
    "Biblioratos",
    "Carpetas",
    "Material de Arte, Manualidades, Decoraciones"
  ];

  const [selectedCategory, setSelectedCategory] = useState(null);
  const [selectedItem, setSelectedItem] = useState(null);
  const [cartItems, setCartItems] = useState([]);

  const categoryListRef = useRef(null);

  const handleCategorySelect = (category) => {
    setSelectedCategory(category);
    scrollToCategoryList();
  };

  const handleItemClick = (item) => {
    setSelectedItem(item);
  };

  const handleAddToCart = (item) => {
    const existingItem = cartItems.find((cartItem) => cartItem.id === item.id);
    if (existingItem) {
      setCartItems(
        cartItems.map((cartItem) =>
          cartItem.id === item.id ? { ...cartItem, quantity: cartItem.quantity + 1 } : cartItem
        )
      );
    } else {
      setCartItems([...cartItems, { ...item, quantity: 1 }]);
    }
    setSelectedItem(null); // Cierra el modal después de agregar al carrito
  };

  const removeFromCart = (id) => {
    setCartItems(cartItems.filter((item) => item.id !== id));
  };

  const closeModal = () => {
    setSelectedItem(null);
  };

  const filteredArticles = selectedCategory
    ? Articles.filter((article) => article.category === selectedCategory)
    : Articles;

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
          onAddToCart={handleAddToCart}  // Pasa la función al modal
          onClose={closeModal} 
        />
      )}
    </div>
  );
}

export default App;
