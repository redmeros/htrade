package strategies

// NewOandaFxBroker zwraca nowego brokera
// z domyślnymi wartościami
func NewOandaFxBroker() *OandaFxBroker {
	b := OandaFxBroker{
		Leverage:      100,
		StartingMoney: 1000,
	}
	return &b
}

// OandaFxBroker to broker który zarządza
// pozycjami i realizuje zlecenia
type OandaFxBroker struct {
	Leverage      int
	StartingMoney float64
}

// OnMoneyMResults kolejkuje nowe zlecenia na podstawie
// danych z moneyManagement
func (b *OandaFxBroker) OnMoneyMResults(results []*MMResult) {
	panic("not implemented")
}

// OnData realizuje zakolejkowane zlecenia i ustawia
// ceny instrumentów dla konkretnych pozycji
func (b *OandaFxBroker) OnData(data interface{}) {
	panic("not implemented")
}

// Results zwraca wszystkie dostępne rezultaty
func (b *OandaFxBroker) Results() []*Result {
	panic("not implemented")
}

// CurrentPositions zwraca aktualne pozycje
func (b *OandaFxBroker) CurrentPositions() *Positions {
	panic("not implemented")
}

// CurrentOrders zwraca aktualne niezrealizowane zlecenia
func (b *OandaFxBroker) CurrentOrders() []*Order {
	panic("not implemented")
}

// TotalValue zwraca aktualną wartość całego portfela
func (b *OandaFxBroker) TotalValue() float64 {
	panic("not implemented")
}
