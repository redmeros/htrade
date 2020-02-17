package strategies

import (
	"time"

	"github.com/redmeros/htrade/models"
)

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
	Results() []*Result
	CurrentPositions() *Positions
	CurrentOrders() []*Order
	TotalValue() float64
}

// MMResult służy do przechowywania informacji
// o rezultatach obliczeń MoneyManagera
type MMResult struct {
	Ticker    *models.Pair
	Value     float64
	Direction int
}

// Order zawiera dane o zleceniu
type Order struct {
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

// Result zawiera informacje
// o rezultatach w danym dniu
type Result struct {
	Time time.Time
}

// Run jest główną funkcją uruchamiającą test
func Run(dataFeeder Feeder, algo Algorithm, moneyM MoneyM, broker Broker) ([]*Result, error) {

	moneyM.SetBroker(broker)
	algo.SetMoneyM(moneyM)
	dataFeeder.Subscribe(broker)
	dataFeeder.Subscribe(algo)
	dataFeeder.StartFeeding()

	return broker.Results(), nil
}
