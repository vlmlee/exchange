package orderbook

import (
	"time"

	"github.com/google/btree"
)

type Order struct {
	ID         int
	LimitPrice float64
	Side       string
	Amount     int // No fractional orders allowed
	Expiration time.Time
}

type OrderBook struct {
	Ticker         string
	Bids           *btree.BTree
	Asks           *btree.BTree
	Convexity      float64 // A relative measure of liquidity
	LastTradePrice float64
}

func NewOrderBook(ticker string) *OrderBook {
	return &OrderBook{
		Ticker:         ticker,
		Bids:           btree.New(3),
		Asks:           btree.New(3),
		Convexity:      0,
		LastTradePrice: 0,
	}
}
