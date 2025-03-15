import React from 'react';

function ItemList({ items, onItemClick }) {
  return (
    <div className="item-list">
      <h2>Artículos</h2>
      <ul>
        {items.map((item, index) => (
          <li key={index} data-aos="fade-up" onClick={() => onItemClick(item)}>
            <h3>{item.title}</h3>
            <img src={item.image} alt={item.title} />
            <p>{item.description}</p>
            {/* <p>{item.price}</p> */}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ItemList;
