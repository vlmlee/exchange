package orderbook

import "time"

type Bid struct {
	ID int
	BidPrice float32
	Amount int // No fractional orders allowed
	Expiration time.Time
}

type Ask struct {
	ID int
	AskPrice float32
	Amount int
	Expiration time.Time
}

type OrderBook struct {
	Ticker string
	Bids []*Bid
	Asks []*Ask
	Concavity float32 // A relative measure of liquidity
}