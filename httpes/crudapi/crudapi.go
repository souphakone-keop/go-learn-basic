package crudapi

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func init() {
	books = append(books, Book{
		Id:     1,
		Title:  "1900",
		Author: "sukumni",
	})

	books = append(books, Book{
		Id:     2,
		Title:  "English",
		Author: "phoy",
	})
}

// GET ALL BOOKS
func GetBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

// GET BOOK BY ID
func GetBooksId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	for _, book := range books {
		if book.Id == id {
			return c.JSON(book)
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"message": "Book not found",
	})
}