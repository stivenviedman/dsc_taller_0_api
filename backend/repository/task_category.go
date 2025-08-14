package repository

import (
	"back-end-todolist/models"
	"net/http"

	"time"

	"github.com/gofiber/fiber/v2"
)

/*---Task functions----*/
func (r *Repository) CreateTask(context *fiber.Ctx) error {

	task := models.Task{}

	err := context.BodyParser(&task)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
	}

	// Asignar la fecha actual
	task.CreationDate = time.Now()

	// Validar que el User existe
	user := models.User{}
	if err := r.DB.First(&user, task.UserID).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Usuario no encontrado"},
		)
	}

	// Validar que la categoria existe
	category := models.Category{}
	if err := r.DB.First(&category, task.CategoryID).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Categoria no encontrada"},
		)
	}

	errCreate := r.DB.Create(&task).Error

	if errCreate != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo crear el task"})

		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se creo el task correctamente"})

	return nil
}

func (r *Repository) UpdateTask(context *fiber.Ctx) error {

	task := models.Task{}

	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacio"})

		return nil
	}

	err := r.DB.Where("id = ?", id).First(&task).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo obtener el task por id"})

		return err
	}

	taskDTO := models.Task{}

	errParser := context.BodyParser(&taskDTO)

	if errParser != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
	}

	task.Description = taskDTO.Description
	task.State = taskDTO.State
	task.FinalizationDate = taskDTO.FinalizationDate
	task.CategoryID = taskDTO.CategoryID
	task.UserID = taskDTO.UserID

	// Guardar cambios
	if err := r.DB.Save(&task).Error; err != nil {
		return context.Status(http.StatusInternalServerError).JSON(
			fiber.Map{"message": "No se pudo actualizar el task"},
		)
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se actualizo el task correctamente"})

	return nil
}

func (r *Repository) DeleteTask(context *fiber.Ctx) error {

	taskModel := models.Task{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacio"})

		return nil
	}

	err := r.DB.Delete(taskModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo eliminar el task"})

		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se elimino el task corretamente"})

	return nil
}

func (r *Repository) GetTaskById(context *fiber.Ctx) error {

	task := &models.Task{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacio"})

		return nil
	}

	if err := r.DB.
		Preload("User").
		Preload("Category").
		Where("id = ?", id).
		First(task).Error; err != nil {

		return context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error al obtener el task"},
		)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvo el task corretamente",
		"data":    task,
	})

	return nil
}

func (r *Repository) GetTasks(context *fiber.Ctx) error {

	tasks := &[]models.Task{}

	if err := r.DB.
		Preload("User").
		Preload("Category").
		Find(&tasks).Error; err != nil {

		return context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error al obtener los tasks"},
		)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvieron los tasks corretamente",
		"data":    tasks,
	})

	return nil
}

func (r *Repository) GetTasksByUserId(context *fiber.Ctx) error {

	userID := context.Params("userId")

	tasks := &[]models.Task{}

	if err := r.DB.
		Preload("User").
		Preload("Category").
		Where("user_id = ?", userID).
		Find(&tasks).Error; err != nil {

		return context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "Error al obtener los tasks"},
		)
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvieron los tasks corretamente del usuario",
		"data":    tasks,
	})

	return nil
}
