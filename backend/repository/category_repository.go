package repository

import (
	"back-end-todolist/models"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*---Category functions----*/
func (r *Repository) CreateCategory(context *fiber.Ctx) error {

	userID := context.Locals("userID").(uint)

	category := models.Category{}

	err := context.BodyParser(&category)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// Validar que el nombre no esté vacío
	if category.Name == "" {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "El nombre de la categoría no puede estar vacío"},
		)
	}

	// Validar que la descripción no esté vacía
	if category.Description == "" {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "La descripción de la categoría no puede estar vacía"},
		)
	}

	// Verificar si ya existe una categoría con ese nombre
	var existingCategory models.Category
	if err := r.DB.Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
		return context.Status(http.StatusConflict).JSON(
			&fiber.Map{"message": "Ya existe una categoría con ese nombre"},
		)
	}

	/*Le asigna el userID*/
	category.UserID = userID

	errCreate := r.DB.Create(&category).Error
	if errCreate != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo crear la categoría"})
		return errCreate
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se creó la categoría correctamente", "data": category})
	return nil
}

func (r *Repository) GetCategories(context *fiber.Ctx) error {

	userID := context.Locals("userID").(uint)

	categoryModels := &[]models.Category{}

	err := r.DB.Where("user_id", userID).Find(categoryModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudieron obtener las categorías"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvieron las categorías del usuario correctamente",
		"data":    categoryModels,
	})
	return nil
}

func (r *Repository) DeleteCategory(context *fiber.Ctx) error {
	categoryModel := models.Category{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacío"})
		return nil
	}

	// Verificar si la categoría existe
	if err := r.DB.First(&categoryModel, id).Error; err != nil {
		return context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "Categoría no encontrada"},
		)
	}

	// Verificar si es la categoría "Sin Categoría" (ID 1) - no permitir eliminarla
	if id == "1" {
		return context.Status(http.StatusForbidden).JSON(
			&fiber.Map{"message": "No se puede eliminar la categoría 'Sin Categoría'"},
		)
	}

	// Verificar si hay tareas usando esta categoría
	var taskCount int64
	if err := r.DB.Model(&models.Task{}).Where("category_id = ?", id).Count(&taskCount).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error verificando tareas de la categoría"})
		return err
	}

	// Si hay tareas, moverlas a la categoría "Sin Categoría" (ID 1)
	if taskCount > 0 {
		// Actualizar todas las tareas de esta categoría para que usen la categoría "Sin Categoría"
		if err := r.DB.Model(&models.Task{}).Where("category_id = ?", id).Update("category_id", 1).Error; err != nil {
			context.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{"message": "Error moviendo tareas a categoría 'Sin Categoría'"})
			return err
		}
		log.Printf("Se movieron %d tareas a la categoría 'Sin Categoría'", taskCount)
	}

	// Ahora eliminar la categoría
	err := r.DB.Delete(&categoryModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo eliminar la categoría"})
		return err.Error
	}

	// Mensaje personalizado según si se movieron tareas o no
	if taskCount > 0 {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{
				"message":        "Se eliminó la categoría correctamente",
				"info":           "Se movieron tareas a la categoría 'Sin Categoría'",
				"tareas_movidas": taskCount,
			})
	} else {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "Se eliminó la categoría correctamente"})
	}
	return nil
}
