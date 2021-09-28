package citazioni

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CitazioniList(c *gin.Context) {
	db := common.GetDB()
	var citazioni []Citazione
	var count int64
	err := db.Where(&Citazione{IsApproved: true, IsPubblica: true}).Find(&citazioni).Order("-id").Count(&count).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"results": CitazioniSerializer(citazioni),
		"count":   count,
	})
}

func CitazioneDetail(c *gin.Context) {
	db := common.GetDB()
	var citazione Citazione
	id := c.Params.ByName("id")
	err := db.Where(&Citazione{IsApproved: true, IsPubblica: true}).First(&citazione, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, CitazioneSerializer(citazione))
}

func CitazioneRandomDetail(c *gin.Context) {
	db := common.GetDB()
	var count int64
	var citazione Citazione
	search := c.DefaultQuery("search", "")
	qs := db.Model(&Citazione{}).Where(&Citazione{IsApproved: true, IsPubblica: true}).Where("frase ILIKE ? OR autore ILIKE ?", "%"+search+"%", "%"+search+"%").Count(&count)
	err := qs.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no quotes found"})
		return
	}
	i := rand.Int63n(count)
	qs.Limit(1).Offset(int(i)).First(&citazione)
	c.JSON(http.StatusOK, CitazioneSerializer(citazione))
}
