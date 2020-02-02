package models

import (
	"strconv"
	"time"
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

// MarshalJSON konwertuje date na unix timestamp
func (t ITime) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func ParseITime(format string, value string) (ITime, error) {
	t, err := time.Parse(time.RFC3339, value)
	return ITime(t), err
}

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
