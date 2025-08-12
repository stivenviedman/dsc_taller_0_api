# API de Categorías - TodoList

## Endpoints de Categorías

### 1. Crear Categoría
```
POST /api/categories
Content-Type: application/json

{
    "name": "Nombre de la categoría"
}
```

**Respuesta exitosa (200):**
```json
{
    "message": "Se creó la categoría correctamente",
    "data": {
        "id": 1,
        "name": "Nombre de la categoría"
    }
}
```

### 2. Obtener Todas las Categorías
```
GET /api/categories
```

**Respuesta exitosa (200):**
```json
{
    "message": "Se obtuvieron las categorías correctamente",
    "data": [
        {
            "id": 1,
            "name": "Personal"
        },
        {
            "id": 2,
            "name": "Trabajo"
        }
    ]
}
```

### 3. Obtener Categoría por ID
```
GET /api/categories/:id
```

**Respuesta exitosa (200):**
```json
{
    "message": "Se obtuvo la categoría correctamente",
    "data": {
        "id": 1,
        "name": "Personal"
    }
}
```

### 4. Actualizar Categoría
```
PUT /api/categories/:id
Content-Type: application/json

{
    "name": "Nuevo nombre de la categoría"
}
```

**Respuesta exitosa (200):**
```json
{
    "message": "Se actualizó la categoría correctamente",
    "data": {
        "id": 1,
        "name": "Nuevo nombre de la categoría"
    }
}
```

### 5. Eliminar Categoría
```
DELETE /api/categories/:id
```

**Respuesta exitosa (200):**
```json
{
    "message": "Se eliminó la categoría correctamente"
}
```

**Error (409) - Categoría tiene tareas:**
```json
{
    "message": "No se puede eliminar la categoría porque tiene tareas asociadas"
}
```

### 6. Obtener Tareas por Categoría
```
GET /api/categories/:id/tasks
```

**Respuesta exitosa (200):**
```json
{
    "message": "Se obtuvieron las tareas de la categoría correctamente",
    "data": [
        {
            "id": 1,
            "description": "Descripción de la tarea",
            "date": "2024-01-15",
            "user_id": 1,
            "category_id": 1
        }
    ],
    "category": {
        "id": 1,
        "name": "Personal"
    }
}
```

## Categorías por Defecto

Al iniciar la aplicación, se crean automáticamente las siguientes categorías:
- Personal
- Trabajo
- Estudio
- Hogar
- Salud

## Validaciones

- **Nombre único**: No se pueden crear categorías con nombres duplicados
- **Nombre obligatorio**: El nombre de la categoría no puede estar vacío
- **Integridad referencial**: No se puede eliminar una categoría que tenga tareas asociadas
- **Categoría obligatoria**: Toda tarea debe tener una categoría asignada

## Códigos de Estado HTTP

- **200**: Operación exitosa
- **400**: Bad Request (datos inválidos)
- **404**: Categoría no encontrada
- **409**: Conflicto (nombre duplicado o categoría con tareas)
- **422**: Unprocessable Entity (error en el parsing del body)