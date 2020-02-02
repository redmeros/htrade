package data

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redmeros/htrade/models"
	cc "github.com/redmeros/htrade/web/controllers"
	h "github.com/redmeros/htrade/web/helpers"
)

// var logger = logging.NewLogger("web.log")

// Query odpowiada za wszystkie zapytania
// zwiazane z candlesami
type Query struct {
	Instrument string `binding:"required" json:"instrument" form:"instrument"`
	// date_from int `binding:`
}

// GetDataRange zwraca dla danej pary walutowej
// zakres dat dla ktorych sa dane
func GetDataRange(c *gin.Context) {
	var q Query
	if err := c.BindQuery(&q); err != nil {
		h.Badf(c, "Cannot bind qurey: %s", 400, err.Error())
		return
	}

	db, err := cc.GetDB(c)
	if err != nil {
		h.Badf(c, "Cannot get db: %s", 500, err.Error())
		return
	}
	var pair models.Pair
	if db.Where("CONCAT(major, minor) = ?", q.Instrument).First(&pair).Error != nil {
		h.Badf(c, "Cannot pair from db: %s", 400, err.Error())
		return
	}

	row := db.Raw("SELECT max(time) as maxtime, min(time) as mintime, count(time) as count FROM candles WHERE pair_id = ?;", pair.ID).Row()
	if err != nil {
		h.Badf(c, "Error during asking db", 500, err.Error())
		return
	}
	var (
		maxtime time.Time
		mintime time.Time
		count   uint
	)

	if err = row.Scan(&maxtime, &mintime, &count); err != nil {
		h.Badf(c, "Error with response to db %s", 400, err.Error())
		return
	}

	c.JSON(200, gin.H{"max_time": maxtime.Unix(), "min_time": mintime.Unix(), "count": count})
}

// GetCandles zwraca swieczki
func GetCandles(c *gin.Context) {

	var cr models.CandleRequest
	if err := c.BindQuery(&cr); err != nil {
		h.Badf(c, "Cannot bind request %s", 400, err.Error())
		return
	}
	var candles []models.Candle
	db, err := cc.GetDB(c)
	if err != nil {
		h.Badf(c, "Cannot get db %s", 500, err.Error())
		return
	}

	q := db.Joins("LEFT JOIN  pairs ON candles.pair_id = pairs.id")
	q = q.Where("CONCAT(pairs.major, pairs.minor) = ?", cr.Instrument)

	fromTime, err := cr.FromDate()
	const layout = "2006-01-02"
	logger.Debugf("From date: %s", fromTime.Format(layout))
	if err != nil && len(cr.From) > 0 {
		h.Badf(c, "Cannot parse time: %s", 400, err.Error())
		return
	} else if len(cr.From) > 0 {

		q = q.Where("time >= ?", fromTime)
	}

	toTime, err := cr.ToDate()
	logger.Debugf("To date: %s", toTime.Format(layout))
	if err != nil && len(cr.To) > 0 {
		h.Badf(c, "Cannot parse time: %s", 400, err.Error())
		return
	} else if len(cr.To) > 0 {
		q = q.Where("time <= ?", toTime)
	}

	q = q.Order("time")

	if err := q.Find(&candles).Error; err != nil {
		h.Badf(c, "Cannot get candles: %s", 400, err.Error())
		return
	}

	c.JSON(200, candles)
}
