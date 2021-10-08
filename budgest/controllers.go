package budgest

import (
	"errors"
	"net/http"

	"github.com/dennybiasiolli/go-dennybiasiolli-api/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AmbitiList(c *gin.Context) {
	db := common.GetDB()
	var ambiti []Ambito
	var count int64
	// var ambitoFilter = Ambito{IsActive: true}
	err := db.Where("budgest_ambito.is_active = true").Joins("Owner").Order("budgest_ambito.num").Find(&ambiti).Count(&count).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"results": AmbitiSerializer(ambiti),
		"count":   count,
	})
}

func AmbitiDetail(c *gin.Context) {
	db := common.GetDB()
	var ambito Ambito
	id := c.Params.ByName("id")
	// var ambitoFilter = Ambito{IsActive: true}
	err := db.Where("budgest_ambito.is_active = true").Joins("Owner").First(&ambito, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, AmbitoSerializer(ambito))
}
