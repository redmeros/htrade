package strategies

import (
	"fmt"

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

func (ps *Position) CurrentValue(t *models.Candle) (float64, error) {
	if &t.Pair != ps.Ticker {
		return -1, fmt.Errorf("Wrong candle (%s) for position ticker (%s)", t.Pair.Name(), ps.Ticker.Name())
	}
	curPrice := 0.0
	// if ps.Direction > 0 {
	// 	t.
	// }
	return curPrice, nil
}

func (ps *Position) CurrentPL(t *models.Candle) {

}
