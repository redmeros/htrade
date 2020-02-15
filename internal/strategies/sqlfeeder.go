package strategies

import (
	"github.com/redmeros/htrade/internal/db"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
)

var logger = logging.NewLogger("sqlfeeder.log")

// SQLFeeder jest konkretnym typem implementującym
// Feeder'a
type SQLFeeder struct {
	consumers []Consumer
	pairs     []*models.Pair
}

// Subscribe dodaje do listy subskrybentów
// consumera
func (f *SQLFeeder) Subscribe(consumer Consumer) {
	if consumer != nil {
		f.consumers = append(f.consumers, consumer)
	}
}

// Unsubscribe usuwa z listy subskrybentów konkretnego subskrybenta
func (f *SQLFeeder) Unsubscribe(consumer Consumer) {
	for i, x := range f.consumers {
		if x == consumer {
			f.consumers = append(f.consumers[:i], f.consumers[i+1:]...)
		}
	}
}

// Consumers zwraca listę aktualnych
// subskrybentów
func (f *SQLFeeder) Consumers() []Consumer {
	return f.consumers
}

func (f *SQLFeeder) AddPair(pair string) {
	d, err := db.TryGet()
	if err != nil {
		logger.Errorf("Cannot connect to db: %s", err.Error())
		return
	}
	mpair, err := models.GetPairByName(d, pair)
	if err != nil {
		logger.Errorf("Cannot get pair: %s", err.Error())
		return
	}
	f.pairs = append(f.pairs, mpair)
	// append()
}

// StartFeeding rozpoczyna karmienie danymi
func (f *SQLFeeder) StartFeeding() error {
	d, err := db.TryGet()
	if err != nil {
		return err
	}
	var candles []*models.Candle
	if err := d.Find(&candles).Error; err != nil {
		return err
	}

	var series models.TimeSeries
	for _, c := range candles {
		series.AddCandle(c)
	}

	for 

}
