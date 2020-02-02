package oanda

import (
	"fmt"
	"strconv"
	"time"

	// . "github.com/ahmetb/go-linq"

	"github.com/redmeros/htrade/models"
)

// OHLC reprezentuje strukture
// ohlc oandy (dlatego sa stringi)
type OHLC struct {
	O string `json:"o"`
	H string `json:"h"`
	L string `json:"l"`
	C string `json:"c"`
}

// Candle reprezentuje odpowiedz typu Candle
// z serwera oandy
type Candle struct {
	Ask      OHLC   `json:"ask"`
	Bid      OHLC   `json:"bid"`
	Mid      OHLC   `json:"mid"`
	Complete bool   `json:"complete"`
	Time     string `json:"time"`
	Volume   int    `json:"volume"`
}

// CandleResponse reprezentuje odpowiedz
// z serwerów oandy na "instrument/PARA/candles"
type CandleResponse struct {
	Candles     []Candle `json:"candles"`
	Granularity string   `json:"granularity"`
	Instrument  string   `json:"instrument"`
}

// ToCandle konwertuje do modelu bazy danych
func (o *CandleResponse) ToCandle(pair *models.Pair) ([]*models.Candle, error) {

	var curPair *models.Pair = pair
	if curPair == nil {
		return nil, fmt.Errorf("Cannot find pair: %s", o.Instrument)
	}

	candles := make([]*models.Candle, len(o.Candles))
	for i, ocandle := range o.Candles {
		if ocandle.Complete != true {
			continue
		}
		var c = models.Candle{}
		t, err := time.Parse(time.RFC3339, ocandle.Time)
		if err != nil {
			return nil, fmt.Errorf("Cannot parse time: %s", ocandle.Time)
		}

		c.Time = models.ITime(t)
		c.Pair = *curPair
		c.PairID = curPair.ID
		c.Granularity = o.Granularity
		c.Volume = ocandle.Volume
		c.OpenAsk, _ = strconv.ParseFloat(ocandle.Ask.O, 64)
		c.HighAsk, _ = strconv.ParseFloat(ocandle.Ask.H, 64)
		c.LowAsk, _ = strconv.ParseFloat(ocandle.Ask.L, 64)
		c.CloseAsk, _ = strconv.ParseFloat(ocandle.Ask.C, 64)
		c.OpenBid, _ = strconv.ParseFloat(ocandle.Bid.O, 64)
		c.HighBid, _ = strconv.ParseFloat(ocandle.Bid.H, 64)
		c.LowBid, _ = strconv.ParseFloat(ocandle.Bid.L, 64)
		c.CloseBid, _ = strconv.ParseFloat(ocandle.Bid.C, 64)
		candles[i] = &c
	}

	return candles, nil
}
