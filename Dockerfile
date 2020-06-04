# ---- Base Node ----
FROM golang:1.14.4 AS base

RUN mkdir -p /go/src/candles
WORKDIR /go/src/candles

ADD . /go/src/candles

# ---- Linter ----
FROM golangci/golangci-lint:latest

COPY --from=base /go/src/candles /go/src/candles
WORKDIR /go/src/candles

RUN golangci-lint run --timeout 3m0s -v .

# ---- Dependencies ----
FROM base AS dependencies

RUN go get -v





