package auth

type User struct {
	ID          int    `gorm:"AUTO_INCREMENT; primaryKey"`
	IsSuperuser bool   `gorm:"not null;"`
	Username    string `gorm:"uniqueIndex; not null; size:150"`
	Password    string `gorm:"not null; size:128"`
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
