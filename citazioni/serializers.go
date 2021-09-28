package citazioni

import (
	"github.com/gin-gonic/gin"
)

func CitazioniSerializer(citazioni []Citazione) []gin.H {
	response := []gin.H{}
	for _, citazione := range citazioni {
		fraseLen := len(citazione.Frase)
		appendStr := ""
		if fraseLen > 100 {
			fraseLen = 100
			appendStr = "..."
		}
		response = append(response, gin.H{
			"id":              citazione.ID,
			"frase":           citazione.Frase[:fraseLen] + appendStr,
			"autore":          citazione.Autore,
			"visualizzazioni": citazione.Visualizzazioni,
			"likes":           citazione.Likes,
		})
	}
	return response
}

func CitazioneSerializer(citazione Citazione) gin.H {
	return gin.H{
		"id":              citazione.ID,
		"frase":           citazione.Frase,
		"autore":          citazione.Autore,
		"visualizzazioni": citazione.Visualizzazioni,
		"likes":           citazione.Likes,
		"is_pubblica":     citazione.IsPubblica,
		"is_approved":     citazione.IsApproved,
	}
}
