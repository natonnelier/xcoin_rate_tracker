package rates

type Acx struct {
  Pair       string
  Ask        float64  		`json:"sell"`
  Bid        float64			`json:"buy"`
}
