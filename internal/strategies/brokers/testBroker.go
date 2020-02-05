package brokers

import (
	"github.com/redmeros/htrade/internal/strategies"
	"github.com/redmeros/htrade/models"
)

// TestBroker to zwykły offline'owy broker
type TestBroker struct {
	orders []*strategies.Order
}

// Orders implementacja dla IBroker
// Zwraca slic'a z wynikami
func (b *TestBroker) Orders() []*strategies.Order {
	return b.orders[:]
}

// PushOrders implementacja dla IBroker
func (b *TestBroker) PushOrders(orders []*strategies.Order) {
	panic("not implemented")
}

// RealizeOrders to implementacja dla IBroker
func (b *TestBroker) RealizeOrders(candles map[string]*models.Candle) []*strategies.Order {
	for _, order := range b.Orders() {
		// zamykam pozycję
		if order.Direction == 0 {
			candles[order.Instrument]
		}
	}
}
