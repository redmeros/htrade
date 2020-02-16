package strategies

import "github.com/redmeros/htrade/models"

// BuyAndHoldStrategy to zwykła strategia testująca
// ten system - polega na kupieniu w pierwszym możliwym
// momencie i zamknięciu wszystkich pozycji po 200 świeczkach
type BuyAndHoldStrategy struct {
	moneyManager MoneyM
	bought       map[string]bool
}

// OnData jest implementacją Consumer'a
// przewiduje się że data będzie type []*Candle
func (s *BuyAndHoldStrategy) OnData(data interface{}) {
	d := data.([]*models.Candle)
	results := make([]*AlgoResult, len(d))
	for _, candle := range d {
		if s.bought[candle.Pair.Name()] == true {
			continue
		}
		s.bought[candle.Pair.Name()] = true
		result := AlgoResult{
			Pair:   candle.Pair,
			Rating: 1,
		}
		results = append(results, &result)
	}
	// To wykonanie musi być końcu nie ma innej możliwości
	// przynajmniej na razie nie widzę
	s.moneyManager.OnAlgoResults(results)
}

// SetMoneyM ustawia managera
func (s *BuyAndHoldStrategy) SetMoneyM(moneyManager MoneyM) {
	if moneyManager != nil {
		s.moneyManager = moneyManager
	}
}

// MoneyM zwraca używanego menedzera
func (s *BuyAndHoldStrategy) MoneyM() MoneyM {
	return s.MoneyM()
}
