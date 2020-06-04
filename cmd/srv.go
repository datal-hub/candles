package cmd

import (
	"github.com/urfave/cli"

	"github.com/datal-hub/candles/models"
	"github.com/datal-hub/candles/pkg/database"
	log "github.com/datal-hub/candles/pkg/logger"
	"github.com/datal-hub/candles/pkg/settings"
	"github.com/datal-hub/candles/pkg/truefx"
)

//description run the application from the command line in server mode
var Srv = cli.Command{
	Name:        "srv",
	Usage:       "Start candles",
	Description: `Candles service gets ticks and makes candles.`,
	Action:      runSrv,
}

func runSrv(c *cli.Context) {
	log.Info("Start runSrv")
	db, err := database.NewDB()
	if err != nil {
		log.Error(err.Error())
		return
	}
	tickers, err := db.GetOrCreateTickers(settings.Tickers)
	if err != nil {
		log.Error(err.Error())
		return
	}

	initCandles := make([]models.Candle, 0, len(tickers))
	for _, ticker := range tickers {
		initCandles = append(initCandles, models.Candle{TickerId: ticker.ID})
	}
	counter := make(map[int64]models.CandlesTimer, len(settings.Intervals))
	for _, interval := range settings.Intervals {
		counter[interval] = models.CandlesTimer{
			Timer:   interval,
			Candles: initCandles,
		}
	}
	for {
		candles, err := truefx.GetMinuteCandle(tickers)
		if err != nil {
			log.Error(err.Error())
		}
		for i, v := range counter {
			models.AddCandles(counter[i].Candles, candles)
			v.Timer--
			counter[i] = v
			if counter[i].Timer == 0 {
				for _, candle := range counter[i].Candles {
					candle.Interval = i
					if err := db.Save(&candle); err != nil {
						log.Error(err.Error())
					}
				}
				counter[i] = models.ResetTimer(v, i)

			}
		}
	}
}
