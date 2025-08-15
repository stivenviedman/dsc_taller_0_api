# dsc_taller_0_api
Taller de nivelacion del curso Desarrollo de Soluciones Cloud

# API de Tareas (Tasks)

Esta API permite gestionar tareas (`Tasks`) asociadas a usuarios y categorías. Está desarrollada con **Go** y **Fiber**.

**Base URL:** `http://localhost:8080/api`

---

## Endpoints

### 1. Obtener todos los tasks
- **URL:** `/tasks`
- **Método:** `GET`
- **Descripción:** Obtiene todas las tareas, incluyendo información de usuario y categoría asociada.
- **Respuesta exitosa:**
```json
{
  "message": "Se obtuvieron los tasks corretamente",
  "data": [
    {
      "ID": 1,
      "Description": "Tarea ejemplo",
      "State": "pendiente",
      "CreationDate": "2025-08-14T12:00:00Z",
      "FinalizationDate": null,
      "User": { "ID": 1, "Name": "Juan" },
      "Category": { "ID": 2, "Name": "Trabajo" }
    }
  ]
}
