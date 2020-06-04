package models

//go:generate reform

//reform:candles
type Candle struct {
	ID        int64   `reform:"id,pk"`
	High      float64 `reform:"high"`
	Low       float64 `reform:"low"`
	Open      float64 `reform:"open"`
	Close     float64 `reform:"close"`
	TickerId  int64   `reform:"ticker_id"`
	BeginDttm int64   `reform:"begin_dttm"`
	EndDttm   int64   `reform:"end_dttm"`
	Interval  int64   `reform:"interval"`
}

func (c *Candle) add(sc Candle) {
	if (sc.Low != 0 && sc.Low < c.Low) || (c.Low == 0 && sc.Low != 0) {
		c.Low = sc.Low
	}
	if sc.High > c.High {
		c.High = sc.High
	}
	if sc.EndDttm > c.EndDttm {
		c.Close = sc.Close
		c.EndDttm = sc.EndDttm
	}
	if (c.BeginDttm == 0 && sc.BeginDttm != 0) || (sc.BeginDttm > 0 && sc.BeginDttm < c.BeginDttm) {
		c.Open = sc.Open
		c.BeginDttm = sc.BeginDttm
	}
	if c.TickerId == 0 {
		c.TickerId = sc.TickerId
	}
}

func AddCandles(fcc []Candle, scc []Candle) {
	for i := range fcc {
		for _, secondCandle := range scc {
			if fcc[i].TickerId == secondCandle.TickerId {
				fcc[i].add(secondCandle)
			}
		}
	}
}

type CandlesTimer struct {
	Timer   int64
	Candles []Candle
}

func ResetTimer(ct CandlesTimer, newTimer int64) CandlesTimer {
	newCandelsTimer := CandlesTimer{}
	newCandelsTimer.Timer = newTimer
	for _, c := range ct.Candles {
		newCandelsTimer.Candles = append(newCandelsTimer.Candles, Candle{TickerId: c.TickerId})
	}
	return newCandelsTimer
}
