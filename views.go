package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetArticoli(c *gin.Context) {
	var articoli []Articolo
	db.Where("data_pubblicazione IS NOT NULL").Find(&articoli)
	c.JSON(200, articoli)
}
func GetArticolo(c *gin.Context) {
	var articolo Articolo
	id := c.Params.ByName("id")
	err := db.Where("data_pubblicazione IS NOT NULL").First(&articolo, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, articolo)
}

func GetCitazioni(c *gin.Context) {
	var citazioni []Citazione
	db.Where(&Citazione{IsApproved: true, IsPubblica: true}).Find(&citazioni)
	c.JSON(200, citazioni)
}
func GetCitazione(c *gin.Context) {
	var citazione Citazione
	id := c.Params.ByName("id")
	err := db.Where(&Citazione{IsApproved: true, IsPubblica: true}).First(&citazione, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, citazione)
}
