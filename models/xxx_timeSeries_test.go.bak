package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddingInOrder(t *testing.T) {
	var lst OrderedListOfTimes
	t1 := time.Now()
	t2 := t1.Add(5 * time.Second)
	t3 := t2.Add(5 * time.Second)
	t4 := t3.Add(5 * time.Second)
	lst.Add(t1)
	lst.Add(t2)
	lst.Add(t3)
	lst.Add(t4)

	el1 := lst.Front()
	el2 := el1.Next()
	el3 := el2.Next()
	el4 := el3.Next()
	assert.Equal(t, t1, el1.Value)
	assert.Equal(t, t2, el2.Value)
	assert.Equal(t, t3, el3.Value)
	assert.Equal(t, t4, el4.Value)
}

func TestAddingInWrongOrder(t *testing.T) {
	var lst OrderedListOfTimes
	t1 := time.Now()
	t2 := t1.Add(5 * time.Second)
	t3 := t2.Add(5 * time.Second)
	t4 := t3.Add(5 * time.Second)
	lst.Add(t3)
	lst.Add(t4)
	lst.Add(t2)
	lst.Add(t1)

	el1 := lst.Front()
	el2 := el1.Next()
	el3 := el2.Next()
	el4 := el3.Next()
	assert.Equal(t, t1, el1.Value)
	assert.Equal(t, t2, el2.Value)
	assert.Equal(t, t3, el3.Value)
	assert.Equal(t, t4, el4.Value)
}

func TestAddingDuplicateShouldDoNothing(t *testing.T) {
	var lst OrderedListOfTimes
	t1 := time.Now()
	t2 := t1
	t3 := t1
	lst.Add(t1)
	lst.Add(t2)
	lst.Add(t3)
	assert.Equal(t, 1, lst.Len())
}

func getTestDataDiffTime() []*Candle {
	eurusd := Pair{
		Major: "EUR",
		Minor: "USD",
	}
	usdjpy := Pair{
		Major: "EUR",
		Minor: "USD",
	}
	t1 := time.Now()
	t2 := t1.Add(5 * time.Minute)

	c1 := Candle{
		Pair: eurusd,
		Time: ITime(t1),
	}

	c2 := Candle{
		Pair: usdjpy,
		Time: ITime(t2),
	}
	return []*Candle{&c1, &c2}
}
func TestAddToTimeSeries(t *testing.T) {
	var series TimeSeries
	candles := getTestDataDiffTime()
	candles[0].Time = candles[1].Time

	series.AddCandle(candles[0])
	series.AddCandle(candles[1])
	assert.Equal(t, 1, series.Len())
}

func TestAddToTimeSeriesDifferentTime(t *testing.T) {
	var series TimeSeries
	candles := getTestDataDiffTime()

	series.AddCandle(candles[0])
	series.AddCandle(candles[1])
	assert.Equal(t, 2, series.Len())
}

func TestIterationIsSuccesfull(t *testing.T) {
	var series TimeSeries
	candles := getTestDataDiffTime()
	series.AddCandle(candles[0])
	series.AddCandle(candles[1])
	for cs := series.Front(); cs != nil; cs = cs.Next() {
		key, ok := cs.Value.(time.Time)
		if ok != true {
			assert.Fail(t, "Zwrot z elementu musi byc time")
		}
		arr := series.Get(key)
		assert.NotNil(t, arr)
		assert.GreaterOrEqual(t, len(arr), 1)
	}
}
