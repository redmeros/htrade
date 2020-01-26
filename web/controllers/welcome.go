package controllers

import (
	"github.com/gin-gonic/gin"
)

// Welcome wyswietla sie tylko dla zalogowanych
// uzytkownikow - uwaga ta wiadomosc jest scisle tajna!
func Welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome",
	})
}
