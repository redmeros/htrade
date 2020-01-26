package controllers

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

func resolveCollectorExec() (string, error) {
	var paths = []string{
		"../dist/dataCollector",
		"dataCollector",
	}
	for _, path := range paths {
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}
		if !info.IsDir() {
			return path, nil
		}
	}
	return "", errors.New("Żaden ze standardowych plików kolektora nie istnieje")
}

// StartCollector startuje nowy data collector
// i tworzy dla niego pid
func StartCollector(c *gin.Context) {
	// path, err := resolveCollectorExec()
	// if err != nil {
	// 	h.Bad(c, "cannot find collector exec", 500)
	// 	c.Abort()
	// 	return
	// }

	// cmd := exec.Command(path)

}
