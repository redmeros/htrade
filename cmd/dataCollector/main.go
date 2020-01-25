package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/redmeros/htrade/config"
	"github.com/redmeros/htrade/internal/db"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
	"github.com/redmeros/htrade/pkg/oanda"
)

func tryResloveConfig() (string, error) {
	files := []string{
		"config.json",
		"../../config/config.json",
	}
	for _, filename := range files {
		info, err := os.Stat(filename)
		if os.IsNotExist(err) {
			continue
		}
		if !info.IsDir() {
			return filename, nil
		}
	}
	return "", fmt.Errorf("Żaden ze standardowych plikow nie istnieje")
}

func main() {
	logger := logging.NewLogger()
	defer logging.Shutdown()
	var configArgIdx = -1
	var configFileName string
	for i, el := range os.Args {
		if el == "-c" {
			configArgIdx = i
			break
		}
	}

	if configArgIdx == -1 {
		resolvedConfigFileName, err := tryResloveConfig()
		if err != nil {
			logger.Fatal("Nie znalazłem żadnego pliku config")
		}
		configFileName = resolvedConfigFileName
	}

	config, err := config.LoadConfig(configFileName)
	if err != nil {
		logger.Fatalf("Nie moge zaladowac pliku konfiguracyjnego, %s", err.Error())
	}

	if _, err = db.GetDb(config); err != nil {
		logger.Fatalf("Cannot get db: %s", err.Error())
		return
	}

	pairs, err := readPairs()

	if err != nil {
		logger.Fatalf("Cannot read pairs: %s", err.Error())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				logger.Info("ctrl+c has been discovered quiting")
				d, err := db.GetDB()
				if err == nil {
					d.Close()
				}
				os.Exit(0)
			}
		}
	}()

	d := time.Second * 60 * 5
	for range time.Tick(d) {
		overallDo(pairs, config)
	}
}

func overallDo(pairs []*models.Pair, config *config.Config) {
	logger := logging.NewLogger()
	var wg sync.WaitGroup
	for _, pair := range pairs {
		wg.Add(1)
		go doJob(config, pair, &wg)
	}
	logger.Info("Waiting for goroutines to be done...")
	wg.Wait()
	logger.Info("Goroutines done")
}

func readPairs() ([]*models.Pair, error) {
	mdb, err := db.GetDB()
	logger := logging.NewLogger()
	if err != nil {
		logger.Errorf("Cannot get pairs: %s", err.Error())
		return nil, err
	}
	pairs := []*models.Pair{}
	if err := mdb.Find(&pairs).Error; err != nil {
		logger.Errorf("Database error: %s", err.Error())
		return nil, err
	}
	return pairs, nil
}

func doJob(cfg *config.Config, pair *models.Pair, wg *sync.WaitGroup) error {
	defer wg.Done()
	var params map[string]string = make(map[string]string)
	params["price"] = "BA"
	params["count"] = "2"
	params["granularity"] = "M5"

	oanda := oanda.NewOanda(&cfg.Oanda)
	logger := logging.NewLogger()
	logger.Infof("Starting job for %s", pair.Name())

	candle, err := oanda.GetCandles(pair.NameWithSep("_"), params)

	if err != nil {
		logger.Fatal(err.Error())
		return err
	}

	mdb, err := db.GetDb(cfg)
	if err != nil {
		logger.Fatalf("Cannot create connection with db: %s", err.Error())
		return err
	}

	mcs, err := candle.ToCandle(pair)
	if err != nil {
		logger.Errorf("Error during conversion to pair: %s", err.Error())
		return err
	}

	tx := mdb.Begin()
	for _, c := range mcs {
		if err := tx.Create(c).Error; err != nil {
			logger.Errorf("Error during updating db: %s", err.Error())
			tx.Rollback()
			return err
		}
		logger.Infof("Created row for %s, at %s", c.Pair.Name(), c.Time.String())
	}
	tx.Commit()
	return nil
}
