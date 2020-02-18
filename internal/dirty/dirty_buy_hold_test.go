package dirty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunBuyHold(t *testing.T) {
	assert.NotPanics(t, func() { RunBuyHold(10000.0, 100.0) })
	//kupione za 1.10170
	//sprzedane za 1.10239
}
