package budgest

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func UserSerializer(user User) gin.H {
	return gin.H{
		"id":        user.ID,
		"full_name": strings.TrimSpace(user.FirstName + " " + user.LastName),
	}
}

func AmbitiSerializer(ambiti []Ambito) []gin.H {
	response := []gin.H{}
	for _, ambito := range ambiti {
		response = append(response, AmbitoSerializer(ambito))
	}
	return response
}

func AmbitoSerializer(ambito Ambito) gin.H {
	return gin.H{
		"id":              ambito.ID,
		"owner_id":        ambito.OwnerId,
		"owner":           UserSerializer(ambito.Owner),
		"num":             ambito.Num,
		"descrizione":     ambito.Descrizione,
		"is_active":       ambito.IsActive,
		"is_investimento": ambito.IsInvestimento,
	}
}
