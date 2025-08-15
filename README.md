# dsc_taller_0_api
Taller de nivelacion del curso Desarrollo de Soluciones Cloud

# API de Tareas (Tasks)

Esta API permite gestionar tareas (`Tasks`) asociadas a usuarios y categorías. Está desarrollada con **Go** y **Fiber**.

**Base URL:** `http://localhost:8080/api`

---

## Endpoints

### 1. Obtener todos los tasks por UserId
- **URL:** `/tasks/{UserId}`
- **Método:** `GET`
- **Descripción:** Obtiene todas las tareas dado el id de un usuario, incluyendo información de usuario y categoría asociada.
- **Respuesta exitosa:**
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
- **URL:** `/get_tasks/{TaskId}`
- **Método:** `GET`
- **Descripción:** Obtiene todas las tareas dado el id de una tarea, incluyendo información de usuario y categoría asociada.
- **Respuesta exitosa:**
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
- **Método:** `GET`
- **Cuerpo de la solicitud (JSON):**
```json
{
    "description": "Sacar al perro",
    "finalizationDate": "2027-02-01T00:00:00Z",
    "state": "Sin Empezar",
    "category_id" : 3,
    "user_id": 1
}
```
- **Descripción:** Obtiene todas las tareas dado el id de una tarea, incluyendo información de usuario y categoría asociada.
- **Respuesta exitosa:**
```json
{
    "message": "Se creo el task correctamente"
}
```




