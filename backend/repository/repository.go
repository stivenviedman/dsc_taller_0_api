package repository

import (
	"back-end-todolist/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
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

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Task routes
	api.Post("/create_tasks", r.CreateTask)
	api.Put("/update_task/:id", r.UpdateTask)
	api.Delete("/delete_task/:id", r.DeleteTask)
	api.Get("/get_tasks/:id", r.GetTaskById)
	api.Get("/tasks", r.GetTasks)
	//pendiente endpoint para buscar tasks por categoria y/o por estado

	// User routes
	api.Post("/create_users", r.CreateUser)
	api.Get("/tasks/:userId", r.GetTasksByUserId)

	// Category routes
	api.Post("/categorias", r.CreateCategory)
	api.Get("/categorias", r.GetCategories)
	api.Delete("/categorias/:id", r.DeleteCategory)
}
