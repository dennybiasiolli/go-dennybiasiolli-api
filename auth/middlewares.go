package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
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
