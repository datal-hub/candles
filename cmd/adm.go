package cmd

import (
	"time"

	"github.com/urfave/cli"

	"github.com/datal-hub/candles/pkg/database"
	log "github.com/datal-hub/candles/pkg/logger"
)

//description run the application from the command line in administrator mode
var Adm = cli.Command{
	Name:        "adm",
	Usage:       "Administrative tools for candles service",
	Description: ``,
	Subcommands: []cli.Command{
		{
			Name:   "initdb",
			Usage:  "Initialize database",
			Action: initDB,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "Drop candles database!"},
			},
		},
	},
}

func initDB(c *cli.Context) error {
	var db database.DB
	var err error
	for {
		db, err = database.NewDB()
		if err != nil {
			log.ErrorF("Error database initialization", log.Fields{"message": err.Error()})
			time.Sleep(5 * time.Second)
			continue
		} else {
			break
		}
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()
	return db.Init(c.Bool("force"))
}
