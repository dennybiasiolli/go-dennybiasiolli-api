package articoli

import (
	"errors"
	"net/http"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticoliList(c *gin.Context) {
	db := common.GetDB()
	var articoli []Articolo
	var count int64
	err := db.Where("data_pubblicazione IS NOT NULL").Find(&articoli).Order("-id").Count(&count).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"results": ArticoliSerializer(articoli),
		"count":   count,
	})
}

func ArticoliDetail(c *gin.Context) {
	db := common.GetDB()
	var articolo Articolo
	id := c.Params.ByName("id")
	err := db.Where("data_pubblicazione IS NOT NULL").First(&articolo, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ArticoloSerializer(articolo))
}
