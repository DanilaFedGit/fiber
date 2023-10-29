package handler

import (
	"fmt"
	"github.com/DanilaFedGit/fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
)

type Repository struct {
	DataBase *gorm.DB
}

func (r *Repository) GetBookById(contex *fiber.Ctx) error {
	bookModel := &[]models.Books{}
	id := contex.Params("id")
	if id == "" {
		contex.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "id is empty"})
		return nil
	}
	fmt.Println(id)
	err := r.DataBase.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		contex.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "can't find this book"})
		return nil
	}
	contex.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "get book by id",
			"data": bookModel})
	return nil
}
func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := &[]models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"mesage": "id is empty"})
		return nil
	}
	err := r.DataBase.Delete(bookModel, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "couldn't delete this book"})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "delete success"})
	return nil
}
func (r *Repository) GetBook(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DataBase.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "couldn't get books"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "books fetched successfully",
			"data": bookModels})
	return nil
}
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request was failed"})
		return err
	}
	err = r.DataBase.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "the book wasn't created"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "the book was created"})
	return nil
}
func (r *Repository) SetupRouters(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookById)
	api.Get("/books", r.GetBook)
}

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}
