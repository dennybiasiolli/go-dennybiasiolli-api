package citazioni

import "time"

type Citazione struct {
	ID              int       `gorm:"AUTO_INCREMENT; primaryKey"`
	Frase           string    `gorm:"not null;type:text"`
	Autore          string    `gorm:"not null; size:100"`
	IsPubblica      bool      `gorm:"not null;"`
	Visualizzazioni uint      `gorm:"not null;"`
	Likes           uint      `gorm:"not null;"`
	IsApproved      bool      `gorm:"not null;"`
	UserTrackJson   *string   `gorm:"type:text"`
	CreatedDate     time.Time `gorm:"not null;"`
	ModifiedDate    time.Time `gorm:"not null;"`
}

// TableName set Citazione's table name to be `citazioni_citazione`
func (Citazione) TableName() string {
	return "citazioni_citazione"
}

type CreateCitazioneInput struct {
	Frase  string `json:"frase" validate:"required"`
	Autore string `json:"autore" validate:"required"`
}
