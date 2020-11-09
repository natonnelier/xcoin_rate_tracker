package bitstamp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"xcoin_rate_tracker/rates"
)

const base_uri string = "https://www.bitstamp.net/api/"
const api_version string = "v2/"

var pairs = [...]string{"ltcusd", "ethusd", "xrpeur", "bchusd", "bcheur", "xrpbtc", "eurusd", "bchbtc", "ltceur", "btcusd", "btceur", "ltcbtc", "xrpusd", "ethbtc", "etheur", "eoseur", "eosusd"}

func TickerRate(pair string) *rates.Rate {
	uri := base_uri + api_version + "ticker/" + pair
	resp, err := http.Get(uri)
	var rate rates.Rate
	if err != nil {
		fmt.Println("Bitstamp TickerRate error: %s\n", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			var ticker rates.Bitstamp
			json.Unmarshal([]byte(body), &ticker)
			ticker.Pair = pair
			fmt.Printf("%+v\n", ticker)
			rate.Pair = pair
			if ask, err := strconv.ParseFloat(ticker.Ask, 64); err == nil {
				rate.Ask = ask
			}
			if bid, err := strconv.ParseFloat(ticker.Bid, 64); err == nil {
				rate.Bid = bid
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
