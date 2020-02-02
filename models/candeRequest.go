package models

import "time"

const layout = "2006-01-02"

// CandleRequest zawiera request info
// o swieczkach
type CandleRequest struct {
	Instrument string `json:"instrument" binding:"required" form:"instrument"`
	From       string `json:"from" form:"from"`
	To         string `json:"to" form:"to"`
	// Kind       string `json:"kind" form:"kind" binding:"required"`
}

// FromDate parsuje date i ja zwraca
func (r *CandleRequest) FromDate() (time.Time, error) {
	return time.Parse(layout, r.From)
}

// ToDate parsuje date i ja zwraca
func (r *CandleRequest) ToDate() (time.Time, error) {
	return time.Parse(layout, r.To)
}
