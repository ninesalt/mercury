package main

import (
	"errors"
	"fmt"
)

// TODO: move this to another file
type Order struct {
	Price    int
	Amount   int
	Stock    string // should probably rename this to something more general
	Type     string // buy/sell
	IsMarket bool   // if market or limit order
}

type OrderBook struct {
	Asks []Order
	Bids []Order
}

func (ob *OrderBook) SubmitOrder(o Order) error {

	t := o.Type

	if t == "" {
		return errors.New("Order type must be defined")
	}

	// add to bids
	if t == "buy" {
		ob.Bids = append(ob.Bids, o)
	}

	if t == "sell" {
		ob.Asks = append(ob.Asks, o)
	}

	return nil
}

func (*OrderBook) CalculateSpread() int {
	return 0
}

func main() {
	o := Order{Type: "buy"}
	book := OrderBook{}
	book.SubmitOrder(o)
	fmt.Println(book)
}
