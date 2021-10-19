package auth

import "github.com/golang-jwt/jwt/v4"

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

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JwtUserInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	IsStaff  bool   `json:"is_staff"`
}
type JwtCustomClaims struct {
	TokenType string      `json:"token_type"`
	UserId    int         `json:"user_id"`
	UserInfo  JwtUserInfo `json:"user_info"`
	jwt.RegisteredClaims
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Name          string `json:"name"`
	Locale        string `json:"locale"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
