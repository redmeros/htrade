package config

import "io/ioutil"

import "encoding/json"

import "fmt"

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
	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s", c.Host, c.User, c.Name, c.Password)
}

// WebConfig zawiera konfiguracje dla backendu
type WebConfig struct {
	BindingAddress string `json:"binding_address"`
	BindingPort    string `json:"binding_port"`
	Secret         string `json:"secret"`
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
