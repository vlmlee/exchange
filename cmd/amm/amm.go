package amm

import (
	"time"

	orderbook "github.com/vlmlee/exchange/cmd/orderbook"
)

// Market makers picks up info from a feed and uses that data to find opportunities for various tickers
type MarketMaker struct {
	ID string
	Balance float32
	Profit float32
}

// Places a bid  order at the specific price and amount at the specific scheduleAhead time. Can be placed immediately.
func (m *MarketMaker) PlaceBid() {

}

// Places an ask order at the specific price and amount at the specific scheduleAhead time. Can be placed immediately.
func (m *MarketMaker) PlaceAsk(orderBook *orderbook.OrderBook, price float32, amount int, scheduleAhead time.Time) {

}