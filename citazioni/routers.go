package citazioni

import (
	"github.com/gofiber/fiber/v2"
)

func CitazioniAnonymousRegister(router fiber.Router) {
	router.Get("/", CitazioniList)
	router.Get("/:id/", CitazioneDetail)
	router.Post("/", CitazioneCreate)
}
func CitazioneAnonymousRegister(router fiber.Router) {
	router.Get("/", CitazioneRandomDetail)
}
