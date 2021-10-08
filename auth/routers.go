package auth

import "github.com/gin-gonic/gin"

func JwtTokenRegister(router *gin.RouterGroup) {
	router.POST("/", TokenObtain)
}
