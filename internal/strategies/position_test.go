package strategies

import (
	"testing"

	"github.com/redmeros/htrade/models"
)

func TestPositionPanicWhenWrongTick(t *testing.T) {
	pair := models.Pair{
		Major: "EUR",
		Minor: "USD",
	}

	pos := Position{
		Ticker:    &pair,
		Direction: 1,
		OpenPrice: 1.1,
		Quantity:  10,
	}
}
