package repository

import (
	"back-end-todolist/middlewares"
	"back-end-todolist/models"
	"fmt"
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

	dbuser := models.User{}
	errSelect := r.DB.Where("username = ?", user.Username).First(&dbuser).Error

	if errSelect != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo encontrar el usuario"})

		return errSelect
	}

	token, errToken := middlewares.GenerarToken(*user.Username, dbuser.ID)

	if errToken != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo generar el token"})
		return errToken
	}

	return context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se creo el user correctamente",
			"token": token})
}

func (r *Repository) LoginUser(context *fiber.Ctx) error {

	user := models.User{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Datos inválidos"})

		return err
	}

	dbuser := models.User{}
	errSelect := r.DB.Where("username = ?", user.Username).First(&dbuser).Error

	if errSelect != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo encontrar el usuario"})

		return errSelect
	}
	fmt.Printf("Usuario: %s\n", *user.Password)
	fmt.Printf("Usuario DB %s\n", *dbuser.Password)
	if *dbuser.Password != *user.Password {
		return context.Status(http.StatusForbidden).JSON(
			&fiber.Map{"message": "Contraseña incorrecta"})
	}

	token, errToken := middlewares.GenerarToken(*dbuser.Username, dbuser.ID)

	if errToken != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo generar el token"})

		return errToken
	}

	return context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Ingreso exitoso",
			"token": token,
			"image": dbuser.ImageP})
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Task routes
	api.Post("/create_tasks", middlewares.AutValidation, r.CreateTask)
	api.Put("/update_task/:id", middlewares.AutValidation, r.UpdateTask)
	api.Delete("/delete_task/:id", middlewares.AutValidation, r.DeleteTask)
	api.Get("/get_tasks/:id", middlewares.AutValidation, r.GetTaskById)
	api.Get("/tasks", middlewares.AutValidation, r.GetTasks)
	//pendiente endpoint para buscar tasks por categoria y/o por estado

	// User routes
	api.Post("/create_users", r.CreateUser)
	api.Post("/login_users", r.LoginUser)
	//Pendiente endpoint de obtener tasks by user id
	api.Get("/user_tasks", middlewares.AutValidation, r.GetTasksByUserId)

	// Category routes
	api.Post("/categorias", middlewares.AutValidation, r.CreateCategory)
	api.Get("/categorias", middlewares.AutValidation, r.GetCategories)
	api.Delete("/categorias/:id", middlewares.AutValidation, r.DeleteCategory)
}
