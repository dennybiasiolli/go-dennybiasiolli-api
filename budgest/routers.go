package budgest

import "github.com/gin-gonic/gin"

func BudgestRegister(router *gin.RouterGroup) {
	router.GET("/ambiti/", AmbitiList)
	router.GET("/ambiti/:id/", AmbitiDetail)
}
