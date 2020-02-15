package brokers

import (
	"testing"
	"time"

	"github.com/redmeros/htrade/internal/strategies"
	"github.com/redmeros/htrade/models"
	"github.com/stretchr/testify/assert"
)

func TestBrokerImplementsIBroker(t *testing.T) {
	b := NewTestBroker()
	assert.Implements(t, (*strategies.IBroker)(nil), b, "Test if TestBroker implements IBroker")

}
func TestBrokerReutrnsPositions(t *testing.T) {
	b := NewTestBroker()
	assert.NotNil(t, b.Positions())
	q := b.Positions().Len()
	assert.Equal(t, 0, q)
}

func TestBrokerReturnsOrders(t *testing.T) {
	b := NewTestBroker()
	assert.NotNil(t, b.Orders())
	q := b.Orders().Len()
	assert.Equal(t, 0, q)
}

func TestPushOrderReflectsInOrders(t *testing.T) {
	b := NewTestBroker()
	var order = strategies.Order{Instrument: "EURPLN"}
	orders := []*strategies.Order{&order}
	b.PushOrders(orders)
	element := b.Orders().Front()
	elOrder := element.Value.(*strategies.Order)
	assert.Same(t, &order, elOrder)
}

func getTestCandles() map[string]*models.Candle {
	const DateFormat = "2006-01-02 15:04:05"
	candles := make(map[string]*models.Candle)
	t, _ := time.Parse(DateFormat, "2020-01-28 09:45:00")
	it := models.ITime(t)
	candles["EURUSD"] = &models.Candle{
		ID:     2,
		Time:   it,
		PairID: 3,
		Pair: models.Pair{
			ID:    1,
			Major: "EUR",
			Minor: "USD",
		},
		Volume:      193,
		Granularity: "M5",
		OpenAsk:     1.30166,
		OpenBid:     1.30152,
		HighAsk:     1.30170,
		HighBid:     1.30154,
		LowAsk:      1.30086,
		LowBid:      1.3007,
		CloseAsk:    1.30089,
		CloseBid:    1.30073,
	}
	return candles
}

func TestBrokerFullifyOrder(t *testing.T) {
	b := NewTestBroker()
	candles := getTestCandles()

	order := strategies.Order{
		PlacementTime: models.Date(2020, 01, 28, 0, 45, 0, 0, time.UTC),
		Direction:     1,
		Instrument:    "EURUSD",
	}
	orders := []*strategies.Order{&order}
	b.PushOrders(orders)
	b.RealizeOrders(candles)
	for _, order := range orders {
		assert.Equal(t, true, order.Realized, "All orders need to be marked as realized")

	}

	t.Fail()
}
