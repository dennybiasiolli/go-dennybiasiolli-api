package citazioni

type Citazione struct {
	ID              int     `gorm:"AUTO_INCREMENT; primaryKey"`
	Frase           string  `gorm:"not null;type:text"`
	Autore          string  `gorm:"not null; size:100"`
	IsPubblica      bool    `gorm:"not null;"`
	Visualizzazioni uint    `gorm:"not null;"`
	Likes           uint    `gorm:"not null;"`
	IsApproved      bool    `gorm:"not null;"`
	UserTrackJson   *string `gorm:"type:text"`
}

// TableName set Citazione's table name to be `citazioni_citazione`
func (Citazione) TableName() string {
	return "citazioni_citazione_go"
}
