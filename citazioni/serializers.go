package citazioni

import "github.com/gofiber/fiber/v2"

func CitazioniSerializer(citazioni []Citazione) []fiber.Map {
	response := []fiber.Map{}
	for _, citazione := range citazioni {
		fraseLen := len(citazione.Frase)
		appendStr := ""
		if fraseLen > 100 {
			fraseLen = 100
			appendStr = "..."
		}
		response = append(response, fiber.Map{
			"id":              citazione.ID,
			"frase":           citazione.Frase[:fraseLen] + appendStr,
			"autore":          citazione.Autore,
			"visualizzazioni": citazione.Visualizzazioni,
			"likes":           citazione.Likes,
		})
	}
	return response
}

func CitazioneSerializer(citazione Citazione) fiber.Map {
	return fiber.Map{
		"id":              citazione.ID,
		"frase":           citazione.Frase,
		"autore":          citazione.Autore,
		"visualizzazioni": citazione.Visualizzazioni,
		"likes":           citazione.Likes,
		"is_pubblica":     citazione.IsPubblica,
		"is_approved":     citazione.IsApproved,
	}
}
