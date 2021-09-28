package citazioni

import "github.com/gin-gonic/gin"

func CitazioniAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", CitazioniList)
	router.GET("/:id", CitazioneDetail)
}
func CitazioneAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", CitazioneRandomDetail)
}
