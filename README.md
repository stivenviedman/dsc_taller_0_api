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
            "state": "Sin Empezar",
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
        }
    ],
    "message": "Se obtuvieron los tasks corretamente del usuario"
}
```
### 1. Obtener todos los tasks por UserId
- **URL:** `/tasks/{UserId}`
- **Método:** `GET`
- **Descripción:** Obtiene todas las tareas dado el id de un usuario, incluyendo información de usuario y categoría asociada.
- **Respuesta exitosa:**

