package strategies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqlFeederIsFeeder(t *testing.T) {
	var f SQLFeeder
	assert.Implements(t, (*Feeder)(nil), &f)
}

func TestSubscribeUnsubscribe(t *testing.T) {
	sqlFeeder := SQLFeeder{}
	var f Feeder
	f = &sqlFeeder
	f.Subscribe(nil)
	assert.Len(t, f.Consumers(), 0)

	consumer := PrintingDataConsumer{}
	f.Subscribe(&consumer)
	assert.Len(t, f.Consumers(), 1)

	f.Unsubscribe(&consumer)
	assert.Len(t, f.Consumers(), 0)

	lst := f.Consumers()
	lst = append(lst, &consumer)
	assert.Len(t, f.Consumers(), 0)
}
