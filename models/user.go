package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User is a user model
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100)" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(128)" binding:"required"`
	Email    string `json:"email" gorm:"type:varchar(100);unique_index" binding:"required"`
}

// HashPassword zwraca hash hasla
// Bcryptem
func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// WithMaskedPass password in user
// end return user
func (u *User) WithMaskedPass() *User {
	u.Password = "xxxx"
	return u
}

// CheckPassword zwraca nil jesli hash hasla
// i haslo sie zgadzaja, jesli nie zwraca error
func (u *User) CheckPassword(pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return err
	}
	return nil
}

// HashPass replacing a password with
// a hash
func (u *User) HashPass(secret string) error {
	hash, err := HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return nil
}
