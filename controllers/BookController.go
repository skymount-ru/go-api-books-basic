package books

import (
	"awesomeProject/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name string
	Country string
	Released string
}

func GetBooks(c *fiber.Ctx)  {
	db := database.DBConn
	var books []Book
	db.Table("books").Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx)  {
	db := database.DBConn
	var book Book
	db.Table("books").First(&book, c.Params("id"))
	c.JSON(book)
}
