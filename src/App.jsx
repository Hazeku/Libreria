import React, { useEffect, useState,useRef } from 'react';
import logo from './logo.svg';
import './Styles/App.css';
import './Styles/Navbar.css';
import './Styles/CategoryList.css';
import './Styles/ItemList.css';
import './Styles/Footer.css';
import './Styles/Carousel.css';
import './Styles/Modal.css'; // Importar los estilos del modal
import './Styles/Servicios.css';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import CategoryList from './Components/CategoryList';
import ItemList from './Components/ItemList';
import Navbar from './Components/Navbar';
import Servicios from './Components/Servicios';
import articles from './Data/Articles';
import Footer from './Components/Footer';
import Carousel from './Components/Carousel';
import Modal from './Components/Modal'; // Importar el componente Modal
import AOS from 'aos';
import 'aos/dist/aos.css';

function App() {
  useEffect(() => {
    AOS.init({
      duration: 1000,
      offset: 150,
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

  const categoryListRef = useRef(null); // Referencia al elemento de la lista de categorías

  const handleCategorySelect = (category) => {
    setSelectedCategory(category);
    console.log("Se seleccionó la categoría:", category);
    scrollToCategoryList(); // Desplaza la ventana de visualización hasta la lista de categorías
  };

  const handleItemClick = (item) => {
    setSelectedItem(item);
  };

  const closeModal = () => {
    setSelectedItem(null);
  };

  const filteredArticles = selectedCategory
    ? articles.filter(article => article.category === selectedCategory)
    : articles;

  const scrollToCategoryList = () => {
    categoryListRef.current.scrollIntoView({ behavior: 'smooth' }); // Desplaza suavemente hasta la lista de categorías
  };

  return (
    <div className="App">
      <Navbar categories={categories} onSelectCategory={handleCategorySelect} />
      <Carousel />
      <Servicios />
      <div className="container" ref={categoryListRef}> {/* Agrega la referencia al contenedor */}
        <CategoryList categories={categories} onSelectCategory={handleCategorySelect} />
        <ItemList items={filteredArticles} onItemClick={handleItemClick} />
      </div>
      <Footer />
      {selectedItem && <Modal item={selectedItem} onClose={closeModal} />}
    </div>
  );
}

export default App;