package strategies

import (
	"github.com/redmeros/htrade/models"
)

// IAlgo to interfejs dla algorytmu
type IAlgo interface {
	OnTick(map[string]*models.Candle) []*AlgoResult
}

// IBroker to interfesj dla Brokera
type IBroker interface {
	Orders() []*Order
	RealizeOrders(map[string]*models.Candle) []*Order
	// PushOrders odkłada zlecenia na stos,
	// zlecenia zamkniecia sa powinny być realizowane
	// cenie zamkniecia z danego ticku!
	PushOrders(orders []*Order)
}

// IWallet to interfejs dla portfela
type IWallet interface {
	UpdatePositions(orders []*Order)
	Positions() []*Position
	FilterAlgosResult([]*AlgoResult) []*Order
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
// `Direction` oznacza kierunek - ,1 - kupuj, -1 sprzedaj, 0 - zamknij
// Podejście z Direction -1, 0, 1 jest złe, bo np dla 0 - skąd
// Broker ma wiedzieć czy ma sprzedać czy kupić/
type Order struct {
	PlacementTime   models.ITime
	RealizationTime models.ITime
	Instrument      string
	Direction       uint8
	StopLoss        float64
	TakeProfit      float64
	Realized        bool
}

// AlgoResult zawiera informacje o wynikach
// z na danym etapie z algorytmu
type AlgoResult struct {
	Instrument string
	Time       models.ITime
	Result     int
}
