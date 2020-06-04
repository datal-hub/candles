package models

//go:generate reform

//reform:tickers
type Ticker struct {
	ID    int64  `reform:"id,pk"`
	Label string `reform:"label"`
}
