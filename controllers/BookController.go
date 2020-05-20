package books

import (
	"awesomeProject/database"
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx)  {
	db := database.DBConn
	var books []database.Book
	db.Table("books").Find(&books)
	err := c.JSON(books)
	if err != nil {
		panic("Unable to get books")
	}
}

func GetBook(c *fiber.Ctx)  {
	db := database.DBConn
	var book database.Book
	db.Table("books").First(&book, c.Params("id"))
	if book.ID == 0 {
		panic("Book not found")
	}
	err := c.JSON(book)
	if err != nil {
		panic("Unable to get the book")
	}
}
