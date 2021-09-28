package citazioni

import "github.com/gin-gonic/gin"

func CitazioniAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", CitazioniList)
	router.GET("/:id", CitazioneDetail)
}
