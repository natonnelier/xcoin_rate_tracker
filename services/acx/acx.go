package acx

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "xcoin_rate_tracker/rates"
)

const baseUri string = "https://acx.io/api/"
const apiVersion string = "v2/"
var pairs = [...]string {"btcaud","ethaud","bchaud","ltcaud","hsraud","usdtaud","btcusdt","ethusdt","dashbtc","fuelbtc","ubtcbtc"}

type Ticker struct {
  Key rates.Acx  `json:"ticker"`
}

func TickerRate(pair string) *rates.Rate {
  uri := baseUri + apiVersion + "tickers/" + pair
  resp, err := http.Get(uri)
  var rate rates.Rate
  if err != nil {
    fmt.Println("Acx TickerRate error: %s\n", err)
  } else {
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Printf("%s\n", err)
    } else {
      var ticker Ticker
      json.Unmarshal([]byte(body), &ticker)
      ticker.Key.Pair = pair
      fmt.Println("Acx TickerRate: %s\n")
      fmt.Printf("%+v\n", ticker)
      rate.Pair = pair
      rate.Ask = ticker.Key.Ask
      rate.Bid = ticker.Key.Bid
    }
  }
  return &rate
}

func AllRates() {
  for i := 0; i < len(pairs); i++ {
    TickerRate(pairs[i])
  }
}
