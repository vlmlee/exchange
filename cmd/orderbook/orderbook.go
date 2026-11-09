package orderbook

import (
	"time"

	"github.com/google/btree"
)

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
	Bids *btree.BTree
	Asks *btree.BTree
	Concavity float32 // A relative measure of liquidity
}