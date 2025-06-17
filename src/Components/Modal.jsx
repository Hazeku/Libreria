import React from 'react';
import { HiOutlineShoppingBag } from 'react-icons/hi2'; // Modern bag icon


function Modal({ item, onClose, onAddToCart }) {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <img src={item.image} alt={item.title} className="modal-image" />
        <div className="modal-description">
          <h3>{item.title}</h3>
          <p>{item.description}</p>
        </div>
        
        {/* Botón para agregar al carrito */}
        <button className="modal-add-to-cart" onClick={() => onAddToCart(item)}>
  <HiOutlineShoppingBag size={20} style={{ marginRight: '8px' }} />
  Agregar al carrito
</button>


        {/* Botón de cierre */}
        <button className="modal-close" onClick={onClose}>X</button>
      </div>
    </div>
  );
}

export default Modal;
