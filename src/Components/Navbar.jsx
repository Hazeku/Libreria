import React from 'react';

function Navbar({ categories, onSelectCategory, cartItems, removeFromCart }) {
  return (
    <nav className="navbar">
      <div className="navcontainer">
        <a href="/" className="navbar-brand">Logo</a>
        <ul className="navbar-nav">
          <li className="nav-item">
            <a href="/" className="nav-link">Inicio</a>
          </li>
          <li className="nav-item dropdown">
            <a href="#!" className="nav-link dropdown-toggle" id="navbarDropdown" role="button">
              Categorías
            </a>
            <div className="dropdown-menu">
              {categories.map((category, index) => (
                <a
                  key={index}
                  className="dropdown-item"
                  href="#!"
                  onClick={() => onSelectCategory(category)}
                >
                  {category}
                </a>
              ))}
            </div>
          </li>
        </ul>

        {/* Carrito dentro de la barra */}
        <div className="cart">
          <button className="cart-toggle">
            🛒 {cartItems.length}
          </button>
          {cartItems.length > 0 && (
            <div className="cart-dropdown">
              <ul>
                {cartItems.map((item, index) => (
                  <li key={index}>
                    {item.title} x {item.quantity} {/* Cambié name por title */}
                    <button onClick={() => removeFromCart(item.id)}>❌</button>
                  </li>
                ))}
              </ul>
            </div>
          )}
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
