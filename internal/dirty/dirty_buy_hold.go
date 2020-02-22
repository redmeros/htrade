package dirty

import (
	"time"

	"github.com/redmeros/htrade/internal/db"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
)

var logger = logging.NewLogger("dirty.log")

// NewBuyAndHoldStrategy returns new strategy
func NewBuyAndHoldStrategy() models.IStrategy {
	s := &BuyHoldStrategy{}
	s.SetupDefault()
	return s
}

// BuyHoldStrategy to struct odpowiadający za realizację strategii
type BuyHoldStrategy struct {
	models.StrategyBase
}

// SetupDefault ustawia domyślne wartości dla strategii
func (s *BuyHoldStrategy) SetupDefault() {
	s.SetCodeName("buyhold")
	s.SetDisplayName("Buy & Hold")
	s.SetDescription("Strategia Buy&Hold to tak naprawdę testowa strategia która sprawdza czy silnik działa.")
	params := s.InputParameters()
	initialMoney := models.NewFloatParameter("Wartość startowa portfela", "initial_money", 10000)
	leverage := models.NewFloatParameter("Lewar", "leverage", 100)
	params = append(params, initialMoney)
	params = append(params, leverage)

	db, err := db.TryGet()
	if err == nil {
		var minDate time.Time
		row := db.Table("candles").Select("Min(time)").Row()
		if err = row.Scan(&minDate); err == nil {
			params = append(params, models.NewDateParameter("Data początkowa", "date_from", minDate.Year(), int(minDate.Month()), minDate.Day()))
		}

		var maxDate time.Time
		row = db.Table("candles").Select("Max(time)").Row()
		if err = row.Scan(&maxDate); err == nil {
			params = append(params, models.NewDateParameter("Data końcowa", "date_to", maxDate.Year(), int(maxDate.Month()), maxDate.Day()))
		}
	}

	s.SetInputParameters(params)
}

// func (s *BuyHoldStrategy) MarshallJson

// Run Uruchamia strategię
func (s *BuyHoldStrategy) Run() (*models.PortfolioResult, error) {
	d, err := db.TryGet()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	candles, err := models.GetCandlesByPairName(d, "EURUSD")
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if len(candles) == 0 {
		return nil, err
	}

	bought := false

	pairby, err := models.GetPairByName(d, "EURUSD")
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var position *models.Position
	var results models.PortfolioResult
	for i, candle := range candles {
		if !bought {
			bought = true
			position = models.NewPosition(pairby, 10000, 1, candle.OpenAsk, time.Time(candle.Time))
		}
		if i == len(candles)-1 {
			position.Close(candle.CloseBid, 1.0/candle.CloseBid, time.Time(candle.Time))
			logger.Info(position.PL(candle.CloseBid, 1))
		}
		rec := models.PortfolioRecord{
			Value: 10000.0 + position.PL(candle.CloseBid, 1.0/candle.CloseBid),
		}
		results.AddRecord(&rec)
	}
	results.AddPosition(position)
	return &results, nil
}
