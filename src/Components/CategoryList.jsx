import React from 'react';

function CategoryList({ categories, onSelectCategory }) {
  return (
    <div className="category-list">
      <h2>Categorías</h2>
      <ul>
        {categories.map((category, index) => (
          <li key={index} onClick={() => onSelectCategory(category)}>
            {category}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default CategoryList;
