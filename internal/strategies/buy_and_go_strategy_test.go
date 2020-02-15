package strategies 


import (
	"testing"

	"github.com/hhkbp2/testify/assert"
)

func TestBuyAndHoldStrategyIsAlgo(t* testing.T) {
	var s BuyAndHoldStrategy
	assert.Implements(t, (*Algorithm)(nil), &s)
}

