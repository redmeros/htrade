package dirty

import (
	"encoding/json"
	"testing"

	"github.com/redmeros/htrade/models"
	"github.com/stretchr/testify/assert"
)

func TestIsStrategy(t *testing.T) {
	strategy := new(BuyHoldStrategy)
	assert.Implements(t, (*models.IStrategy)(nil), strategy)
}

func TestRunBuyHold(t *testing.T) {
	strategy := BuyHoldStrategy{}
	strategy.SetupDefault()

	assert.NotPanics(t, func() { strategy.Run() })
	results, err := strategy.Run()
	assert.NoError(t, err)
	assert.Len(t, results.Positions, 1)
	for _, pos := range results.Positions {
		assert.True(t, pos.Closed, "All positions should be closed")
	}
	assert.Equal(t, 1.10155, results.Positions[0].OpenRate)
	assert.Equal(t, 1.10292, results.Positions[0].CloseRate)
}

func TestJsonWorks(t *testing.T) {
	s := NewBuyAndHoldStrategy()
	s.SetupDefault()

	b, err := json.Marshal(&s)
	assert.NoError(t, err)
	var mm map[string]interface{}
	json.Unmarshal(b, &mm)
	assert.Equal(t, mm["code_name"], "buyhold")
	assert.Equal(t, mm["display_name"], "Buy & Hold")
	// etc...
}
