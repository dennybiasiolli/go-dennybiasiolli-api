package citazioni

import "github.com/gin-gonic/gin"

func CitazioniAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", CitazioniList)
	router.GET("/:id", CitazioneDetail)
	router.POST("/", CitazioneCreate)
}
func CitazioneAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", CitazioneRandomDetail)
}
