package truefx

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/datal-hub/candles/models"
	log "github.com/datal-hub/candles/pkg/logger"
	"github.com/datal-hub/candles/pkg/settings"
)

const LoopCounter = 60

func tickInfoToCandle(tick []string, tickerId int64) models.Candle {
	if len(tick) != 9 {
		return models.Candle{}
	}
	time, err := strconv.Atoi(tick[1])
	if err != nil {
		return models.Candle{}
	}
	low, err := strconv.ParseFloat(tick[6], 8)
	if err != nil {
		return models.Candle{}
	}
	high, err := strconv.ParseFloat(tick[7], 8)
	if err != nil {
		return models.Candle{}
	}
	open, err := strconv.ParseFloat(tick[8], 8)
	if err != nil {
		return models.Candle{}
	}
	return models.Candle{
		High:      high,
		Low:       low,
		Open:      open,
		Close:     open,
		TickerId:  tickerId,
		BeginDttm: int64(time),
		EndDttm:   int64(time),
	}
}

func GetMinuteCandle(tickers []models.Ticker) ([]models.Candle, error) {
	candlesChan := make(chan []models.Candle, LoopCounter)
	tickersStr := ""
	tickersMap := make(map[string]int64)
	for i, ticker := range tickers {
		tickersStr += ticker.Label
		if i != len(tickers)-1 {
			tickersStr += ","
		}
		tickersMap[ticker.Label] = ticker.ID
	}
	requestURL := fmt.Sprintf("%s&c=%s", settings.FX.Url, tickersStr)
	counter := 0
	wg := sync.WaitGroup{}
	for counter < LoopCounter {
		wg.Add(1)
		go func() {
			res := make([]models.Candle, 0, len(tickers))
			defer wg.Done()
			resp, err := http.Get(requestURL)
			if err != nil {
				log.Error(err.Error())
				return
			}
			defer func() {
				if resp != nil && resp.Body != nil {
					resp.Body.Close()
				}
			}()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error(err.Error())
				return
			}
			tickInfo := csv.NewReader(strings.NewReader(string(body)))
			for {
				record, err := tickInfo.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Error(err.Error())
				}
				res = append(res, tickInfoToCandle(record, tickersMap[record[0]]))
			}
			candlesChan <- res
		}()
		time.Sleep(time.Second)
		counter++
	}
	wg.Wait()
	close(candlesChan)
	candles := <-candlesChan
	for curCandles := range candlesChan {
		models.AddCandles(candles, curCandles)
	}
	return candles, nil
}
