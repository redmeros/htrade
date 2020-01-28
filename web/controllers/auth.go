package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
	h "github.com/redmeros/htrade/web/helpers"
	wmodels "github.com/redmeros/htrade/web/models"
)

var loginOffset = time.Minute * 60

// Login loguje uzytkownika
// tzn podpisuje mu token jwt
func Login(c *gin.Context) {
	logger := logging.NewLogger("web.log")
	var json wmodels.Credentials
	if err := c.ShouldBindJSON(&json); err != nil {
		h.BadRequest(c, err.Error())
		return
	}

	db, err := GetDB(c)
	if err != nil {
		h.Bad(c, err.Error(), 500)
		return
	}

	cfg, err := GetConfig(c)
	if err != nil {
		h.Bad(c, err.Error(), 500)
		return
	}

	var dbuser models.User
	if err := db.Where("username = ? OR email = ?", json.GetIdentifier(), json.GetIdentifier()).Find(&dbuser).Error; err != nil {
		h.Bad(c, "Bad user or password", 403)
		return
	}

	if err = dbuser.CheckPassword(json.Password); err != nil {
		logger.Error(err.Error())
		logger.Error(len(json.Password))
		h.Bad(c, "Bad user or password", 403)
		return
	}

	secret := []byte(cfg.Web.Secret)
	issuedTime := time.Now()
	expirationTime := issuedTime.Add(loginOffset)
	claims := &wmodels.Claims{
		Username: dbuser.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  issuedTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		h.Bad(c, "Error creating jwt", 500)
		return
	}
	c.JSON(200, gin.H{
		"token":   tokenString,
		"expires": expirationTime.Unix(),
	})
	return
}

// Logout wylogowuje uzytkownika
func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "logout"})
}

// Refresh odswierza token jwt
func Refresh(c *gin.Context) {
	cuser, exists := c.Get("user")
	if !exists {
		h.Bad(c, "cannot get user", 403)
		c.Abort()
		return
	}
	user, ok := cuser.(*models.User)
	if !ok {
		h.Bad(c, "user is of wrong type", 500)
		c.Abort()
		return
	}

	cfg, err := GetConfig(c)
	if err != nil {
		h.Bad(c, "cannot get config", 500)
		c.Abort()
		return
	}

	secret := []byte(cfg.Web.Secret)
	issuedTime := time.Now()
	expirationTime := issuedTime.Add(loginOffset)
	claims := &wmodels.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  issuedTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		h.Bad(c, "Error during creating jwt", 500)
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"token":   tokenString,
		"expires": expirationTime.Unix(),
	})
	return
}

// SignUp jest handlerem dla signup
// metoda powinna byc POST
func SignUp(c *gin.Context) {
	var json models.User
	if err := c.ShouldBindJSON(&json); err != nil {
		h.BadRequest(c, err.Error())
		return
	}
	cfg, err := GetConfig(c)
	if err != nil {
		h.Bad(c, err.Error(), 500)
		return
	}
	if cfg.Web.SignupBlocked {
		h.Bad(c, "Signup blocked by admin", 403)
	}
	json.HashPass(cfg.Web.Secret)
	db, err := GetDB(c)
	if err != nil {
		h.Bad(c, err.Error(), 500)
		return
	}
	// spew.Dump(json)
	if err = db.Create(&json).Error; err != nil {
		h.Bad(c, err.Error(), 500)
		return
	}

	c.JSON(200, json.WithMaskedPass())
}
