package httpes

import (
	"souphakone-keop/go-learn-basic/httpes/crudapi"

	"github.com/gofiber/fiber/v2"
)

func Httpes() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	//CRUD
	app.Get("/book", crudapi.GetBooks)
	app.Get("/book/:id", crudapi.GetBooksId)
	/* 	app.Get("/book/:id", getBooks)
	   	app.Post("/book", createBooks)
	   	app.Put("/book", updateBook)
	   	app.Delete("/book", deleteBooks) */
	app.Listen(":8080")
}
