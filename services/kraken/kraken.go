package kraken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"xcoin_rate_tracker/rates"
)

const base_uri string = "https://api.kraken.com/"
const api_version string = "0"

var pairs = [...]string{"BCHEUR", "BCHUSD", "BCHXBT", "DASHEUR", "DASHUSD", "DASHXBT", "EOSETH", "EOSEUR", "EOSUSD", "EOSXBT", "GNOETH", "GNOEUR",
	"GNOUSD", "GNOXBT", "USDTZUSD", "XETCXETH", "XETCZEUR", "XETCZUSD", "XETHZCAD", "XETHZEUR",
	"XETHZGBP", "XETHZJPY", "XICNXETH", "XICNXXBT", "XLTCXXBT", "XLTCZEUR", "XLTCZUSD",
	"XMLNXETH", "XMLNXXBT", "XREPXETH", "XREPXXBT", "XREPZEUR", "XREPZUSD", "XXBTZCAD", "XXBTZEUR", "XXBTZGBP",
	"XXBTZJPY", "XXBTZUSD", "XXDGXXBT", "XXLMXXBT", "XXLMZEUR", "XXLMZUSD", "XXMRXXBT", "XXMRZEUR",
	"XXMRZUSD", "XXRPXXBT", "XXRPZCAD", "XXRPZEUR", "XXRPZJPY", "XXRPZUSD", "XZECXXBT", "XZECZEUR", "XZECZJPY", "XZECZUSD"}

type Ticker struct {
	Kraken map[string]Kraken `json:"result"`
}

type Kraken struct {
	Ask []string `json:"a"`
	Bid []string `json:"b"`
}

func TickerRate(pair string) *rates.Rate {
	normalizedpair := normalizePair(pair)
	uri := base_uri + api_version + "/public/" + "Ticker" + "?pair=" + normalizedpair
	resp, err := http.Get(uri)
	var rate rates.Rate
	if err != nil {
		fmt.Println("Kraken TickerRate error: %s\n", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			var ticker Ticker
			json.Unmarshal([]byte(body), &ticker)
			fmt.Printf("%+v\n", ticker)
			// var p = strings.ToUpper(pair)
			if len(ticker.Kraken[normalizedpair].Ask) > 0 {
				rate.Pair = pair
				if ask, err := strconv.ParseFloat(ticker.Kraken[normalizedpair].Ask[0], 64); err == nil {
					rate.Ask = ask
				}
				if bid, err := strconv.ParseFloat(ticker.Kraken[normalizedpair].Bid[0], 64); err == nil {
					rate.Bid = bid
				}
			}
		}
	}
	return &rate
}

func AllRates() {
	for i := 0; i < len(pairs); i++ {
		TickerRate(pairs[i])
	}
}

func normalizePair(pair string) string {
	switch pair {
	case "btceur":
		return "XXBTZEUR"
	case "btcusd":
		return "XXBTZUSD"
	case "btcgbp":
		return "XXBTZGBP"
	case "etheur":
		return "XETHZEUR"
	case "ethusd":
		return "XETHZUSD"
	case "ethgbp":
		return "XETHZGBP"
	case "bcheur":
		return "BCHEUR"
	case "bchusd":
		return "BCHUSD"
	case "bchbtc":
		return "BCHXBT"
	case "ltceur":
		return "XLTCZEUR"
	case "ltcusd":
		return "XLTCZUSD"
	case "ltcbtc":
		return "XLTCXXBT"
	case "eosusd":
		return "EOSUSD"
	case "eoseur":
		return "EOSEUR"
	}
	return ""
}
