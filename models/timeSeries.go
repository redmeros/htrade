package models

import (
	"container/list"
	"fmt"
	"time"
)

// TimeSeries to podstawowa struktura danych
// dla serii danych opartych na czasie
type TimeSeries struct {
	keys OrderedListOfTimes
	rows map[time.Time][]*Candle
}

// Len zwraca liczbę wierszy w serii, panikuje jeśli
// liczba kluczy jest inna od liczby wierszy
func (t *TimeSeries) Len() int {
	if t.keys.Len() != len(t.rows) {
		msg := fmt.Sprintf("t.keys.len() : %d, len(t.rows) : %d", t.keys.Len(), len(t.rows))
		panic(msg)
	}
	return t.keys.Len()
}

// Get zwraca Swieczki odpowiadajace dacie
func (t *TimeSeries) Get(idx time.Time) []*Candle {
	return t.rows[idx]
}

// Front zwraca klucz początkowy serii
func (t *TimeSeries) Front() *list.Element {
	return t.keys.Front()
}

// AddCandle dodaje świeczkę do serii
func (t *TimeSeries) AddCandle(candle *Candle) {
	candleTime := time.Time(candle.Time)
	t.keys.Add(candleTime)
	if t.rows == nil {
		t.rows = make(map[time.Time][]*Candle, 0)
	}

	arr := t.rows[candleTime]
	if arr == nil {
		arr = make([]*Candle, 0)
	}
	arr = append(arr, candle)
	t.rows[candleTime] = arr
}

// OrderedListOfTimes jest listą która jest
// zawsze posortowana służy jako lista kluczy dla
// TimeSeries
type OrderedListOfTimes struct {
	times list.List
}

// Add dodaje time.Time do serii, przy czym
// Ustawia je w kolejności przy czym zaczyna od tyłu
// (zakładam że z bazy dane będą w miarę uporządkowane)
// jeśli t już jest w liście - nic się nie dzieje
func (c *OrderedListOfTimes) Add(t time.Time) {
	for e := c.times.Back(); e != nil; e = e.Prev() {
		w := e.Value.(time.Time)
		if t == w {
			return
		}
		if t.After(w) {

			c.times.InsertAfter(t, e)
			return
		}
	}
	c.times.PushFront(t)
}

// Front jest wrapper od list Front
func (c *OrderedListOfTimes) Front() *list.Element {
	return c.times.Front()
}

// Len jest wrapperrem od list len
func (c *OrderedListOfTimes) Len() int {
	return c.times.Len()
}
