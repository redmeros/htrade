package strategies

import (
	"testing"

	"github.com/redmeros/htrade/models"
	"github.com/stretchr/testify/assert"
)

func TestPositionPanicWhenWrongTick(t *testing.T) {
	pair := models.Pair{
		Major: "EUR",
		Minor: "USD",
	}

	pos := Position{
		Ticker:    &pair,
		Direction: 1,
		OpenPrice: 1.0,
		Quantity:  10,
	}

	assert.EqualValues(t, 12.0, pos.CurrentValue(1.2))
	assert.EqualValues(t, 10, pos.CurrentValue(1))
}
