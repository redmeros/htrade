package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// ITime jest typem ktory pozwala podmienic zachowanie sie time.Time
// podczas konwersji do daty !!!!! UWAGA NIE WIEM JAK TO SIEZACHOWA DLA COLLECTORA
// ale nie powinno miec problemu ??
type ITime time.Time

// Candle to model bazy danych
type Candle struct {
	ID          uint64  `gorm:"primary_key" json:"id"`
	Time        ITime   `gorm:"unique_index:idx_pair_time" json:"date"`
	PairID      uint64  `gorm:"ForeignKey:PairID;unique_index:idx_pair_time" json:"pair_id"`
	Pair        Pair    `json:"-"`
	Volume      int     `json:"volume"`
	Granularity string  `gorm:"type:varchar(2)" json:"granularity"`
	OpenAsk     float64 `json:"open_ask"`
	OpenBid     float64 `json:"open_bid"`
	HighAsk     float64 `json:"high_ask"`
	HighBid     float64 `json:"high_bid"`
	LowAsk      float64 `json:"low_ask"`
	LowBid      float64 `json:"low_bid"`
	CloseAsk    float64 `json:"close_ask"`
	CloseBid    float64 `json:"close_bid"`
}

// LowHigh zwraca zwykłego tupla z low i z high
func (c *Candle) LowHigh(priceType string) (float64, float64) {
	switch priceType {
	case "bid":
		return c.LowBid, c.HighAsk
	case "ask":
		return c.LowAsk, c.LowBid
	default:
		return 0, 0
	}
}

// MarshalJSON konwertuje date na unix timestamp
func (t ITime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// func ParseITime(format string, value string) (ITime, error) {
// 	t, err := time.Parse(time.RFC3339, value)
// 	return ITime(t), err
// }

// UnmarshalJSON konwertujez unix timestamp do daty
func (t *ITime) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

// Date helper co Date'a robi ITime'a
func Date(year int, month time.Month, day int, hour int, minute int, second int, nsec int, loc *time.Location) ITime {
	return ITime(time.Date(year, month, day, hour, minute, second, nsec, loc))
}

// GetCandlesByPairName zwraca swieczki po nazwie - wszystkie
func GetCandlesByPairName(db *gorm.DB, pair string) ([]Candle, error) {
	var candles []Candle
	q := db.Joins("LEFT JOIN  pairs ON candles.pair_id = pairs.id")
	q = q.Where("CONCAT(pairs.major, pairs.minor) = ?", pair)
	q = q.Order("time")
	if err := q.Find(&candles).Error; err != nil {
		return nil, err
	}
	return candles, nil
}
