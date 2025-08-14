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

/*---Task functions----*/
func (r *Repository) CreateTask(context *fiber.Ctx) error {

	task := models.Task{}

	err := context.BodyParser(&task)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
	}

	// Validar que el User existe
	user := models.User{}
	if err := r.DB.First(&user, task.UserID).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Usuario no encontrado"},
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

	token, errToken := middlewares.GenerarToken(*user.Username)

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

	token, errToken := middlewares.GenerarToken(*dbuser.Username)

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
	api.Post("/create_tasks", r.CreateTask)
	//Pendiente endpoint de update task
	api.Delete("/delete_task/:id", r.DeleteTask)
	api.Get("/get_tasks/:id", r.GetTaskById)
	api.Get("/tasks", middlewares.AutValidation, r.GetTasks)
	api.Post("/create_users", r.CreateUser)
	api.Post("/login_users", r.LoginUser)
	//Pendiente endpoint de obtener tasks by user id
}
