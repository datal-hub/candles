package pgsql

import (
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/datal-hub/candles/models"
)

func (db *PgSQL) GetOrCreateTickers(labels []string) ([]models.Ticker, error) {
	rdb := reform.NewDB(db.SqlDB(), postgresql.Dialect, nil)
	tickers := make([]models.Ticker, 0, len(labels))
	for _, label := range labels {
		var ticker models.Ticker
		err := rdb.FindOneTo(&ticker, "label", label)
		if err != nil && err == reform.ErrNoRows {
			ticker.Label = label
			if err := rdb.Save(&ticker); err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
		tickers = append(tickers, ticker)
	}
	return tickers, nil
}
