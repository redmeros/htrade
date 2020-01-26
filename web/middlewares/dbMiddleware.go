package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/redmeros/htrade/config"
	"github.com/redmeros/htrade/internal/db"
)

var idb *gorm.DB = nil

func getOrCreateDb(c *config.Config) (*gorm.DB, error) {
	if idb != nil {
		return idb, nil
	}
	var err error
	idb, err = db.GetDb(c)
	if err != nil {
		return nil, err
	}
	return idb, nil
}

// DbMiddleware wstrzykuje wskaznik do bazy danych
// do contextu
func DbMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		d, err := getOrCreateDb(cfg)
		if err != nil {
			panic(err)
		}
		c.Set("DB", d)
		c.Next()
	}
}
