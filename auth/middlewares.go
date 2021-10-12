package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	sDec, err := base64.StdEncoding.DecodeString(auth_token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse base64 basic credentials"})
		return
	}
	authParts := strings.SplitN(string(sDec), ":", 2)
	if len(authParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse basic credentials"})
		return
	}

	user, err := LoginDjangoUser(authParts[0], authParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("user", user)
}

func DjangoJwtAuth(c *gin.Context) {
	auth_token, err := getAuthorizationHeader(c, "Bearer")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	// check token validity
	token, err := jwt.Parse(auth_token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(common.JWT_HMAC_SAMPLE_SECRET), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid && claims["token_type"] == "access") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	db := common.GetDB()
	var user User = User{
		IsActive: true,
	}
	err = db.Where(&user).First(&user, claims["user_id"]).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("user", user)
}

var JwtMiddleware = jwtware.New(jwtware.Config{
	SigningKey: []byte(common.JWT_HMAC_SAMPLE_SECRET),
})
