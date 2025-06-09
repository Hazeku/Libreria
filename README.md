📚 Librería App

Bienvenido a Librería App, un sistema de gestión de librerías que permite administrar libros, usuarios y compras de manera eficiente. Este proyecto está compuesto por un backend en Go con GORM y PostgreSQL, y un frontend desarrollado con React.

🚀 Características principales

🔹 Backend (Go + GORM + PostgreSQL)

Autenticación con JWT.

Gestión de usuarios, libros y compras.

Conexión con PostgreSQL usando GORM.

Implementación de middleware para proteger rutas.

🔹 Frontend (React + Vite)

Interfaz intuitiva y amigable.

Consumo de API mediante Axios.

Autenticación de usuarios.

Gestión de libros y compras desde el dashboard.

🛠️ Instalación y configuración

📌 Prerrequisitos

Tener instalado Go y Node.js

PostgreSQL en ejecución

🔧 Instalación del Backend

cd backend-go
go mod tidy
go run main.go

Configurar base de datos en config.yaml o .env:

db_host: "localhost"
db_port: 5432
db_user: "tu_usuario"
db_password: "tu_contraseña"
db_name: "libreria_db"

🔧 Instalación del Frontend

cd frontend-react
npm install
npm run dev

🌐 Uso

Accede al frontend en http://localhost:5173

Regístrate o inicia sesión.

Gestiona libros, compras y usuarios.

📜 API Endpoints principales

POST /login → Inicia sesión y devuelve un token JWT.

GET /books → Lista todos los libros.

POST /books → Agrega un libro (requiere autenticación).

DELETE /books/:id → Elimina un libro (requiere autenticación).

🛡️ Seguridad

Uso de JWT para autenticación.

Middleware para proteger rutas sensibles.

CORS configurado en el backend.

🏗️ Tecnologías utilizadas

Backend: Go, GORM, PostgreSQL, JWT, Echo

Frontend: React, Vite, Axios, Tailwind CSS

📌 Contribuciones

¡Las contribuciones son bienvenidas! Abre un issue o envía un pull request.

📞 Contacto

Si tienes preguntas o sugerencias, contáctanos
