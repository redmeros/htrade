package strategies

import (
	"container/list"

	"github.com/redmeros/htrade/models"
)

// IAlgo to interfejs dla algorytmu
type IAlgo interface {
	OnTick(map[string]*models.Candle) ([]*AlgoResult, error)
}

// IBroker to interfesj dla Brokera,
// Broker odpowiedzialny jest za elementy które normalnie robi
// prawdziwy broker:
// 1. Obsługuje zlecenia
// 2. Prowadzi 'księgowość dla rachunki
// 3. Zamyka rachunek w razie bankrutu
type IBroker interface {
	// Orders zwraca zakolejkowane zlecenia
	// Zlecenia zrealizowane są oznaczane jako realized = true
	// i zwracane. Następnie są usuwane z kolejki
	Orders() *list.List
	Positions() *list.List
	// Realize
	RealizeOrders(map[string]*models.Candle) ([]*Position, []*Order, error)
	PushOrders(orders []*Order)
}

// IWallet to interfejs dla portfela
type IWallet interface {
	UpdatePositions(orders []*Order)
	Positions() []*Position
	FilterAlgosResult([]*Position, []*AlgoResult) ([]*Order, error)
}

// Position zawiera informacje o pozycji
type Position struct {
	OpenOrder  *Order
	CloseOrder *Order
	Instrument string
	Closed     bool
}

// UnrealizedPL zwraca wartosc niezrealizowanych zyskow
// strat z otwartej pozycji, na podstawie aktualnie przekazanej swieczki najgorsza mozliwa
// cena dla sprzedazy bedzie to high, dla kupna bedzie to low
// jezeli pozycja jest zamknieta
// skorzystaj z metody RealizedPL
func (pos *Position) UnrealizedPL(candle *models.Candle) float64 {
	panic("not implemented")
}

// RealizedPL zwraca wartosc zrealizowanych zyskow / strat
// na podstawie roznicy pomiedzy zrealizowanym zleceniem
// otwarcia i zamkniecia pozycji
func (pos *Position) RealizedPL() float64 {
	panic("not implemented")
}

// Order jest structem zawierającym
// informacje na temat zlecenie
// jeżeli zlecenie zostało zrealizowane to `Realized` powinno być ustawione na true
// `Direction` oznacza kierunek - ,1 - kupuj, -1 sprzedaj
// Podejście z Direction -1, 0, 1 jest złe, bo np dla 0 - skąd
// Broker ma wiedzieć czy ma sprzedać czy kupić/
type Order struct {
	ID              int
	PositionID      int
	PlacementTime   models.ITime
	RealizationTime models.ITime
	Instrument      string
	Direction       OrderType
	StopLoss        float64
	TakeProfit      float64
	Realized        bool
	Quantity        float64
}

// AlgoResult zawiera informacje o wynikach
// z na danym etapie z algorytmu
type AlgoResult struct {
	Instrument string
	Time       models.ITime
	Result     int
}

// OrderType reprezentuje typ zlecenia
type OrderType int

const (
	// Buy to zwykle zlecenie kupna
	Buy OrderType = iota
	// Sell to zwykle zlecenie sprzedazy (krotka pozycja)
	Sell
	// Close to zamkniecie pozycji
	Close
)

func (o OrderType) String() string {
	return [...]string{"Buy", "Sell", "Close"}[o]
}
