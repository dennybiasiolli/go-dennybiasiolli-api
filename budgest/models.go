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
}

// TableName set Ambito's table name to be `budgest_ambito`
func (Ambito) TableName() string {
	return "budgest_ambito"
}
