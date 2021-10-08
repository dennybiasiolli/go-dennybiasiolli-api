package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func TokenObtain(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := common.GetDB()
	var user User = User{
		IsActive: true,
		Username: input.Username,
	}
	if err := db.Where(&user).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	res, err := CheckPassword(input.Password, user.Password)
	if !res || err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// JWT section
	mySigningKey := []byte(common.JWT_HMAC_SAMPLE_SECRET)

	// Create the Claims
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	c.JSON(http.StatusOK, gin.H{
		"access": ss,
	})
}
