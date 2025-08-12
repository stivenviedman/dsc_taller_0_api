package repository

import (
	"back-end-todolist/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*---Category functions----*/
func (r *Repository) CreateCategory(context *fiber.Ctx) error {
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

	// Verificar si ya existe una categoría con ese nombre
	var existingCategory models.Category
	if err := r.DB.Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
		return context.Status(http.StatusConflict).JSON(
			&fiber.Map{"message": "Ya existe una categoría con ese nombre"},
		)
	}

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
	categoryModels := &[]models.Category{}

	err := r.DB.Find(categoryModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudieron obtener las categorías"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvieron las categorías correctamente",
		"data":    categoryModels,
	})
	return nil
}

func (r *Repository) GetCategoryById(context *fiber.Ctx) error {
	categoryModel := &models.Category{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacío"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(categoryModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo obtener la categoría por id"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvo la categoría correctamente",
		"data":    categoryModel,
	})
	return nil
}

func (r *Repository) UpdateCategory(context *fiber.Ctx) error {
	category := models.Category{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacío"})
		return nil
	}

	// Verificar si la categoría existe
	if err := r.DB.First(&category, id).Error; err != nil {
		return context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "Categoría no encontrada"},
		)
	}

	// Parsear el body
	var updateData models.Category
	if err := context.BodyParser(&updateData); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// Validar que el nombre no esté vacío
	if updateData.Name == "" {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "El nombre de la categoría no puede estar vacío"},
		)
	}

	// Verificar si ya existe otra categoría con ese nombre (excluyendo la actual)
	var existingCategory models.Category
	if err := r.DB.Where("name = ? AND id != ?", updateData.Name, id).First(&existingCategory).Error; err == nil {
		return context.Status(http.StatusConflict).JSON(
			&fiber.Map{"message": "Ya existe otra categoría con ese nombre"},
		)
	}

	// Actualizar solo el nombre
	category.Name = updateData.Name

	errUpdate := r.DB.Save(&category).Error
	if errUpdate != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo actualizar la categoría"})
		return errUpdate
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se actualizó la categoría correctamente", "data": category})
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

	// Verificar si hay tareas usando esta categoría
	var taskCount int64
	if err := r.DB.Model(&models.Task{}).Where("category_id = ?", id).Count(&taskCount).Error; err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error verificando tareas de la categoría"})
		return err
	}

	if taskCount > 0 {
		return context.Status(http.StatusConflict).JSON(
			&fiber.Map{"message": "No se puede eliminar la categoría porque tiene tareas asociadas"},
		)
	}

	err := r.DB.Delete(&categoryModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo eliminar la categoría"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se eliminó la categoría correctamente"})
	return nil
}

func (r *Repository) GetTasksByCategory(context *fiber.Ctx) error {
	categoryID := context.Params("id")
	if categoryID == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID de la categoría no puede estar vacío"})
		return nil
	}

	// Verificar si la categoría existe
	var category models.Category
	if err := r.DB.First(&category, categoryID).Error; err != nil {
		return context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "Categoría no encontrada"},
		)
	}

	// Obtener tareas de la categoría
	var tasks []models.Task
	if err := r.DB.Where("category_id = ?", categoryID).Find(&tasks).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudieron obtener las tareas de la categoría"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":  "Se obtuvieron las tareas de la categoría correctamente",
		"data":     tasks,
		"category": category,
	})
	return nil
}
