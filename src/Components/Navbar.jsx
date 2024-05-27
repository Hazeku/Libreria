import React from 'react';

function Navbar({ categories, onSelectCategory }) {
  return (
    <nav className="navbar">
      <div className="navcontainer">
        <a href="/" className="navbar-brand">Logo</a>
        <ul className="navbar-nav">
          <li className="nav-item">
            <a href="/" className="nav-link">Inicio</a>
          </li>
          <li className="nav-item dropdown">
            <a href="#!" className="nav-link dropdown-toggle" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
              Categorías
            </a>
            <div className="dropdown-menu" aria-labelledby="navbarDropdown">
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
      </div>
    </nav>
  );
}

export default Navbar;
