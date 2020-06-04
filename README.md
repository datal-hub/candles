# CANDLES

This project provide getting currency ticks using API truefx.com
and saving candles to Postgres for set intervals.

## Init database

```bash
go run candles.go -vv -c ./examples/cfg.ini adm initdb
```

## Run srv mode

```bash
go run candles.go -vv -c ./examples/cfg.ini srv
```

## Running with docker-compose

```bash
docker-compose up 
```
