package articoli

import (
	"github.com/gin-gonic/gin"
)

func ArticoliSerializer(articoli []Articolo) []gin.H {
	response := []gin.H{}
	for _, articolo := range articoli {
		response = append(response, gin.H{
			"id":                 articolo.ID,
			"data_pubblicazione": articolo.DataPubblicazione,
			"titolo_it":          articolo.TitoloIt,
			"titolo_en":          articolo.TitoloEn,
			"autore_id":          articolo.AutoreId,
		})
	}
	return response
}

func ArticoloSerializer(articolo Articolo) gin.H {
	return gin.H{
		"id":                 articolo.ID,
		"data_pubblicazione": articolo.DataPubblicazione,
		"titolo_it":          articolo.TitoloIt,
		"titolo_en":          articolo.TitoloEn,
		"testo_it":           articolo.TestoIt,
		"testo_en":           articolo.TestoEn,
		"autore_id":          articolo.AutoreId,
	}
}
