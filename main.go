package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gofiber/fiber/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	common.GetEnvVariables()

	common.ConnectDb()
	// common.GetDB().AutoMigrate(&citazioni.Citazione{})

	app := fiber.New()

	setupFiberRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(common.HTTP_LISTEN))
}
