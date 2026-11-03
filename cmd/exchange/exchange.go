package exchange

import (
	orderbook "github.com/vlmlee/exchange/cmd/orderbook"
)

type Exchange struct {
	OrderBooks []*orderbook.OrderBook
}

func createExchange(orderbooks []*orderbook.OrderBook) *Exchange {
	return &Exchange{
		OrderBooks: orderbooks,
	}
}