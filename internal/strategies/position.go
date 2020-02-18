package strategies

import (
	"github.com/redmeros/htrade/models"
)

// Position to struct pozycji
type Position struct {
	Ticker     *models.Pair
	Direction  uint8
	OpenPrice  float64
	ClosePrice float64
	Quantity   int
}

// CurrentValue na podstawie podanej ceny oblicza
// aktualna wartosc
func (ps *Position) CurrentValue(currentPrice float64) float64 {
	return ps.OpenPrice*float64(ps.Quantity) + ps.CurrentPL(currentPrice)
}

// CurrentPL zwraca aktualna wartosc pozycji
func (ps *Position) CurrentPL(currentPrice float64) float64 {
	return (currentPrice - ps.OpenPrice) * float64(ps.Quantity) * float64(ps.Direction)
}

// Positions to jest wrapper na tablicę pozycji
type Positions struct {
	positions []*Position
}

// Exists zwraca info czy pozycja dla danej pary istnieje
// w kolekcji
func (ps *Positions) Exists(p *models.Pair) bool {
	for _, pos := range ps.positions {
		if pos.Ticker.ID == p.ID {
			return true
		}
	}
	return false
}
