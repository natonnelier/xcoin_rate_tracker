package dsx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"xcoin_rate_tracker/rates"
)

const base_uri string = "https://dsxglobal.com/mapi/v2/"

var pairs = [...]string{"btgeur", "btcusdt", "btcrub", "bcceurs", "btcgbp", "eurseur", "ltcbtc", "bccusdt", "btggbp", "eursusd", "ethgbp",
	"btgbtc", "usdrub", "bccbtc", "ltcusd", "bccusd", "btcusd", "ltceur", "bcceur", "btceur", "btgusd", "ethusd", "etheur",
	"ltcgbp", "gbpusd", "bccgbp", "ethbtc", "ltcusdt", "btceurs", "usdtusd", "ltceurs", "usdteur", "etheurs", "eurusd", "ethusdt", "eoseur", "eosusd"}

type Ticker struct {
	Dsx Dsx `json:"-"`
}

type Dsx struct {
	Pair string
	Ask  string `json:"sell"`
	Bid  string `json:"buy"`
}

func TickerRate(pair string) *rates.Rate {
	uri := base_uri + "ticker/" + pair
	resp, err := http.Get(uri)
	var rate rates.Rate
	fmt.Println("%s\n", pair)
	if err != nil {
		fmt.Println("DSX TickerRate error: %s\n", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			var wrapper map[string]rates.Dsx
			json.Unmarshal([]byte(body), &wrapper)
			fmt.Printf("%+v\n", wrapper[pair].Ask)
			rate.Pair = pair
			rate.Ask = wrapper[pair].Ask
			rate.Bid = wrapper[pair].Bid
		}
	}
	return &rate
}

func AllRates() {
	for i := 0; i < len(pairs); i++ {
		TickerRate(pairs[i])
	}
}
