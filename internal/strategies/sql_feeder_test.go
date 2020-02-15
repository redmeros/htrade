package strategies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqlFeederImplementsFeeder(t *testing.T) {
	assert.Implements(t, (*Feeder)(nil), &SQLFeeder{})
}

func TestAddingEURUSDWorks(t *testing.T) {
	feeder := SQLFeeder{}
	feeder.AddPair("EURUSD")
	assert.Len(t, feeder.pairs, 1)

	pair := feeder.pairs[0]
	assert.Equal(t, "EUR", pair.Major)
	assert.NotEqual(t, pair.ID, 0)
}
