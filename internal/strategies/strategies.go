package strategies

import "fmt"

// Feeder jest ogólnym interfejsem który zasila
// danymi Consumerów którzy są wpisani na listę
// subskrypcyjną
type Feeder interface {
	Subscribe(consumer Consumer)
	Unsubscribe(consumer Consumer)
	Consumers() []Consumer
	StartFeeding()
}

// SQLFeeder jest konkretnym typem implementującym
// Feeder'a
type SQLFeeder struct {
	consumers []Consumer
}

// Subscribe dodaje do listy subskrybentów
// consumera
func (f *SQLFeeder) Subscribe(consumer Consumer) {
	if consumer != nil {
		f.consumers = append(f.consumers, consumer)
	}
}

// Unsubscribe usuwa z listy subskrybentów konkretnego subskrybenta
func (f *SQLFeeder) Unsubscribe(consumer Consumer) {
	for i, x := range f.consumers {
		if x == consumer {
			f.consumers = append(f.consumers[:i], f.consumers[i+1:]...)
		}
	}
}

// Consumers zwraca listę aktualnych
// subskrybentów
func (f *SQLFeeder) Consumers() []Consumer {
	return f.consumers
}

// StartFeeding rozpoczyna karmienie danymi
func (f *SQLFeeder) StartFeeding() {
	panic("not implemented")
}

// Consumer jest interfejsem dla wszystkich elementów
// które są w stanie obsługiwać dane
type Consumer interface {
	OnData(data interface{})
}

// PrintingDataConsumer jest odbiorcą danych
// który drukuje to co dostanie:D - stworzone do testów
type PrintingDataConsumer struct {
}

// OnData tylko wyświetla to co otrzymuje
func (p *PrintingDataConsumer) OnData(data interface{}) {
	fmt.Println(data)
}

// Algorithm zwraca doMoneyM info
// czy kupować czy sprzedawać i ustawia limity
type Algorithm interface {
	OnData(data interface{})
	SetMoneyM(moneyManager MoneyM)
}

// MoneyM zarządza wielkościami pozycji
type MoneyM interface {
	SetBroker(broker Broker)
}

// Broker jest uruchamiany w razie potrzeby
// przez MoneyM
type Broker interface {
	Results() interface{}
}

// Run jest główną funkcją uruchamiającą test
func Run(dataFeeder Feeder, algo Algorithm, moneyM MoneyM, broker Broker) (interface{}, error) {
	moneyM.SetBroker(broker)
	algo.SetMoneyM(moneyM)
	dataFeeder.Subscribe(algo)

	dataFeeder.StartFeeding()
	return broker.Results(), nil
}
