package auth

import (
	"strings"
	"time"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func TokenObtain(c *fiber.Ctx) error {
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validator.New().Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := LoginDjangoUser(input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Set claims
	claims := JwtCustomClaims{
		"access",
		user.ID,
		JwtUserInfo{
			Username: user.Username,
			Email:    user.Email,
			FullName: strings.TrimSpace(user.FirstName + " " + user.LastName),
			IsStaff:  user.IsStaff,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(common.JWT_ACCESS_TOKEN_LIFETIME_SECONDS),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(common.JWT_HMAC_SAMPLE_SECRET))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"access": t})
}
