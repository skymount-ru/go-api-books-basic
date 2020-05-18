package main

import (
	books "awesomeProject/controllers"
	"awesomeProject/database"
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/jinzhu/gorm"
)

func main() {
	app := fiber.New()
	app.Settings.ServerHeader="Sky.Fiber"
	app.Use(logger.New())

	var err error

	database.DBConn, err = gorm.Open("mysql", "db-config")
	if err != nil {
		panic("There was an error opening the database!")
	}
	fmt.Println("Connected to database successfully!")
	database.DBConn.AutoMigrate(&books.Book{})
	defer database.DBConn.Close()

	app.Static("/", "/public")
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Books API")
	})
	app.Get("/books", books.GetBooks)
	app.Get("/books/:id", books.GetBook)
	app.Listen(4000)
}
