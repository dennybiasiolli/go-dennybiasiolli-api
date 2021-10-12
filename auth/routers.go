package auth

import "github.com/gofiber/fiber/v2"

func JwtTokenRegister(router fiber.Router) {
	router.Post("/", TokenObtain)
}
