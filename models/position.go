package models

import (
	"time"
)

// Position to reprezentacja pozycji
type Position struct {
	Ticker        *Pair
	QuantityMajor float64
	QuantityMinor float64

	OpenRate             float64
	OpenTime             time.Time
	CloseRate            float64
	CloseTime            time.Time
	CloseMinorToHomeRate float64
	Closed               bool
	CloseProfitInMinor   float64
	CloseProfitInHome    float64
}

// Duration zwraca ilosc czasu jaka
func (p *Position) Duration() time.Duration {
	return p.OpenTime.Sub(p.CloseTime)
}

// Close zamyka pozycje wylicza PL i zapisuje daty i rate'y
func (p *Position) Close(currentPrice float64, minorToHome float64, time time.Time) {
	p.CloseRate = currentPrice
	p.CloseMinorToHomeRate = minorToHome
	p.CloseTime = time
	p.Closed = true
	p.CloseProfitInMinor = p.PLC(p.CloseRate)
	p.CloseProfitInHome = p.PL(p.CloseRate, p.CloseMinorToHomeRate)
}

// PL Zwraca PL w walucie w walucie rachunku
func (p *Position) PL(currentPrice float64, minorToHome float64) float64 {
	newminorq := p.QuantityMajor * currentPrice
	return (p.QuantityMinor + newminorq) * minorToHome
}

// PLC zwraca PL w walucie minor
func (p *Position) PLC(currentPrice float64) float64 {
	newminorq := p.QuantityMajor * currentPrice
	return p.QuantityMinor + newminorq
}

// NewPosition zwraca nowa pozycje
func NewPosition(ticker *Pair, quantity float64, direction int8, rate float64, time time.Time) *Position {
	pos := Position{
		Ticker:        ticker,
		QuantityMajor: quantity * float64(direction),
		QuantityMinor: quantity * -float64(direction) * rate,
		OpenRate:      rate,
		CloseRate:     0,
		OpenTime:      time,
		Closed:        false,
	}
	return &pos
}
