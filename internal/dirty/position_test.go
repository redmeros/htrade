package dirty

import (
	"testing"
	"time"

	"github.com/redmeros/htrade/models"
	"github.com/stretchr/testify/assert"
)

var tt = time.Now()

// https://www.oanda.com/us-en/trading/how-calculate-profit-loss/
func TestProfit(t *testing.T) {
	eurusd := models.Pair{
		Major: "EUR",
		Minor: "USD",
	}

	// POZYCJA KRÓTKA
	pos := Position{
		Ticker:        &eurusd,
		QuantityMajor: -10000,
		QuantityMinor: 9517,
	}

	assert.EqualValues(t, 12.00, pos.PL(0.9505, 1))
	assert.Equal(t, 12.00, pos.PLC(0.9505))
}

func TestLoss(t *testing.T) {
	usdjpy := models.Pair{
		Major: "USD",
		Minor: "JPY",
	}
	pos := Position{
		Ticker:        &usdjpy,
		QuantityMajor: 10000,
		QuantityMinor: -1150500,
	}
	assert.EqualValues(t, -52.424639580602886, pos.PL(114.45, 1.0/114.45))
	assert.Equal(t, -6000.0, pos.PLC(114.45))
}

func TestUintToFloat(t *testing.T) {
	t1 := uint8(1)
	assert.Equal(t, 1.0, float64(t1))
}

func TestNewLongPosition(t *testing.T) {
	usdjpy := models.Pair{
		Major: "USD",
		Minor: "JPY",
	}

	pos := NewPosition(&usdjpy, 10000, 1, 115.05, tt)
	assert.Equal(t, 10000.0, pos.QuantityMajor)
	assert.Equal(t, -1150500.0, pos.QuantityMinor)
}

func TestShort2(t *testing.T) {
	eurusd := models.Pair{
		Major: "EUR",
		Minor: "USD",
	}
	pos := NewPosition(&eurusd, 10000, -1, 1.35777, tt)
	assert.Equal(t, 49.2928036947388, pos.PL(1.35111, 1.0/1.35111))
}

func TestLong2(t *testing.T) {
	eurgbp := models.Pair{
		Major: "EUR",
		Minor: "GPB",
	}
	pos := NewPosition(&eurgbp, 10000, 1, 0.73044, tt)
	assert.Equal(t, -2.9259999999994397, pos.PL(0.73025, 1.54))

}

func TestNewShortPosition(t *testing.T) {
	eurusd := models.Pair{
		Major: "EUR",
		Minor: "USD",
	}
	pos := NewPosition(&eurusd, 10000, -1, 0.9517, tt)
	assert.Equal(t, -10000.0, pos.QuantityMajor)
	assert.Equal(t, 9517.0, pos.QuantityMinor)
}

func TestLoss3(t *testing.T) {
	usdjpy := models.Pair{
		Major: "USD",
		Minor: "JPY",
	}
	pos := Position{
		Ticker:        &usdjpy,
		QuantityMajor: 10000,
		QuantityMinor: -1150500,
	}
	assert.EqualValues(t, -54.64978595500501, pos.PL(114.45, 1.0/109.79))
}

func TestProfit3(t *testing.T) {
	eurusd := models.Pair{
		Major: "EUR",
		Minor: "USD",
	}
	pos := NewPosition(&eurusd, 10000, 1, 1.10155, tt)
	assert.InDelta(t, 12.696706, pos.PL(1.10292, 1.0/1.07902), 0.01)
}
