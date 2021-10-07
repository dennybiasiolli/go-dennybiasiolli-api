package budgest

import (
	"github.com/dennybiasiolli/go-dennybiasiolli-api/auth"
	"github.com/gin-gonic/gin"
)

func AmbitiSerializer(ambiti []Ambito) []gin.H {
	response := []gin.H{}
	for _, ambito := range ambiti {
		response = append(response, AmbitoSerializer(ambito))
	}
	return response
}

func AmbitoSerializer(ambito Ambito) gin.H {
	return gin.H{
		"id": ambito.ID,
		// "owner_id":        ambito.OwnerId,
		"owner":           auth.UserSerializer(ambito.Owner),
		"num":             ambito.Num,
		"descrizione":     ambito.Descrizione,
		"is_active":       ambito.IsActive,
		"is_investimento": ambito.IsInvestimento,
	}
}
