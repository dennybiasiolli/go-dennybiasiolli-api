package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func getAuthorizationHeader(c *gin.Context, auth_type string) (string, error) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	authParts := strings.SplitN(authorizationHeader, " ", 2)
	if len(authParts) != 2 {
		return "", errors.New("Failed to parse authentication string")
	}
	if authParts[0] != auth_type {
		return "", errors.New(fmt.Sprintf("Authorization scheme is %v, not %v", authParts[0], auth_type))
	}
	return authParts[1], nil
}

func DjangoBasicAuth(c *gin.Context) {
	auth_token, err := getAuthorizationHeader(c, "Basic")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, fiber.Map{"error": err.Error()})
	}

	sDec, err := base64.StdEncoding.DecodeString(auth_token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, fiber.Map{"error": "Failed to parse base64 basic credentials"})
		return
	}
	authParts := strings.SplitN(string(sDec), ":", 2)
	if len(authParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, fiber.Map{"error": "Failed to parse basic credentials"})
		return
	}

	user, err := LoginDjangoUser(authParts[0], authParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, fiber.Map{"error": "Unauthorized"})
		return
	}

	c.Set("user", user)
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
