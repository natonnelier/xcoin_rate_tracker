package rates


type Kraken struct {
  Pair       string
  Ask        float32  		`json:"a"`
  Bid        float32  		`json:"b"`
}
