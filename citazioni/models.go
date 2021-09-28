package citazioni

type Citazione struct {
	ID              int    `gorm:"AUTO_INCREMENT; primaryKey"`
	Frase           string `gorm:"not null;type:text[]"`
	Autore          string `gorm:"not null; size:100"`
	IsPubblica      bool   ``
	Visualizzazioni uint   ``
	Likes           uint   ``
	IsApproved      bool   ``
}

// TableName set Citazione's table name to be `citazioni_citazione`
func (Citazione) TableName() string {
	return "citazioni_citazione"
}
