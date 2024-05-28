import React from 'react';
import '../Styles/Footer.css';

function Footer() {
  return (
    <footer className="footer" data-aos="fade-up">
      <div className="footercontainer">
        <p>© 2024 Mi Librería. Todos los derechos reservados.</p>
        <div className="footer-address">
          <p>Dirección: 25 de Mayo 430, Concepción, Tucumán</p>
        </div>
        <iframe
            src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3544.1034917724173!2d-65.59385878837526!3d-27.341232811147453!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x9423cf43b41514b7%3A0xe47a5fec05f78a4!2s25%20de%20Mayo%20430%2C%20T4146%20Concepci%C3%B3n%2C%20Tucum%C3%A1n!5e0!3m2!1sen!2sar!4v1716923093518!5m2!1sen!2sar"
            width="200"
            height="200"
            style={{ border: 0 }}
            allowFullScreen=""
            loading="lazy"
            referrerPolicy="no-referrer-when-downgrade"
          ></iframe>
      </div>
    </footer>
  );
}

export default Footer;
