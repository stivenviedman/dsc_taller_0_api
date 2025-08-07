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

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

/*---Repository functions----*/
func (r *Repository) CreateBook(context *fiber.Ctx) error {

	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})

		return err
	}

	errCreate := r.DB.Create(&book).Error

	if errCreate != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo crear el book"})

		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se creo el book corretamente"})

	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {

	bookModel := models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacio"})

		return nil
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo eliminar el book"})

		return err.Error
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Se elimino el book corretamente"})

	return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error {

	bookModel := &models.Books{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "El ID no puede estar vacio"})

		return nil
	}

	err := r.DB.Where("id = ?", id).First(bookModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudo obtener el book por id"})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvo el book corretamente",
		"data":    bookModel,
	})

	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {

	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "No se pudieron obtener los libros"})

		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Se obtuvieron los libros corretamente",
		"data":    bookModels,
	})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookById)
	api.Get("/books", r.GetBooks)
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

	errMigrate := models.MigrateBooks(db)

	if errMigrate != nil {
		log.Fatal("Error migrando la base de datos")
	}

	r := Repository{DB: db}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
