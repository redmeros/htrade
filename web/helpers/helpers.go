package helpers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// BadRequest jest helperem/skrotem
// zapisuje status http 400 i zwraca json
// "error": "message"
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}

// Badf to to samo co BAD ale z obsluga "f" czyli formatowania
func Badf(c *gin.Context, message string, status int, a ...interface{}) {
	c.JSON(status, gin.H{
		"error": fmt.Sprintf(message, a),
	})
	c.Abort()
}

// Bad jest helperem/skrotem
// zapisuje status i zwraca json
// "error": "message"
func Bad(c *gin.Context, message string, status int) {
	c.JSON(status, gin.H{"error": message})
	c.Abort()
}

// TryResloveConfig probuje znalezc
// plik config.json
func TryResloveConfig() (string, error) {
	files := []string{
		"config.json",
		"../config/config.json",
		"../../config/config.json",
	}
	for _, filename := range files {
		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			continue
		}
		if !info.IsDir() {
			return filename, nil
		}
	}
	return "", fmt.Errorf("Żaden ze standardowych plikow nie istnieje")
}
