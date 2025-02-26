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
            href="https://maps.app.goo.gl/X5LguGDKd5zubhqU9" 
            target="_blank" 
            rel="noopener noreferrer"
          >
            Ver en Google Maps
          </a>
        </div>
        <div className="footer-section">
          <h4>Contáctanos</h4>
          <p>Teléfono: +54 3865-579016</p>
        </div>
        <div className="footer-section">
          <h4>Horarios de atencion</h4>
          <p>08:00 - 12:00</p>
          <p>14:00 - 21:00</p>
          
        </div>
      </div>
    </footer>
  );
}

export default Footer;
