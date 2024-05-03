import logo from './logo.svg';
import './Styles/App.css';
import './Styles/Navbar.css';
import './Styles/CategoryList.css';
import './Styles/ItemList.css';
import './Styles/Footer.css';
import './Styles/Carousel.css';
import { BrowserRouter,Switch,Route } from 'react-router-dom';
import CategoryList from './Components/CategoryList';
import ItemList from './Components/ItemList';
import Navbar from './Components/Navbar';
import articles from './Data/Articles';
import Footer from './Components/Footer';
import Carousel from './Components/Carousel';



function App() {
  const categories = ["Fantasía", "Ficción", "Educación"]; // Puedes agregar más categorías aquí según sea necesario

  const handleCategorySelect = (category) => {
    // Lógica para filtrar artículos por categoría
    console.log("Se seleccionó la categoría:", category);
  };

  return (
    <div className="App">
      <Navbar />
      <div className="container">
        <CategoryList categories={categories} onSelectCategory={handleCategorySelect} />
        <ItemList items={articles} />
      </div>
      <Footer></Footer>
    </div>
  );
}

export default App;