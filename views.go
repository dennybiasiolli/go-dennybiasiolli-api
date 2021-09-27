package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticoli(c *gin.Context) {
	var users = []Articolo{
		{ID: 1, TitoloIt: "Titolo 1", TestoIt: "Testo 1", AutoreId: 1},
	}
	c.JSON(200, users)
}
func GetArticolo(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, _ := strconv.ParseInt(id, 0, 64)

	if user_id == 1 {
		articolo := Articolo{ID: 1, TitoloIt: "Titolo 1", TestoIt: "Testo 1", AutoreId: 1}
		content := articolo
		c.JSON(200, content)
	} else {
		content := gin.H{"error": "user with id " + id + " not found"}
		c.JSON(404, content)
	}
}
