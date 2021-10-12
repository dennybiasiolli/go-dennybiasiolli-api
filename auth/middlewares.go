package auth

import (
	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func GetDjangoBasicAuthMiddleware() func(*fiber.Ctx) error {
	return basicauth.New(basicauth.Config{
		Authorizer: func(username, password string) bool {
			_, err := LoginDjangoUser(username, password)
			if err != nil {
				return false
			}
			return true
		},
	})
}

func GetDjangoJwtAuthMiddleware() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(common.JWT_HMAC_SAMPLE_SECRET),
		SuccessHandler: func(c *fiber.Ctx) error {
			u := c.Locals("user").(*jwt.Token)
			claims := u.Claims.(jwt.MapClaims)

			if claims["token_type"] != "access" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token type"})
			}

			db := common.GetDB()
			var user User = User{
				IsActive: true,
			}
			err := db.Where(&user).First(&user, claims["user_id"]).Error
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
			}
			c.Locals("user", user)
			return c.Next()
		},
	})
}
