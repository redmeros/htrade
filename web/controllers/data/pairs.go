package data

import (
	"github.com/gin-gonic/gin"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
	cc "github.com/redmeros/htrade/web/controllers"
	h "github.com/redmeros/htrade/web/helpers"
	"strconv"
)

var logger = logging.NewLogger("web.log")

// GetPairs zwraca wszystkie dostepne pary
func GetPairs(c *gin.Context) {
	db, err := cc.GetDB(c)
	if err != nil {
		h.Badf(c, "Cannot get db: %s", 500, err)
		return
	}
	var pairs []models.Pair
	if db.Find(&pairs).Error != nil {
		h.Badf(c, "Cannot get pairs: %s", 500, err)
		return
	}

	c.JSON(200, pairs)
}

// GetPairByName zwraca pare po jej nazwie (np EURUSD)
func GetPairByName(c *gin.Context) {
	db, err := cc.GetDB(c)
	if err != nil {
		h.Badf(c, "Cannot get db: %s", 500, err)
	}
	name := c.Param("name")
	if len(name) != 6 {
		h.Bad(c, "Provided pair name has wrong name", 400)
		return
	}
	var pair models.Pair
	if db.Where("CONCAT(major, minor) = ?", name).First(&pair).Error != nil {
		h.Badf(c, "Cannot pair from db: %s", 400, err.Error())
		return
	}
	c.JSON(200, pair)
}

// NewPair zapisuje do bazy pare i ja zwraca
func NewPair(c *gin.Context) {
	var pair models.Pair
	if err := c.BindJSON(&pair); err != nil {
		h.Badf(c, "Cannot read pair data: %s", 400, err.Error())
		return
	}
	db, err := cc.GetDB(c)
	if err != nil {
		h.Badf(c, "Cannot get db: %s", 500, err.Error())
		return
	}
	if err = db.Create(&pair).Error; err != nil {
		h.Badf(c, "Cannot create pair %s", 400, err.Error())
		return
	}
	c.JSON(200, &pair)
}

// DeletePair usuwa pare z bazy danych
// 
func DeletePair(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.Badf(c, "Cannot get id: %s", 400, err.Error())
		return
	}

	if id == 0 {
		h.Bad(c, "Id must be set...", 400)
		return
	}

	db, err := cc.GetDB(c)
	if err != nil {
		h.Badf(c, "Cannot get db: %s", 500, err.Error())
		return
	}
	var pair models.Pair
	if err := db.First(&pair, id).Error; err != nil {
		h.Badf(c, "Cannot get pair %s", 400, err.Error())
		return
	}

	db.Delete(&pair)
}
