# Taller de Nivelacion - Desarrollo de Soluciones Cloud
Taller de nivelacion del curso Desarrollo de Soluciones Cloud

Este repositorio include todo el codigo de la app full-stack Todo. El backend desarrollado en Go se encuentra en la carpeta `backend` y el frontend
desarrollado en Next.js se encuentra en la carpeta `./ui`

# Steps to run project 🚀
- En la carpeta raiz, ejecutar el comando: `docker compose up --build`, este es el unico comando necesario para ejecutar el proyecto, incluye
los contenedores para la base de datos PostgreSQL, el backend y la interfaz grafica.

El backend es accesible por medio de `http://localhost:8080/api`, mientras que el frontend es accesible a traves de ``http://localhost:3000``

# API Endpoints Documentation

Esta API permite gestionar tareas (`Tasks`) asociadas a usuarios y categorías. Está desarrollada con **Go** y **Fiber**.

**Base URL:** `http://localhost:8080/api`

---

## Endpoints

### 1. Obtener todos los tasks por UserId
- **URL:** `/user_tasks/`
- **Método:** `GET`
- **Descripción:** Obtiene todas las tareas dado el id de un usuario (el userId se obtiene por medio del JWT Token), incluyendo información de usuario y categoría asociada.
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "data": [
        {
            "id": 1,
            "description": "Hacer mercado",
            "creationDate": "2025-08-14T00:00:00Z",
            "finalizationDate": "2027-02-01T00:00:00Z",
            "state": "pendiente",
            "user_id": 1,
            "User": {
                "id": 1,
                "username": "gabo98",
                "password": "1234"
            },
            "category_id": 1,
            "Category": {
                "id": 1,
                "name": "Sin Categoría",
                "description": "Tareas sin categoría"
            }
        },
        {
            "id": 2,
            "description": "Sacar al perro",
            "creationDate": "2025-08-15T00:00:00Z",
            "finalizationDate": "2027-02-01T00:00:00Z",
            "state": "Sin Empezar",
            "user_id": 1,
            "User": {
                "id": 1,
                "username": "gabo98",
                "password": "1234"
            },
            "category_id": 3,
            "Category": {
                "id": 3,
                "name": "Trabajo",
                "description": "Tareas relacionadas con el trabajo y proyectos laborales"
            }
        }
    ],
    "message": "Se obtuvieron los tasks corretamente del usuario"
}
```
### 2. Obtener task por id
- **URL:** `/get_tasks/{taskId}`
- **Método:** `GET`
- **Descripción:** Obtiene una tarea dado el id, incluyendo información de usuario y categoría asociada.
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "data": {
        "id": 1,
        "description": "Hacer mercado",
        "creationDate": "2025-08-14T00:00:00Z",
        "finalizationDate": "2027-02-01T00:00:00Z",
        "state": "pendiente",
        "user_id": 1,
        "User": {
            "id": 1,
            "username": "gabo98",
            "password": "1234"
        },
        "category_id": 1,
        "Category": {
            "id": 1,
            "name": "Sin Categoría",
            "description": "Tareas sin categoría"
        }
    },
    "message": "Se obtuvo el task corretamente"
}
```
### 3. Crear task
- **URL:** `/create_tasks`
- **Método:** `POST`
- **Cuerpo de la solicitud (JSON):**
```json
{
    "description": "Sacar al perro",
    "finalizationDate": "2027-02-01T00:00:00Z",
    "state": "Sin Empezar",
    "category_id" : 3
}
```
- **Descripción:** Crea una tarea, asociandola a un usuario y a una categoria.
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "message": "Se creo el task correctamente"
}
```
### 4. Update task
- **URL:** `/update_task/{taskId}`
- **Método:** `PUT`
- **Cuerpo de la solicitud (JSON):**
```json
{
    "id":2,
    "description": "Sacar al perro pero al parque grande",
    "finalizationDate": "2027-04-01T00:00:00Z",
    "state": "Completado",
    "category_id" : 1
}
```
- **Descripción:** Actualiza una tarea existente, dado el id de la tarea y el cuerpo json con los datos nuevos.
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "message": "Se actualizo el task correctamente"
}
```
### 5. Delete task
- **URL:** `/delete_task/{taskId}`
- **Método:** `DELETE`
- **Descripción:** Elimina una tarea dado su id.
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "message": "Se elimino el task corretamente"
}
```

---

## Endpoints de Categorías

### 1. Crear categoría
- **URL:** `/categorias`
- **Método:** `POST`
- **Autenticación:** Requiere JWT Token (Bearer Token)
- **Cuerpo de la solicitud (JSON):**
```json
{
    "name": "Trabajo",
    "description": "Tareas relacionadas con el trabajo y proyectos laborales"
}
```
- **Descripción:** Crea una nueva categoría asociada al usuario autenticado. El nombre debe ser único.
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "message": "Se creó la categoría correctamente",
    "data": {
        "id": 3,
        "name": "Trabajo",
        "description": "Tareas relacionadas con el trabajo y proyectos laborales",
        "user_id": 1
    }
}
```
- **Errores posibles:**
  - **400 Bad Request:** Nombre o descripción vacíos
  - **409 Conflict:** Ya existe una categoría con ese nombre
  - **401 Unauthorized:** Token JWT inválido o faltante

### 2. Obtener categorías del usuario
- **URL:** `/categorias`
- **Método:** `GET`
- **Autenticación:** Requiere JWT Token (Bearer Token)
- **Descripción:** Obtiene todas las categorías del usuario autenticado.
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "message": "Se obtuvieron las categorías del usuario correctamente",
    "data": [
        {
            "id": 1,
            "name": "Sin Categoría",
            "description": "Tareas sin categoría",
            "user_id": 1
        },
        {
            "id": 2,
            "name": "Personal",
            "description": "Tareas personales y del hogar",
            "user_id": 1
        },
        {
            "id": 3,
            "name": "Trabajo",
            "description": "Tareas relacionadas con el trabajo y proyectos laborales",
            "user_id": 1
        }
    ]
}
```
- **Errores posibles:**
  - **400 Bad Request:** Error al obtener las categorías
  - **401 Unauthorized:** Token JWT inválido o faltante

### 3. Eliminar categoría
- **URL:** `/categorias/{categoryId}`
- **Método:** `DELETE`
- **Autenticación:** Requiere JWT Token (Bearer Token)
- **Descripción:** Elimina una categoría específica. Si la categoría tiene tareas asociadas, estas se mueven automáticamente a la categoría "Sin Categoría" (ID 1). No se puede eliminar la categoría "Sin Categoría".
- **Respuesta exitosa STATUS 200 OK:**
```json
{
    "message": "Se eliminó la categoría correctamente"
}
```
- **Respuesta cuando se mueven tareas:**
```json
{
    "message": "Se eliminó la categoría correctamente",
    "info": "Se movieron tareas a la categoría 'Sin Categoría'",
    "tareas_movidas": 3
}
```
- **Errores posibles:**
  - **400 Bad Request:** Error al eliminar la categoría
  - **403 Forbidden:** Intento de eliminar la categoría "Sin Categoría"
  - **404 Not Found:** Categoría no encontrada
  - **401 Unauthorized:** Token JWT inválido o faltante
  - **500 Internal Server Error:** ID de categoría vacío o error al verificar tareas

### Notas importantes sobre categorías:
- **Categoría "Sin Categoría":** Esta categoría (ID 1) es especial y no se puede eliminar
- **Unicidad:** Los nombres de categorías deben ser únicos por usuario
- **Tareas asociadas:** Al eliminar una categoría, las tareas se mueven automáticamente a "Sin Categoría"
- **Autenticación:** Todos los endpoints de categorías requieren un token JWT válido
- **Validaciones:** Se valida que el nombre y descripción no estén vacíos


