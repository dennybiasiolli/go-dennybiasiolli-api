package budgest

import (
	"github.com/dennybiasiolli/go-dennybiasiolli-api/auth"
	"github.com/gofiber/fiber/v2"
)

func AmbitiSerializer(ambiti []Ambito) []fiber.Map {
	response := []fiber.Map{}
	for _, ambito := range ambiti {
		response = append(response, AmbitoSerializer(ambito))
	}
	return response
}

func AmbitoSerializer(ambito Ambito) fiber.Map {
	return fiber.Map{
		"id": ambito.ID,
		// "owner_id":        ambito.OwnerId,
		"owner":           auth.UserSerializer(ambito.Owner),
		"num":             ambito.Num,
		"descrizione":     ambito.Descrizione,
		"is_active":       ambito.IsActive,
		"is_investimento": ambito.IsInvestimento,
	}
}
