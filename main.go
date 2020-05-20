package main

import (
	books "awesomeProject/controllers"
	"awesomeProject/database"
	// "github.com/gofiber/basicauth"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
)

func main() {
	app := fiber.New()
	app.Settings.ServerHeader="Sky.Fiber"
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowMethods: []string{http.MethodGet},
		AllowOrigins: []string{"http://localhost"},
	}))
  	// app.Use(basicauth.New(basicauth.Config{
   //  	Users: map[string]string{
   //    		"admin":  "123456",
	  //   },
  	// }))

    app.Use(recover.New(recover.Config{
        Handler: func(c *fiber.Ctx, err error) {
            c.SendString(err.Error())
            c.SendStatus(500)
        },
    }))

	database.DBConn = database.OpenDB()
	defer database.DBConn.Close()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Books API")
	})

	api:= app.Group("/api")

	api.Get("/books", books.GetBooks)
	api.Get("/books/:id", books.GetBook)

	err := app.Listen(4000)
	if err != nil {
		log.Println("Error: unable to start listening. Why: " + err.Error())
	}
}
