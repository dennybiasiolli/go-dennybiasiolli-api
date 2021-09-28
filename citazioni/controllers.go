package citazioni

import (
	"errors"
	"net/http"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CitazioniList(c *gin.Context) {
	db := common.GetDB()
	var citazioni []Citazione
	db.Where(&Citazione{IsApproved: true, IsPubblica: true}).Find(&citazioni)
	c.JSON(200, citazioni)
}

func CitazioneDetail(c *gin.Context) {
	db := common.GetDB()
	var citazione Citazione
	id := c.Params.ByName("id")
	err := db.Where(&Citazione{IsApproved: true, IsPubblica: true}).First(&citazione, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, citazione)
}
