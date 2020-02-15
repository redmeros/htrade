package strategies

type BuyAndHoldStrategy struct {
	moneyManager MoneyM
	bought       bool
}

func (s *BuyAndHoldStrategy) OnData(data interface{}) {
	s.bought = true
	// na koncu musi być wywolany moneyM
	s.moneyManager.OnData(data)
}

func (s *BuyAndHoldStrategy) SetMoneyM(moneyManager MoneyM) {
	if moneyManager != nil {
		s.moneyManager = moneyManager
	}
}

func (s *BuyAndHoldStrategy) MoneyM() MoneyM {
	return s.MoneyM()
}
