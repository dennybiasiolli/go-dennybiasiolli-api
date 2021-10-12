package articoli

import "github.com/gofiber/fiber/v2"

func ArticoliSerializer(articoli []Articolo) []fiber.Map {
	response := []fiber.Map{}
	for _, articolo := range articoli {
		response = append(response, fiber.Map{
			"id":                 articolo.ID,
			"data_pubblicazione": articolo.DataPubblicazione,
			"titolo_it":          articolo.TitoloIt,
			"titolo_en":          articolo.TitoloEn,
			"autore_id":          articolo.AutoreId,
		})
	}
	return response
}

func ArticoloSerializer(articolo Articolo) fiber.Map {
	return fiber.Map{
		"id":                 articolo.ID,
		"data_pubblicazione": articolo.DataPubblicazione,
		"titolo_it":          articolo.TitoloIt,
		"titolo_en":          articolo.TitoloEn,
		"testo_it":           articolo.TestoIt,
		"testo_en":           articolo.TestoEn,
		"autore_id":          articolo.AutoreId,
	}
}
