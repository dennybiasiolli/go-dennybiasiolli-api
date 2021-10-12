package main

import (
	"fmt"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/articoli"
	"github.com/dennybiasiolli/go-dennybiasiolli-api/auth"
	"github.com/dennybiasiolli/go-dennybiasiolli-api/budgest"
	"github.com/dennybiasiolli/go-dennybiasiolli-api/citazioni"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func setupFiberRoutes(app *fiber.App) {
	var db = make(map[string]string)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	basicAuthHandler := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"foo":  "bar", // user:foo password:bar
			"manu": "123", // user:manu password:123
		},
	})

	// // Authorized group (uses basicauth middleware)
	// authorized := app.Group("/", basicAuthHandler)

	app.Post("/admin", basicAuthHandler, func(c *fiber.Ctx) error {
		type BodyStruct struct {
			Value  string `json:"value" xml:"value" form:"value" validate:"required,min=3,max=32"`
			Value2 int    `json:"value2" xml:"value2" form:"value2" validate:"required,number"`
		}

		jsonBody := new(BodyStruct)
		if err := c.BodyParser(jsonBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err := validator.New().Struct(jsonBody)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
			// validation
			// response := []fiber.Map{}
			// for _, err := range err.(validator.ValidationErrors) {
			// 	response = append(response, fiber.Map{
			// 		"Tag":             err.Tag(),
			// 		"ActualTag":       err.ActualTag(),
			// 		"Namespace":       err.Namespace(),
			// 		"StructNamespace": err.StructNamespace(),
			// 		"Field":           err.Field(),
			// 		"StructField":     err.StructField(),
			// 		"Value":           err.Value(),
			// 		"Param":           err.Param(),
			// 		"Error":           err.Error(),
			// 	})
			// }
			// return c.Status(fiber.StatusBadRequest).JSON(response)
		}

		username := fmt.Sprintf("%v", c.Locals("username"))
		db[username] = jsonBody.Value
		return c.JSON(fiber.Map{"status": "ok"})
	})

	articoli.ArticoliAnonymousRegister(app.Group("/articoli"))
	citazioni.CitazioniAnonymousRegister(app.Group("/citazioni"))
	citazioni.CitazioneAnonymousRegister(app.Group("/citazione"))
	auth.JwtTokenRegister(app.Group("/token"))
	budgest.BudgestRegister(app.Group("/budgest").Use(auth.GetDjangoJwtAuthMiddleware()))
}
