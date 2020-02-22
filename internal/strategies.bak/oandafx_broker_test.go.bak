package strategies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOandaFxBrokerIsBroker(t *testing.T) {
	b := OandaFxBroker{}
	assert.Implements(t, (*Broker)(nil), &b)
}
