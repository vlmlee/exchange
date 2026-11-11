package exchange

import (
	orderbook "github.com/vlmlee/exchange/cmd/orderbook"
)

type Exchange struct {
	OrderBooks []*orderbook.OrderBook
}

func CreateExchange(orderbooks []*orderbook.OrderBook) *Exchange {
	return &Exchange{
		OrderBooks: orderbooks,
	}
}
