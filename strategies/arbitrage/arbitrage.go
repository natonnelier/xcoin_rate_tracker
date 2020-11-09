package arbitrage

import (
	"sort"
	"xcoin_rate_tracker/services"
)

var pairs = [...]string{"btcusd", "btceur", "ethusd", "etheur", "bchusd", "bcheur", "ltcusd", "ltceur", "eosusd", "eoseur"}

var margin = 5.0

type TickerGap struct {
	Pair        string
	AskExchange string
	AskPrice    float64
	BidExchange string
	BidPrice    float64
	Ratio       float64
}

func RateGap(pair string, compare bool) (*TickerGap, bool) {
	rates := services.GetRates(pair, false)
	var ask = rates.Asks[0]
	var bid = rates.Bids[0]
	var ratio = ((bid.Price / ask.Price) - 1) * 100
	gap := &TickerGap{
		Pair:        pair,
		AskExchange: ask.Exchange,
		AskPrice:    ask.Price,
		BidExchange: bid.Exchange,
		BidPrice:    bid.Price,
		Ratio:       ratio,
	}
	if compare || ratio >= margin {
		return gap, true
	} else {
		return nil, false
	}
}

func GetRateGaps(compare bool) []*TickerGap {
	var response = []*TickerGap{}
	for i := 0; i < len(pairs); i++ {
		gap, diff := RateGap(pairs[i], compare)
		if diff {
			response = append(response, gap)
		}
	}
	sort.Slice(response, func(i, j int) bool {
		return response[i].Ratio > response[j].Ratio
	})
	return response
}
