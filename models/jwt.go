package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims zawiera requesty do authu
type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
