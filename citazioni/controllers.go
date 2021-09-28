package citazioni

import (
	"encoding/json"
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

func CitazioneCreate(c *gin.Context) {
	var input CreateCitazioneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userTrack, _ := json.MarshalIndent(gin.H{
		"User-Agent":      c.Request.Header.Get("User-Agent"),
		"x-forwarded-for": c.Request.Header.Get("x-forwarded-for"),
		// "Remote-Addr": c.Request.Header.Get("Remote-Addr"),
		"ClientIP": c.ClientIP(),
	}, "", "  ")
	userTrackStr := string(userTrack[:])
	citazione := Citazione{
		Frase:         input.Frase,
		Autore:        input.Autore,
		UserTrackJson: &userTrackStr,
	}
	db := common.GetDB()
	db.Create(&citazione)
	c.JSON(http.StatusOK, citazione)
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
