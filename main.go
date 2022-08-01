package main
// create gofiber application
// Language: go
// Path: main.go
import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World cokkk!")
	})
	app.Listen(":8080")
}
