package repository

import (
	"back-end-todolist/models"
	"net/http"

	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

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
	category := models.CategoryTemp{}
	if err := r.DB.First(&category, task.CategoryTempID).Error; err != nil {
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

	taskModel := &models.Task{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacio"})

		return nil
	}

	err := r.DB.Where("id = ?", id).First(taskModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo obtener el task por id"})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvo el task corretamente",
		"data":    taskModel,
	})

	return nil
}

func (r *Repository) GetTasks(context *fiber.Ctx) error {

	taskModels := &[]models.Task{}

	err := r.DB.Find(taskModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudieron obtener los tasks"})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvieron los tasks corretamente",
		"data":    taskModels,
	})

	return nil
}

func (r *Repository) GetTasksByUserId(context *fiber.Ctx) error {

	userID := context.Params("userId")

	tasks := &[]models.Task{}

	// Traer tasks + user + category
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

/*---User functions----*/
func (r *Repository) CreateUser(context *fiber.Ctx) error {

	user := models.User{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
	}

	errCreate := r.DB.Create(&user).Error

	if errCreate != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo crear el user"})

		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se creo el user correctamente"})

	return nil
}

/*---Category functions----*/
func (r *Repository) CreateCategory(context *fiber.Ctx) error {

	category := models.CategoryTemp{}

	err := context.BodyParser(&category)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
	}

	errCreate := r.DB.Create(&category).Error

	if errCreate != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo crear la category"})

		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se creo la category correctamente"})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_tasks", r.CreateTask)
	//Pendiente endpoint de update task
	api.Delete("/delete_task/:id", r.DeleteTask)
	api.Get("/get_tasks/:id", r.GetTaskById)
	api.Get("/tasks", r.GetTasks)
	api.Post("/create_users", r.CreateUser)
	api.Get("/tasks/:userId", r.GetTasksByUserId)
	api.Post("/create_categories", r.CreateCategory)

}
