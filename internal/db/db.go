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

// GetDB zwraca otwarta baze danych,
// jesli nie ma dostepnej bazy danych zwraca blad,
// jesli zwraca blad uzy GetDb razem z configiem
// zeby stworzyc nowe polaczenie
func GetDB() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	return nil, errors.New("No db to get")
}

// GetDb zwraca wskaznik do bazy danych
func GetDb(config *config.Config) (*gorm.DB, error) {
	log := logging.NewLogger("database.log")
	if db == nil {
		log.Debugf("Creating new db object with: %s", config.Db.GetPgConnString())
		var err error
		db, err = gorm.Open("postgres", config.Db.GetPgConnString())
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
