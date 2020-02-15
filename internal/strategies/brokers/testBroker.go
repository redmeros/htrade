package brokers

import (
	"container/list"
	"fmt"

	"github.com/redmeros/htrade/internal/strategies"
	"github.com/redmeros/htrade/models"
)

// TestBroker to zwykły offline'owy broker
type TestBroker struct {
	orders    *list.List
	positions *list.List
}

// NewTestBroker jest konstruktorem dla TestBrokera
func NewTestBroker() *TestBroker {
	b := TestBroker{}
	b.orders = list.New().Init()
	b.positions = list.New().Init()
	return &b
}

// Orders implementacja dla IBroker
// Zwraca slic'a z wynikami
func (b *TestBroker) Orders() *list.List {
	return b.orders
}

// PushOrders implementacja dla IBroker
func (b *TestBroker) PushOrders(orders []*strategies.Order) {
	for _, order := range orders {
		b.orders.PushBack(order)
	}
}

func (b *TestBroker) tryBuy(order *strategies.Order, candle *models.Candle) {

}

// RealizeOrders to implementacja dla IBroker
func (b *TestBroker) RealizeOrders(candles map[string]*models.Candle) ([]*strategies.Position, []*strategies.Order, error) {
	for order := b.Orders().Front(); order != nil; order = order.Next() {
		o := order.Value.(*strategies.Order)
		switch o.Direction {
		case strategies.Buy:
			candle := candles[o.Instrument]
			b.tryBuy(o, candle)
			break
		case strategies.Sell:
			break
		case strategies.Close:
			break
		default:
			return nil, nil, fmt.Errorf("Order of type %s is not implemented", o.Direction)
		}
	}
	return nil, nil, nil
}

// Positions zwraca aktualne pozycje w portfelu
func (b *TestBroker) Positions() *list.List {
	return b.positions
}
