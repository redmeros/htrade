package models

import (
	"time"
)

// Candle to model bazy danych
type Candle struct {
	ID          uint64    `gorm:"primary_key"`
	Time        time.Time `gorm:"unique_index:idx_pair_time"`
	PairID      uint64    `gorm:"ForeignKey:PairID;unique_index:idx_pair_time"`
	Pair        Pair
	Volume      int
	Granularity string `gorm:"type:varchar(2)"`
	OpenAsk     float64
	OpenBid     float64
	HighAsk     float64
	HighBid     float64
	LowAsk      float64
	LowBid      float64
	CloseAsk    float64
	CloseBid    float64
}
