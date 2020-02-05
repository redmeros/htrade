package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/redmeros/htrade/config"
	"github.com/redmeros/htrade/internal/db"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/internal/strategies"
	"github.com/redmeros/htrade/internal/strategies/algos"
	"github.com/redmeros/htrade/internal/strategies/brokers"
	"github.com/redmeros/htrade/internal/strategies/wallets"
	"github.com/redmeros/htrade/models"
)

// var logger := logging.NewLogger("tester.log")

func main() {
	logger := logging.NewLogger("tester.log")

	var pairs = []string{
		"EURUSD",
		"USDJPY",
	}

	cfg, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		logger.Fatalf("Cannot load config: %s", err.Error())
		return
	}

	db, err := db.GetDb(cfg)
	if err != nil {
		logger.Fatalf("Cannot get db %s", err.Error())
		return
	}

	fmt.Println("Starting building an index of times:")

	var dataset = make(map[models.ITime]map[string]*models.Candle)

	// var candles = make(map[string][]models.Candle)
	for _, pair := range pairs {
		pairCandles, err := models.GetCandlesByPairName(db, pair)
		if err != nil {
			logger.Fatalf("Cannot get candles for pair: %s, err: %s", pair, err)
			return
		}
		for _, candle := range pairCandles {

			if dataset[candle.Time] == nil {
				dataset[candle.Time] = make(map[string]*models.Candle)
			}
			dataset[candle.Time][pair] = &candle
		}
	}
	fmt.Println("Pairs read")
	fmt.Printf("Got %d vals in dataset\n", len(dataset))
	fmt.Println("Getting keys and sort them...")

	keys := make([]models.ITime, len(dataset))
	i := 0
	for key := range dataset {
		keys[i] = key
		i++
	}
	fmt.Printf("Got %d keys\n", len(keys))

	sort.Slice(keys, func(i, j int) bool {
		return time.Time(keys[i]).Before(time.Time(keys[j]))
	})
	fmt.Println("Dates are indexed now, now can try to provide values")

	fmt.Println("Ended building an index of times")

	fmt.Println("Building test ecosystem")

	manager := TestManager{
		algo:   new(algos.BuyAndHoldAglo),
		broker: new(brokers.TestBroker),
		wallet: new(wallets.TestWallet),
	}

	for _, k := range keys {
		manager.OnTick(dataset[k])
	}
}

// TestManager to manager testów
type TestManager struct {
	algo   strategies.IAlgo
	broker strategies.IBroker
	wallet strategies.IWallet
}

// OnTick realizuje strategie z zadanymi ustawionymi
// algorytmem brokerem i walletem
// 1. Najpierw realizowane są transakcje oczekujące (tam jest kolejka)
// 2. Na podstawie zrealizowanych transakcji updatowane są pozycje
// 3. Do algorytmu przekazywana jest informacja o nowych świeczkach czego wynikiem
// 	  jest informacja o wyniku
// 4. Na podstawie wyniku/ratingu dokonywana jest analiza portfela i dobierane są
// 	  wielkości pozycji zwracane są nowe zlecenia
// 5. Zlecenia przekazywane są do brokera w celu realizacji w następnej iteracji
// 6. Z metody zwracany jest wynik w postaci pozycji będących w portfelu, zleceń i rezultatów algorytmu
func (t *TestManager) OnTick(candles map[string]*models.Candle) ([]*strategies.Position, []*strategies.Order, []*strategies.AlgoResult) {
	realizedOrders := t.broker.RealizeOrders(candles)
	t.wallet.UpdatePositions(realizedOrders)

	result := t.algo.OnTick(candles)
	orders := t.wallet.FilterAlgosResult(result)
	t.broker.PushOrders(orders)
	return t.wallet.Positions(), orders, result
}
