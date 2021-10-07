package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func UserSerializer(user User) gin.H {
	return gin.H{
		"id":        user.ID,
		"full_name": strings.TrimSpace(user.FirstName + " " + user.LastName),
	}
}
