package budgest

import "github.com/dennybiasiolli/go-dennybiasiolli-api/auth"

type Ambito struct {
	ID             int       `gorm:"AUTO_INCREMENT; primaryKey"`
	OwnerId        int       `gorm:"not null"`
	Owner          auth.User `gorm:"foreignKey:OwnerId"`
	Num            int       `gorm:"not null;"`
	Descrizione    string    `gorm:"not null; size:200"`
	IsActive       bool      `gorm:"not null; default:true"`
	IsInvestimento bool      `gorm:"not null; default:false"`
	Colore         string    `gorm:"not null; size:20; default:'#cccccc'"`
}

// TableName set Ambito's table name to be `budgest_ambito`
func (Ambito) TableName() string {
	return "budgest_ambito"
}

type AmbitoCreateInput struct {
	Num            int    `json:"num" xml:"num" form:"num" validate:"required,number"`
	Descrizione    string `json:"descrizione" xml:"descrizione" form:"descrizione" validate:"required"`
	IsActive       *bool  `json:"is_active" xml:"is_active" form:"is_active" validate:"required"`
	IsInvestimento *bool  `json:"is_investimento" xml:"is_investimento" form:"is_investimento" validate:"required"`
}
