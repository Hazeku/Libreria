// Navbar.jsx
import React, { useState } from 'react';
import { BiMenu } from 'react-icons/bi';


function Navbar({ categories, onSelectCategory, cartItems, removeFromCart }) {
  const [showCart, setShowCart] = useState(false);
  const [showDropdown, setShowDropdown] = useState(false);

  const enviarPorWhatsApp = () => {
    const numeroTelefono = "+543865653191";
    if (cartItems.length === 0) {
      alert("Tu carrito está vacío");
      return;
    }

    const mensaje = encodeURIComponent(
      `¡Hola! Quiero comprar estos productos:\n\n` +
      cartItems.map(item => `- ${item.title} x ${item.quantity}`).join("\n")
    );

    window.open(`https://wa.me/${numeroTelefono}?text=${mensaje}`, "_blank");
  };

  return (
    <nav className="navbar">
      <a href="/" className="navbar-brand">Home</a>

      <div className="navbar-dropdown">
        <button
          onClick={() => setShowDropdown(!showDropdown)}
          className="dropdown-toggle"
        >
          <BiMenu size={24} />
        </button>
        {showDropdown && (
          <div className="dropdown-menu animate-fade-in">
            {categories.map((category, index) => (
              <button
                key={index}
                onClick={() => {
                  onSelectCategory(category);
                  setShowDropdown(false);
                }}
                className="dropdown-item"
              >
                {category}
              </button>
            ))}
          </div>
        )}
      </div>

      <a href="/admin-login" className="navbar-link">Admin</a>

      <div className="cart">
        <button
          onClick={() => setShowCart(!showCart)}
          className="cart-toggle"
        >
          <i className="bi bi-cart"></i> ({cartItems.length})
        </button>
        {showCart && (
          <div className="cart-dropdown animate-fade-in">
            <ul className="cart-items">
              {cartItems.map((item, index) => (
                <li key={index} className="cart-item">
                  <span>{item.title} x {item.quantity}</span>
                  <button onClick={() => removeFromCart(item.id)}>
                    <i className="bi bi-x"></i>
                  </button>
                </li>
              ))}
            </ul>
            <button
              onClick={enviarPorWhatsApp}
              className="whatsapp-btn"
            >
              <i className="bi bi-whatsapp"></i> Enviar pedido
            </button>
          </div>
        )}
      </div>
    </nav>
  );
}

export default Navbar;
