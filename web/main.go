package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redmeros/htrade/config"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/web/middlewares"
)

func tryResloveConfig() (string, error) {
	files := []string{
		"config.json",
		"../config/config.json",
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

func main() {
	logger := logging.NewLogger("web.log")

	filename, err := tryResloveConfig()
	if err != nil {
		logger.Fatalf("Brak konfiguracji zamykam: %s", err.Error())
	}

	cfg, err := config.LoadConfig(filename)
	if err != nil {
		logger.Fatalf("Blad podczas czytania konfiguracji: %s", err.Error())
	}

	r := gin.Default()
	r.Use(middlewares.ConfigWriterMiddleware())
	r.Use(middlewares.DbMiddleware(cfg))

	r.GET("/ping", func(c *gin.Context) {

	})

	setRouting(r)

	log.Fatalf("Blad podczas uruchamiania %s", r.Run(cfg.Web.FullAddress()).Error())
}
