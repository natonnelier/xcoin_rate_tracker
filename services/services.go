package services

import (
	"fmt"
	"sort"
	"xcoin_rate_tracker/rates"
	"xcoin_rate_tracker/services/acx"
	"xcoin_rate_tracker/services/bitstamp"
	"xcoin_rate_tracker/services/dsx"
	"xcoin_rate_tracker/services/kraken"
)

var exchanges = []string{"dsx", "kraken"}

type Ticker struct {
	Exchange string
	Price    float64
}

type Rates struct {
	Bids []Ticker
	Asks []Ticker
}

func ExchangeTickerRate(pair string, exchange string) *rates.Rate {
	switch exchange {
	case "acx":
		return acx.TickerRate(pair)
	case "bitstamp":
		return bitstamp.TickerRate(pair)
	case "dsx":
		return dsx.TickerRate(pair)
	case "kraken":
		return kraken.TickerRate(pair)
	}
	var rate rates.Rate
	return &rate
}

func AllRates(exchange string) {
	switch exchange {
	case "acx":
		acx.AllRates()
	case "bitstamp":
		bitstamp.AllRates()
	case "dsx":
		dsx.AllRates()
	case "kraken":
		kraken.AllRates()
	}
}

func GetRates(pair string, compare bool) *Rates {
	rates := &Rates{}
	if compare {
		exchanges = append(exchanges, "bitstamp")
	}

	for i := 0; i < len(exchanges); i++ {
		rates.AppendExchangeRates(exchanges[i], pair)
	}

	sort.Slice(rates.Asks, func(i, j int) bool {
		return rates.Asks[i].Price < rates.Asks[j].Price
	})

	sort.Slice(rates.Bids, func(i, j int) bool {
		return rates.Bids[i].Price > rates.Bids[j].Price
	})

	return rates
}

func (rates *Rates) AppendExchangeRates(exchange string, pair string) {
	exchanges := ExchangeTickerRate(pair, exchange)
	if exchanges.Ask == 0 {
		fmt.Println(exchange + " empty")
	} else {
		var bid Ticker
		bid.Exchange = exchange
		bid.Price = exchanges.Bid
		var ask Ticker
		ask.Exchange = exchange
		ask.Price = exchanges.Ask
		rates.Bids = append(rates.Bids, bid)
		rates.Asks = append(rates.Asks, ask)
	}
}
