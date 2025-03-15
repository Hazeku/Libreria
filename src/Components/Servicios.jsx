import React, { useState, useEffect } from 'react';

function Servicios() {
  const [servicios, setServicios] = useState([
    {
      nombre: 'Impresión',
      imagenes: ['/Images/impresión1.webp', '/Images/impresión2.webp', '/Images/impresión3.webp'],
      descripcion: 'Ofrecemos servicios de impresión.'
    },
    {
      nombre: 'Anillado',
      imagenes: ['/Images/anillado1.webp', '/Images/anillado2.webp', '/Images/anillado3.webp'],
      descripcion: 'Realizamos anillado de calidad para presentaciones, informes y más.'
    },
    {
      nombre: 'Fotocopias',
      imagenes: ['/Images/fotocopias1.webp', '/Images/fotocopias2.webp', '/Images/fotocopias3.webp'],
      descripcion: 'Contamos con equipos de fotocopiado modernos y rápidos para satisfacer tus necesidades de reproducción.'
    }
    
    // Agrega más objetos de servicios según sea necesario
  ]);
  const [indexImagen, setIndexImagen] = useState(0); // Índice actual de la imagen en cada servicio

  useEffect(() => {
    const intervalId = setInterval(() => {
      setIndexImagen(prevIndex => (prevIndex + 1) % 3); // Cambia a la siguiente imagen
    }, 5000); // Cambia la imagen cada 5 segundos

    return () => clearInterval(intervalId); // Limpia el intervalo al desmontar el componente
  }, []);

  return (
    <section className="servicios">
      <h2>Servicios</h2>
      <div className="servicios-lista">
        {servicios.map((servicio, index) => (
          <div key={index} className="servicio-item">
            <img src={servicio.imagenes[indexImagen]} alt={servicio.nombre} />
            <h3>{servicio.nombre}</h3>
            <p>{servicio.descripcion}</p>
          </div>
        ))}
      </div>
    </section>
  );
}

export default Servicios;
