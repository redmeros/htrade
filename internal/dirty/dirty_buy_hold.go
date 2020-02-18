package dirty

import (
	"github.com/redmeros/htrade/internal/db"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
)

var logger = logging.NewLogger("dirty.log")

func RunBuyHold(initialMoney float64, leverage float64) {
	d, err := db.TryGet()
	if err != nil {
		logger.Error(err)
		panic(err)
	}

	candles, err := models.GetCandlesByPairName(d, "EURUSD")
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	if len(candles) == 0 {
		panic("len of candles cannot be 0")
	}

	bought := false

	pairby, err := models.GetPairByName(d, "EURUSD")
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	var position *Position
	for i, candle := range candles {
		if !bought {
			bought = true
			position = NewPosition(pairby, 10000, 1, candle.CloseAsk)
		}
		if i == len(candles)-1 {
			logger.Info(position.PL(candle.CloseBid, 1))
		}
	}

}
