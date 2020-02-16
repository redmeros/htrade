package strategies

import "github.com/redmeros/htrade/models"

// Feeder jest ogólnym interfejsem który zasila
// danymi Consumerów którzy są wpisani na listę
// subskrypcyjną
type Feeder interface {
	Subscribe(consumer Consumer)
	Unsubscribe(consumer Consumer)
	Consumers() []Consumer
	StartFeeding() error
	AddPair(pairname string)
	// RemovePair(pairneme string)
}

// Consumer jest interfejsem dla wszystkich elementów
// które są w stanie obsługiwać dane
type Consumer interface {
	OnData(data interface{})
}

// Algorithm zwraca doMoneyM info
// czy kupować czy sprzedawać i ustawia limity
type Algorithm interface {
	OnData(data interface{})
	SetMoneyM(moneyManager MoneyM)
	MoneyM() MoneyM
}

// MoneyM zarządza wielkościami pozycji
type MoneyM interface {
	OnAlgoResults(data []*AlgoResult)
	SetBroker(broker Broker)
}

// Broker jest uruchamiany w razie potrzeby
// przez MoneyM
type Broker interface {
	OnMoneyMResults(results []*MMResult)
	OnData(data interface{})
	Results() []Results
	CurrentPositions() *Positions
	CurrentOrders()
	TotalValue() float64
}

type MMResult struct {
	Ticker *models.Pair
	Value  float64
}

// AlgoResult jest rezultatem wysylanym
// z algo do money managementu
// na początku prosty schemat rating
//  1 - long
// -1 - short
//  0 - close
type AlgoResult struct {
	Pair   models.Pair
	Rating int
}

type Position struct {
	Ticker *models.Pair
}

type Positions struct {
	positions []*Position
}

func (ps *Positions) Exists(p *models.Pair) bool {
	for _, pos := range ps.positions {
		if pos.Ticker.ID == p.ID {
			return true
		}
	}
	return false
}

type Results struct {
}

// Run jest główną funkcją uruchamiającą test
func Run(dataFeeder Feeder, algo Algorithm, moneyM MoneyM, broker Broker) ([]Results, error) {

	moneyM.SetBroker(broker)
	algo.SetMoneyM(moneyM)
	dataFeeder.Subscribe(broker)
	dataFeeder.Subscribe(algo)
	dataFeeder.StartFeeding()

	return broker.Results(), nil
}
