package strategies

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
	OnData(data interface{})
	SetBroker(broker Broker)
}

// Broker jest uruchamiany w razie potrzeby
// przez MoneyM
type Broker interface {
	Results() []Results
}

type Results struct {
}

// Run jest główną funkcją uruchamiającą test
func Run(dataFeeder Feeder, algo Algorithm, moneyM MoneyM, broker Broker) ([]Results, error) {

	moneyM.SetBroker(broker)
	algo.SetMoneyM(moneyM)

	dataFeeder.Subscribe(algo)
	dataFeeder.StartFeeding()

	return broker.Results(), nil
}
