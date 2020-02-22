package strategies

// NoMoneyManagement to manager który ustala wielkość pozycji w oparciu
// wartość całego portfela. Zakłada że jest tylko 5 pozycji w portfelu
type NoMoneyManagement struct {
	broker Broker
}

// SetBroker ustawia brokera
func (mm *NoMoneyManagement) SetBroker(b Broker) {
	mm.broker = b
}

// OnAlgoResults przeprowadza analize portfela
// i wystawia nowe zlecenia
func (mm *NoMoneyManagement) OnAlgoResults(data []*AlgoResult) {
	currentMoney := mm.broker.TotalValue()
	currentPositions := mm.broker.CurrentPositions()
	positionSize := currentMoney / 5
	var results []*MMResult
	for _, res := range data {
		if currentPositions.Exists(&res.Pair) {
			continue
		}
		var mr MMResult
		mr.Direction = res.Rating
		mr.Ticker = &res.Pair
		mr.Value = positionSize
		results = append(results, &mr)
	}

	// to musi być wywołane na końcu to jest zle - tutaj powinno być zwracanie rezultatów
	mm.broker.OnMoneyMResults(results)
}
