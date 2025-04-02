import Articles from './articles.js';

// Función para procesar un objeto y generar una descripción si está vacía
function processObject(obj) {
  const newObj = { ...obj }; // Crea una copia del objeto para no modificar el original
  if (newObj.description === "") {
    newObj.description = "Sin descripción disponible."; // Puedes usar otro valor por defecto
  }
  return newObj;
}

// Procesa cada objeto del arreglo Articles
const processedArticles = Articles.map(processObject);

// Convierte el arreglo procesado a JSON
const jsonData = JSON.stringify(processedArticles, null, 2);

console.log(jsonData);