package rates

type Dsx struct {
  Pair       string
  Ask        float64  		`json:"sell"`
  Bid        float64      `json:"buy"`
}
