package db

import (
	"errors"

	"github.com/jinzhu/gorm"

	// to jest zeby postgras  dzialal
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/redmeros/htrade/config"
	"github.com/redmeros/htrade/internal/logging"
	"github.com/redmeros/htrade/models"
)

var db *gorm.DB = nil

// var sslmode = true
var sslmode = false

// GetDB zwraca otwarta baze danych,
// jesli nie ma dostepnej bazy danych zwraca blad,
// jesli zwraca blad uzyj GetDb razem z configiem
// zeby stworzyc nowe polaczenie
func GetDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	return nil, errors.New("No db to get")
}

// Zwraca baze danych, jeśli nei jest zainicjalizowana
// próbuje ją zainicjować
func TryGet() (*gorm.DB, error) {
	d, err := GetDB()
	if err == nil {
		return d, err
	}
	config, err := config.TryLoad()
	if err != nil {
		return nil, err
	}
	return GetDb(config)
}

// GetDb zwraca wskaznik do bazy danych
func GetDb(config *config.Config) (*gorm.DB, error) {
	log := logging.NewLogger("database.log")
	if db == nil {
		cstr := config.Db.GetPgConnString()
		if sslmode == false {
			cstr = cstr + " sslmode=disable"
		}
		log.Debugf("Creating new db object with: %s", cstr)
		var err error
		db, err = gorm.Open("postgres", cstr)
		if err != nil {
			return nil, err
		}
		log.Debug("Migrations...")
		// log.Debug("Dropping tables...")
		// db.DropTableIfExists(&models.Candle{}, &models.Pair{})
		// db.DropTableIfExists(&models.Candle{})
		if db.Error != nil {
			return nil, db.Error
		}
		log.Debug("Automigrating...")
		db.AutoMigrate(
			&models.Pair{},
			&models.Candle{},
			&models.User{},
		)
		if db.Error != nil {
			return nil, db.Error
		}
		log.Debug("Adding foreign key...")
		db.Model(&models.Candle{}).AddForeignKey("pair_id", "pairs(id)", "RESTRICT", "RESTRICT")
		if db.Error != nil {
			return nil, db.Error
		}
		log.Debug("All migrations done.")

		return db, nil
	}
	return db, nil
}
