package articoli

import "github.com/gin-gonic/gin"

func ArticoliAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", ArticoliList)
	router.GET("/:id", ArticoliDetail)
}
