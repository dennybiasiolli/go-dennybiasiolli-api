package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func DjangoBasicAuth(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	authParts := strings.SplitN(authorizationHeader, " ", 2)
	if len(authParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse authentication string"})
		return
	}
	if authParts[0] != "Basic" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Authorization scheme is %v, not Basic", authParts[0])})
		return
	}
	sDec, err := base64.StdEncoding.DecodeString(authParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse base64 basic credentials"})
		return
	}
	authParts = strings.SplitN(string(sDec), ":", 2)
	if len(authParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse basic credentials"})
		return
	}

	db := common.GetDB()
	var user User = User{
		IsActive: true,
		Username: authParts[0],
	}
	err = db.Where(&user).First(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	res, err := CheckPassword(authParts[1], user.Password)
	if !res || err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("user", user)
}

func DjangoJwtAuth(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	authParts := strings.SplitN(authorizationHeader, " ", 2)
	if len(authParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse authentication string"})
		return
	}
	if authParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Authorization scheme is %v, not Bearer", authParts[0])})
		return
	}

	// check token validity
	token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
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
