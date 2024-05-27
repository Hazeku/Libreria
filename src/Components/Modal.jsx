import React from 'react';

function Modal({ item, onClose }) {
  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <img src={item.image} alt={item.title} className="modal-image" />
        <div className="modal-description">
          <h3>{item.title}</h3>
          <p>{item.description}</p>
        </div>
        <button className="modal-close" onClick={onClose}>X</button>
      </div>
    </div>
  );
}

export default Modal;
