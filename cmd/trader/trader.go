package trader

type Ack = interface{}

type Trader interface {
	PlaceBid() Ack
	PlaceAsk() Ack
}

func PlaceBid(t *Trader) Ack {
	return nil
}

func PlaceAsk(t *Trader) Ack {
	return nil
}