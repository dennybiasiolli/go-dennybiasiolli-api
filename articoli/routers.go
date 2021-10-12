package articoli

import "github.com/gofiber/fiber/v2"

func ArticoliAnonymousRegister(router fiber.Router) {
	router.Get("/", ArticoliList)
	router.Get("/:id/", ArticoliDetail)
}
