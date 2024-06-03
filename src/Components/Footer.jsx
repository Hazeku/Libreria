import React from 'react';
import '../Styles/Footer.css';

function Footer() {
  return (
    <footer className="footer" data-aos="fade-up">
      <div className="footer-container">
        <div className="footer-section">
          <p>© 2024 Mi Librería. Todos los derechos reservados.</p>
        </div>
        <div className="footer-section">
          <p>Dirección: 25 de Mayo 430, Concepción, Tucumán</p>
          <a 
            href="https://maps.app.goo.gl/ab4CGUjjyeduAHZ58" 
            target="_blank" 
            rel="noopener noreferrer"
          >
            Ver en Google Maps
          </a>
        </div>
        <div className="footer-section">
          <h4>Contáctanos</h4>
          <p>Email: info@milibreria.com</p>
          <p>Teléfono: +54 9 381 123 4567</p>
        </div>
        <div className="footer-section">
          <h4>Síguenos</h4>
          <a href="https://www.facebook.com/milibreria" target="_blank" rel="noopener noreferrer">Facebook</a>
          <a href="https://www.twitter.com/milibreria" target="_blank" rel="noopener noreferrer">Twitter</a>
          <a href="https://www.instagram.com/milibreria" target="_blank" rel="noopener noreferrer">Instagram</a>
        </div>
        <div className="footer-section">
          <h4>Términos y Condiciones</h4>
          <a href="/terms-and-conditions" target="_blank" rel="noopener noreferrer">Leer más</a>
        </div>
        <div className="footer-section">
          <h4>Suscríbete a nuestro boletín</h4>
          <form className="subscription-form">
            <input type="email" placeholder="Tu correo electrónico" />
            <button type="submit">Suscribirse</button>
          </form>
        </div>
      </div>
    </footer>
  );
}

export default Footer;
