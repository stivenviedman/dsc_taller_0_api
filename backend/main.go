package main

import (
	"back-end-todolist/models"
	"back-end-todolist/storage"
	"log"
	"os"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Task struct {
	Description string `json:"description"`
	Date        string `json:"date"`
}

type Repository struct {
	DB *gorm.DB
}

/*---Repository functions----*/
func (r *Repository) CreateTask(context *fiber.Ctx) error {

	task := Task{}

	err := context.BodyParser(&task)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
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

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_tasks", r.CreateTask)
	api.Delete("/delete_task/:id", r.DeleteTask)
	api.Get("/get_tasks/:id", r.GetTaskById)
	api.Get("/tasks", r.GetTasks)
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Error cargando la base de datos")
	}

	errMigrate := models.MigrateTasks(db)

	if errMigrate != nil {
		log.Fatal("Error migrando la base de datos")
	}

	r := Repository{DB: db}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
