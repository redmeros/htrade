package oanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/redmeros/htrade/config"
)

// Oanda zawiera metody do wymiany danych z oanda
type Oanda struct {
	config *config.Oanda
}

// GetCandles pobiera z oandy aktualne dane
func (o *Oanda) GetCandles(pair string, params map[string]string) (*CandleResponse, error) {
	url := fmt.Sprintf("%s/v3/instruments/%s/candles", o.config.URL, pair)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", o.config.GetBearerToken())
	buildURL(req.URL, params)
	// spew.Dump(req.URL.String())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Status zwróconego kodu to: %d \r\n body to: \r\n %s", resp.StatusCode, string(bytes))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// spew.Dump(string(body))

	candleresp := CandleResponse{}
	err = json.Unmarshal(body, &candleresp)
	if err != nil {
		return nil, err
	}
	return &candleresp, nil
}

// NewOanda tworzy nowa instancje oandy
func NewOanda(config *config.Oanda) Oanda {
	var oanda Oanda
	oanda.config = config
	return oanda
}

func buildURL(url *url.URL, params map[string]string) error {
	q := url.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	url.RawQuery = q.Encode()
	return nil
}
