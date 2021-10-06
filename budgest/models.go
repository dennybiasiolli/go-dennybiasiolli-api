package budgest

type User struct {
	ID          int    `gorm:"AUTO_INCREMENT; primaryKey"`
	IsSuperuser bool   `gorm:"not null;"`
	Username    string `gorm:"uniqueIndex; not null; size:150"`
	FirstName   string `gorm:"not null; size:150"`
	LastName    string `gorm:"not null; size:150"`
	Email       string `gorm:"not null; size:254"`
	IsStaff     bool   `gorm:"not null;"`
	IsActive    bool   `gorm:"not null;"`
}

// TableName set User's table name to be `auth_user`
func (User) TableName() string {
	return "auth_user"
}

type Ambito struct {
	ID             int    `gorm:"AUTO_INCREMENT; primaryKey"`
	OwnerId        int    `gorm:"not null"`
	Owner          User   `gorm:"foreignKey:OwnerId"`
	Num            int    `gorm:"not null;"`
	Descrizione    string `gorm:"not null; size:200"`
	IsActive       bool   `gorm:"not null; default:true"`
	IsInvestimento bool   `gorm:"not null; default:false"`
}

// TableName set Ambito's table name to be `budgest_ambito`
func (Ambito) TableName() string {
	return "budgest_ambito"
}
