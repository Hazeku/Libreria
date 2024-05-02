import React from 'react';

function ItemList({ items }) {
  return (
    <div className="item-list">
      <h2>Artículos</h2>
      <ul>
        {items.map((item, index) => (
          <li key={index}>
            <h3>{item.title}</h3>
            <p>{item.description}</p>
            <p>Categoría: {item.category}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ItemList;
