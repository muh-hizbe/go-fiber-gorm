package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"log"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//	Validation Required Image
	var filenameString string

	filename := ctx.Locals("filename")
	log.Println("filename = ", filename)
	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image cover is required.",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newBook,
	})
}
