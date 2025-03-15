import React, { useState, useEffect } from 'react';
import '../Styles/Carousel.css'; // Asegúrate de tener este archivo

function Carousel() {
  const [currentIndex, setCurrentIndex] = useState(0);
  const images = [
    '/Images/slider1.jpeg',
    '/Images/slider2.jpeg',
    '/Images/slider3.jpeg'
  ];

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentIndex((prevIndex) => (prevIndex + 1) % images.length);
    }, 10000);

    return () => clearInterval(interval);
  }, [images.length]);

  return (
    <div className="carousel-container">
      {images.map((image, index) => (
        <img
          key={index}
          src={image}
          alt={`Slide ${index + 1}`}
          className={index === currentIndex ? 'active' : 'inactive'}
        />
      ))}
    </div>
  );
}

export default Carousel;
