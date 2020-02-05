package algos

import (
	"github.com/redmeros/htrade/internal/strategies"
	"github.com/redmeros/htrade/models"
)

// BuyAndHoldAglo to algorytm który kupuje
// na wejsciu i trzyma pozycje
type BuyAndHoldAglo struct {
	boughtMap map[string]bool
}

// OnTick Na pierwszym ticku ustawia w mapie na true i daje rekomendacje
// kupna bez stop lossa
func (a *BuyAndHoldAglo) OnTick(map[string]*models.Candle) []*strategies.AlgoResult {
	panic("Not implemented")
}
