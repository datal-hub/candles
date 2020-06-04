package settings

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

//DBType is the structure which define parameters for database connection
type DBType struct {
	Url string
}

//FXType is the structure which define parameters for FX source connection
type FXType struct {
	Url string
}

var (
	//ListenAddr is the address on which the server is running
	ListenAddr string
	// VerboseMode define logging mode
	// True - for verbose mode
	// False - for common mode
	VerboseMode bool
	//DB define parameters for database connection
	DB DBType
	//FX define parameters for FX source connection
	FX FXType
	//Intervals define parameters for candles intervals
	Intervals []int64
	//Tickers parameter define list of target values for saving
	Tickers []string
)

func init() {
	ListenAddr = "localhost:8181"
	VerboseMode = false
	DB.Url = "postgresql://xxx:xxxd@127.0.0.1:5432/candles?sslmode=disable"
	Intervals = []int64{1, 3, 10}
	Tickers = []string{"EUR/USD", "USD/JPY"}
	FX.Url = "https://webrates.truefx.com/rates/connect.html?id=jsTrader:string:ozrates:12345&f=csv"
}

//FromFile initialize settings from file
func FromFile(file string) error {
	log.Printf("Loading configuration from '%s'", file)

	cfg, err := ini.InsensitiveLoad(file)
	if err != nil {
		return fmt.Errorf("Error parsing file '%s'. Message: %s", file, err)
	}

	newListenAddr := cfg.Section("").Key("ListenAddr").String()
	newIntervals := cfg.Section("").Key("Intervals").Int64s(",")
	newTickers := cfg.Section("").Key("Tickers").Strings(",")

	var newDB DBType
	newDB.Url = cfg.Section("DB").Key("Url").String()

	var newFX FXType
	newFX.Url = cfg.Section("TrueFX").Key("Url").String()

	ListenAddr = newListenAddr
	Intervals = newIntervals
	Tickers = newTickers
	DB = newDB
	FX = newFX

	return nil
}
