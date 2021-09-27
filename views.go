package main

import (
	"github.com/gin-gonic/gin"
)

func GetArticoli(c *gin.Context) {
	var articoli []Articolo
	db.Find(&articoli)
	c.JSON(200, articoli)
}
func GetArticolo(c *gin.Context) {
	var articolo Articolo
	id := c.Params.ByName("id")
	db.First(&articolo, id)
	c.JSON(200, articolo)
}
