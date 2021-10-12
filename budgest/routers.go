package budgest

import "github.com/gofiber/fiber/v2"

func BudgestRegister(router fiber.Router) {
	router.Get("/ambiti/", AmbitiList)
	router.Get("/ambiti/:id/", AmbitiDetail)
}
