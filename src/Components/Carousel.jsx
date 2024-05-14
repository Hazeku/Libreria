import React, { useState, useEffect } from 'react';

function Carousel() {
  const [currentIndex, setCurrentIndex] = useState(0);
  const images = [
    '/Images/slider1.jpg',
    '/Images/slider2.jpg',
    '/Images/slider3.jpg'
  ];

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentIndex((prevIndex) => (prevIndex + 1) % images.length);
    }, 5000);

    return () => clearInterval(interval);
  }, [images.length]);

  return (
    <div className="carousel-container">
      <div className="carousel" style={{ transform: `translateX(-${currentIndex * 100}%)` }}>
        {images.map((image, index) => (
          <img key={index} src={image} alt={`Slide ${index + 1}`} />
        ))}
      </div>
    </div>
  );
}

export default Carousel;
