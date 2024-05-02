import React from 'react';

function Navbar() {
  return (
    <nav className="navbar">
      <div className="container">
        <a href="/" className="navbar-brand">Logo</a>
        <ul className="navbar-nav">
          <li className="nav-item">
            <a href="/" className="nav-link">Inicio</a>
          </li>
          <li className="nav-item">
            <a href="/categorias" className="nav-link">Categorías</a>
          </li>
          <li className="nav-item">
            <a href="/carrito" className="nav-link">Carrito</a>
          </li>
          <li className="nav-item">
            <a href="/perfil" className="nav-link">Perfil</a>
          </li>
        </ul>
      </div>
    </nav>
  );
}

export default Navbar;
