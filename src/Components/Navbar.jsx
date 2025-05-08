import React from 'react';

function Navbar({ categories, onSelectCategory, cartItems, removeFromCart }) {
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

  /*     const calcularTotal = () => {
        let total = 0;
        cartItems.forEach(item => {
          const price = parseFloat(item.price.replace('$', ''));
          if (!isNaN(price)) {
            total += price * item.quantity;
          }
        });
        return total.toFixed(2);
      };
  
      const total = calcularTotal(); */

  /*     const mensaje = encodeURIComponent(
        `¡Hola! Quiero comprar estos productos:\n\n` +
        cartItems.map(item => `- ${item.title} x ${item.quantity} 
        (${item.price})`).join("\n") + `\n\nTotal: $${total}`
      );
  
      window.open(`https://wa.me/${numeroTelefono}?text=${mensaje}`, "_blank");
    }; */

  return (
    <nav className="navbar">
      <div className="navcontainer">
        <a href="/" className="navbar-brand">Inicio</a>
        <ul className="navbar-nav">
          <li className="nav-item">
            
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
        <li className="nav-item">
  <a href="/admin-login" className="nav-link">Admin</a>
</li>


        {/* Carrito dentro de la barra */}
        <div className="cart">
          <button className="cart-toggle">
            <i className="bi bi-cart"></i> {cartItems.length}
          </button>
          {cartItems.length > 0 && (
            <div className="cart-dropdown">
              <ul>
                {cartItems.map((item, index) => (
                  <li key={index}>
                    {item.title} x {item.quantity}
                    <button onClick={() => removeFromCart(item.id)}>
                      <i className="bi bi-cart-dash"></i>
                    </button>
                  </li>
                ))}
              </ul>
              {/* Botón para enviar por WhatsApp */}
              <button onClick={enviarPorWhatsApp} className="whatsapp-btn">
                <i className="bi bi-whatsapp"></i> Enviar pedido
              </button>
            </div>
          )}
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
