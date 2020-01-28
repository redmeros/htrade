package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Oanda zawiera konfiguracje danych oandy
type Oanda struct {
	URL       string `json:"url"`
	Token     string `json:"token"`
	AccountID string `json:"account_id"`
}

// GetBearerToken zwraca wartość
// nagłówka Authorization
func (o *Oanda) GetBearerToken() string {
	return fmt.Sprintf("Bearer %s", o.Token)
}

// DB zawiera konfiguracje bazy danych
type DB struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Name     string `json:"name"`
}

// GetPgConnString zwraca connecction string
// obslugiwany przez gorm dla bazy postgres
func (c *DB) GetPgConnString() string {
	connstring := ""
	if len(c.User) != 0 {
		connstring += fmt.Sprintf("user=%s", c.User)
	}

	if len(c.Password) != 0 {
		connstring += fmt.Sprintf(" password=%s", c.Password)
	}

	if len(c.Host) != 0 {
		connstring += fmt.Sprintf(" host=%s", c.Host)
	}

	if c.Port != 0 {
		connstring += fmt.Sprintf(" port=%d", c.Port)
	}
	if len(c.Name) != 0 {
		connstring += fmt.Sprintf(" dbname=%s", c.Name)
	}

	return connstring
}

// WebConfig zawiera konfiguracje dla backendu
type WebConfig struct {
	BindingAddress string `json:"binding_address"`
	BindingPort    string `json:"binding_port"`
	Secret         string `json:"secret"`
	SignupBlocked  bool   `json:"signup_blocked"`
}

// FullAddress zwraca adres przekazywany
// do serwera
func (w *WebConfig) FullAddress() string {
	return fmt.Sprintf("%s:%s", w.BindingAddress, w.BindingPort)
}

// Config jest głównym structem
// konfiguracyjnym
type Config struct {
	Oanda Oanda     `json:"oanda"`
	Db    DB        `json:"db"`
	Web   WebConfig `json:"web"`
}

// LoadConfig ładuje plik filename
// jako struct Config
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
