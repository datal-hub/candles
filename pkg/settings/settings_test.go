package settings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/datal-hub/candles/pkg/settings"
)

func TestDefaults(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("postgresql://xxx:xxxd@127.0.0.1:5432/candles?sslmode=disable", DB.Url)
	assert.Equal("localhost:8181", ListenAddr)
	assert.Equal([]int64{1, 3, 10}, Intervals)
	assert.Equal([]string{"EUR/USD", "USD/JPY"}, Tickers)
	assert.Equal(false, VerboseMode)
}

func TestFromFile(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	err := FromFile("testdata/test.ini")
	require.Nil(err)

	assert.Equal("postgresql://tuser:TPasSwd@testhost:4242/tdb", DB.Url)
	assert.Equal("0.0.0.0:443", ListenAddr)
	assert.Equal([]int64{1, 3, 5, 10}, Intervals)
	assert.Equal([]string{"EUR/USD", "USD/JPY"}, Tickers)
	assert.Equal(false, VerboseMode)
}
