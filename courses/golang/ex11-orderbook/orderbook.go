package orderbook

type Orderbook struct {
	Ask []*Order
	Bid []*Order
}

func New() *Orderbook {
	return &Orderbook{}
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	switch order.Side {
	case SideBid:
		return orderbook.onSideBid(order)
	case SideAsk:
		return orderbook.onSideAsk(order)
	}
	return nil, nil
}

func (orderbook *Orderbook) onSideBid(order *Order) ([]*Trade, *Order) {
	var trades []*Trade
	for order.Volume > 0 {
		var min uint64
		min = 0
		count := 0
		var currentAsk *Order

		//check if there is any ask
		if len(orderbook.Ask) <= 0 && order.Kind == KindLimit {
			orderbook.Bid = append(orderbook.Bid, order)
			return trades, nil
		}

		//look for best price
		for i, ask := range orderbook.Ask {
			if (ask.Price <= order.Price || order.Kind == KindMarket) && (min > ask.Price || min == 0) {
				min = ask.Price
				currentAsk = ask
				count = i + 1
			}
		}

		if count == 0 && order.Kind == KindMarket {
			return trades, order
		}

		//no match
		if count == 0 {
			orderbook.Bid = append(orderbook.Bid, order)
			return trades, nil
		}

		//order fullfilled, ask isn't
		if order.Volume < currentAsk.Volume {
			currentAsk.Volume = currentAsk.Volume - order.Volume
			orderbook.Ask[count-1] = currentAsk
			trades = createTrade(trades, order.Volume, currentAsk.Price, currentAsk, order)
			return trades, nil
		}

		//order fullfilled
		if order.Volume == currentAsk.Volume {
			trades = createTrade(trades, currentAsk.Volume, currentAsk.Price, currentAsk, order)
			orderbook.removeItem(count, SideAsk)
			return trades, nil
		}

		if order.Volume > currentAsk.Volume {
			trades = createTrade(trades, currentAsk.Volume, currentAsk.Price, currentAsk, order)
			order.Volume = order.Volume - currentAsk.Volume
			orderbook.removeItem(count, SideAsk)
		}
	}
	return trades, nil
}

func (orderbook *Orderbook) onSideAsk(order *Order) ([]*Trade, *Order) {
	var trades []*Trade
	for order.Volume > 0 {
		var max uint64
		max = 0
		count := 0
		var currentBid *Order

		//check if there is any bid
		if len(orderbook.Bid) <= 0 && order.Kind == KindLimit {
			orderbook.Ask = append(orderbook.Ask, order)
			return trades, nil
		}

		//look for best price
		for i, bid := range orderbook.Bid {
			if (bid.Price >= order.Price || order.Kind == KindMarket) && max < bid.Price {
				max = bid.Price
				currentBid = bid
				count = i + 1
			}
		}

		if count == 0 && order.Kind == KindMarket {
			return trades, order
		}

		//no match
		if count == 0 {
			orderbook.Ask = append(orderbook.Ask, order)
			return trades, nil
		}

		//order fullfilled, bid isn't
		if order.Volume < currentBid.Volume {
			currentBid.Volume = currentBid.Volume - order.Volume
			orderbook.Bid[count-1] = currentBid
			trades = createTrade(trades, order.Volume, currentBid.Price, currentBid, order)
			return trades, nil
		}

		//order fullfilled
		if order.Volume == currentBid.Volume {
			trades = createTrade(trades, currentBid.Volume, currentBid.Price, currentBid, order)
			orderbook.removeItem(count, SideBid)
			return trades, nil
		}

		if order.Volume > currentBid.Volume {
			trades = createTrade(trades, currentBid.Volume, currentBid.Price, currentBid, order)
			order.Volume = order.Volume - currentBid.Volume
			orderbook.removeItem(count, SideBid)
		}
	}
	return trades, nil
}

func createTrade(trades []*Trade, amount uint64, price uint64, bid *Order, ask *Order) (output []*Trade) {
	trade := &Trade{
		Price:  price,
		Bid:    bid,
		Ask:    ask,
		Volume: amount,
	}
	trades = append(trades, trade)
	return trades
}

func (orderbook *Orderbook) removeItem(i int, side Side) {
	i--
	switch side {
	case SideBid:
		copy(orderbook.Bid[i:], orderbook.Bid[i+1:])
		orderbook.Bid[len(orderbook.Bid)-1] = nil
		orderbook.Bid = orderbook.Bid[:len(orderbook.Bid)-1]
	case SideAsk:
		copy(orderbook.Ask[i:], orderbook.Ask[i+1:])
		orderbook.Ask[len(orderbook.Ask)-1] = nil
		orderbook.Ask = orderbook.Ask[:len(orderbook.Ask)-1]
	}
}
