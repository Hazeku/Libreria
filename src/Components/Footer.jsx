import React from 'react';
import '../Styles/Footer.css';

function Footer() {
  return (
    <footer className="footer" data-aos="fade-up">
      <div className="footercontainer">
        <p>© 2024 Mi Librería. Todos los derechos reservados.</p>
        <div className="footer-address">
          <p>Dirección: 25 de Mayo 430, Concepción, Tucumán</p>
          <a 
            href="https://maps.app.goo.gl/ab4CGUjjyeduAHZ58" 
            target="_blank" 
            rel="noopener noreferrer"
          >
            Ver en Google Maps
          </a>
        </div>
      </div>
    </footer>
  );
}

export default Footer;
