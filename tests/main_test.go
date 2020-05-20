package tests

import (
	books "awesomeProject/controllers"
	"awesomeProject/database"
	"github.com/gofiber/fiber"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func dbSetup() {
	database.DBConn = database.OpenDB()
	defer database.DBConn.Close()
}

func testStatus200(t *testing.T, app *fiber.App, url string, method string) {
	req := httptest.NewRequest(method, url, nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 200, resp.StatusCode, "Status code")
}

func TestAllBooks(t *testing.T) {
	app := fiber.New()
	dbSetup()
	time.Sleep(3 * time.Second)
	app.Get("/api/books", books.GetBooks)
	testStatus200(t, app, "/api/books", "GET")
}