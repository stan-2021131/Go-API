# Go-API

API REST en **Go** para administrar mis videojuegos favoritos. Implementa operaciones CRUD sobre una base de datos SQLite, con soporte para filtros y manejo de errores JSON consistentes.  

Esta API fue desarrollada para el curso de **Sistemas y Tecnologías Web** y se ejecuta en el puerto asignado al carnet (`:24759`).  

---

## 📦 Estructura del proyecto

```

|-Database
|   |-db24759.db
|   |-DDL.sql
|   |-seed.sql
|-Evidencia
|   |-capturas_postman.pdf
|   |-Go-API.json
|-.gitignore
|-docker-compose.yml
|-docker-compose.yml.example
|-.Dockerfile
|-main.go
|-README.md

````

- `Database/` → Contiene la base de datos SQLite y scripts de creación/seed.  
- `Evidencia/` → Capturas de Postman y colección de pruebas.  
- `Dockerfile` y `docker-compose.yml` → Permiten ejecutar la API en contenedor.  
- `main.go` → Código principal de la API.  

---

## 🚀 Instalación y ejecución

### Con Docker

```bash
# Construir la imagen
docker build -t go-api .

# Levantar el contenedor usando docker-compose
docker-compose up -d
````

### Sin Docker (requiere Go instalado)

```bash
# Descargar dependencias
go mod tidy

# Ejecutar la API
go run main.go
```

La API estará disponible en:

```
http://localhost:24759/api/videojuegos/
```

---

## 🛠 Endpoints disponibles

Todos los endpoints responden con **JSON**.

| Método | URL                     | Parámetros                  | Descripción                                                                |
| ------ | ----------------------- | --------------------------- | -------------------------------------------------------------------------- |
| GET    | `/api/videojuegos/`     | `?genero=&publicado=`       | Obtener todos los videojuegos o filtrados por género y año de publicación. |
| GET    | `/api/videojuegos/{id}` | path param: `id`            | Obtener un videojuego por su ID.                                           |
| POST   | `/api/videojuegos/`     | Body JSON                   | Crear un nuevo videojuego.                                                 |
| PUT    | `/api/videojuegos/{id}` | path param: `id`, Body JSON | Actualizar un videojuego existente.                                        |
| DELETE | `/api/videojuegos/{id}` | path param: `id`            | Eliminar un videojuego por ID.                                             |

---

## 📄 Estructura del Body JSON

```json
{
  "nombre": "Nombre del juego",
  "publicado": 2000,
  "genero": "Acción",
  "plataformas": "PC, PS4",
  "desarrollador": "Nombre del desarrollador"
}
```

> Nota: No es necesario enviar el `id` en POST o PUT; se toma del path o se genera automáticamente.

---

## 🔹 Ejemplos de requests y responses

### GET todos los videojuegos

```bash
GET http://localhost:24759/api/videojuegos/
```

**Response 200 OK**

```json
[
  {
    "id": 1,
    "nombre": "Halo 3",
    "publicado": 2007,
    "genero": "Shooter",
    "plataformas": "Xbox 360",
    "desarrollador": "Bungie"
  },
  {
    "id": 2,
    "nombre": "Need for Speed Underground",
    "publicado": 2003,
    "genero": "Carreras",
    "plataformas": "PC, PS2, GameCube",
    "desarrollador": "EA Black Box"
  }
]
```

---

### GET videojuego por ID

```bash
GET http://localhost:24759/api/videojuegos/2
```

**Response 200 OK**

```json
{
  "id": 2,
  "nombre": "Need for Speed Underground",
  "publicado": 2003,
  "genero": "Carreras",
  "plataformas": "PC, PS2, GameCube",
  "desarrollador": "EA Black Box"
}
```

**Response 404 (no encontrado)**

```json
{
  "Message": "Videojuego no encontrado"
}
```

---

### POST (crear nuevo videojuego)

```bash
POST http://localhost:24759/api/videojuegos/
Content-Type: application/json

{
  "nombre": "FIFA 23",
  "publicado": 2022,
  "genero": "Deportes",
  "plataformas": "PC, PS5, Xbox",
  "desarrollador": "EA Sports"
}
```

**Response 200 OK**

```json
{
  "id": 11,
  "nombre": "FIFA 23",
  "publicado": 2022,
  "genero": "Deportes",
  "plataformas": "PC, PS5, Xbox",
  "desarrollador": "EA Sports"
}
```

**Response 400 (error en JSON)**

```json
{
  "Message": "Parametros faltantes"
}
```

---

### PUT (actualizar videojuego)

```bash
PUT http://localhost:24759/api/videojuegos/2
Content-Type: application/json

{
  "nombre": "Need for Speed Underground Remastered",
  "publicado": 2004,
  "genero": "Carreras Arcade",
  "plataformas": "PC, PS2, GameCube, Xbox",
  "desarrollador": "EA Black Box Studios"
}
```

**Response 200 OK**

```json
{
  "Message": "Videojuego actualizado correctamente"
}
```

---

### DELETE (eliminar videojuego)

```bash
DELETE http://localhost:24759/api/videojuegos/2
```

**Response 200 OK**

```json
{
  "Message": "Videojuego eliminado correctamente"
}
```

**Response 404 (no encontrado)**

```json
{
  "Message": "Videojuego no encontrado"
}
```

---

## 🧪 Evidencia de pruebas

Dentro de la carpeta `Evidencia/`:

* `capturas_postman.pdf` → capturas de todos los endpoints funcionando y errores.
* `Go-API.json` → colección de Postman con todos los requests listos para probar.

---

## ⚙️ Notas adicionales

* La API está diseñada para **portabilidad con Docker**, usando `docker-compose` para levantar la base de datos y la API.
* Todas las respuestas de error usan **JSON estructurado**, con mensajes claros.
* Se permiten **filtros combinados** en GET, usando query parameters `?genero=` y `?publicado=`.
