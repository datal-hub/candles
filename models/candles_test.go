package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var emptyOne = Candle{
	TickerId: 1,
}

var emptyTwo = Candle{
	TickerId: 2,
}

var candleOne = Candle{
	ID:        0,
	High:      10,
	Low:       1,
	Open:      2,
	Close:     5,
	TickerId:  1,
	BeginDttm: 1591265436967,
	EndDttm:   1591265436910,
}

var candleTwo = Candle{
	ID:        0,
	High:      10,
	Low:       1,
	Open:      2,
	Close:     5,
	TickerId:  2,
	BeginDttm: 1591265436967,
	EndDttm:   1591265436910,
}

var firstTickerOne = Candle{
	ID:        0,
	High:      5,
	Low:       2,
	Open:      3,
	Close:     4,
	TickerId:  1,
	BeginDttm: 1591265436967,
	EndDttm:   1591265436980,
}

var secondTickerOne = Candle{
	ID:        0,
	High:      10,
	Low:       1,
	Open:      2,
	Close:     5,
	TickerId:  1,
	BeginDttm: 1591265436980,
	EndDttm:   1591265437000,
}

var firstTickerTwo = Candle{
	ID:        0,
	High:      5,
	Low:       2,
	Open:      3,
	Close:     4,
	TickerId:  2,
	BeginDttm: 1591265436967,
	EndDttm:   1591265436980,
}

var secondTickerTwo = Candle{
	ID:        0,
	High:      10,
	Low:       1,
	Open:      2,
	Close:     5,
	TickerId:  2,
	BeginDttm: 1591265436980,
	EndDttm:   1591265437000,
}

func TestAddToEmpty(t *testing.T) {
	emptyOne.add(candleOne)
	assert.Exactly(t, emptyOne, candleOne, "TestAddToEmpty OK")
}

func TestAddNotEmptyCandles(t *testing.T) {
	firstTickerOne.add(secondTickerOne)
	expected := Candle{
		High:      secondTickerOne.High,
		Low:       secondTickerOne.Low,
		Open:      firstTickerOne.Open,
		Close:     secondTickerOne.Close,
		TickerId:  firstTickerOne.TickerId,
		BeginDttm: firstTickerOne.BeginDttm,
		EndDttm:   secondTickerOne.EndDttm,
	}
	assert.Exactly(t, firstTickerOne, expected, "TestAddNotEmptyCandles OK")
}

func TestAddCandlesOK(t *testing.T) {
	expectedTickerOne := Candle{
		High:      secondTickerOne.High,
		Low:       secondTickerOne.Low,
		Open:      firstTickerOne.Open,
		Close:     secondTickerOne.Close,
		TickerId:  1,
		BeginDttm: firstTickerOne.BeginDttm,
		EndDttm:   secondTickerOne.EndDttm,
	}
	expectedTickerTwo := Candle{
		High:      secondTickerOne.High,
		Low:       secondTickerOne.Low,
		Open:      firstTickerOne.Open,
		Close:     secondTickerOne.Close,
		TickerId:  2,
		BeginDttm: firstTickerOne.BeginDttm,
		EndDttm:   secondTickerOne.EndDttm,
	}
	fcc := []Candle{firstTickerOne, firstTickerTwo}
	scc := []Candle{secondTickerOne, secondTickerTwo}
	AddCandles(fcc, scc)
	assert.Exactly(t, scc, scc)
	assert.Exactly(t, fcc, []Candle{expectedTickerOne, expectedTickerTwo})
}

func TestAddCandlesWithEmpty(t *testing.T) {
	fcc := []Candle{emptyOne, emptyTwo}
	scc := []Candle{candleOne, candleTwo}
	AddCandles(fcc, scc)
	assert.Exactly(t, scc, scc)
	assert.Exactly(t, fcc, fcc)
}

func TestResetTimer(t *testing.T) {
	ct := CandlesTimer{
		Timer:   0,
		Candles: []Candle{candleOne, candleTwo},
	}
	resetedTimer := ResetTimer(ct, 3)
	assert.Exactly(t, CandlesTimer{Timer: 3, Candles: []Candle{emptyOne, emptyTwo}}, resetedTimer)
}
