package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/redmeros/htrade/config"
	"github.com/redmeros/htrade/web/helpers"
)

var cfg *config.Config

func init() {
	filename, err := helpers.TryResloveConfig()
	if err != nil {
		panic(err.Error())
	}

	c, err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
		// panic("Cannot read config from middleware")
	}
	cfg = c
}

// ConfigWriterMiddleware wstrzykuje
// config do contextu
func ConfigWriterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	}
}
