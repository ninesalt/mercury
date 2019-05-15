package main

import (
	"errors"
	"fmt"
)

type OrderBook struct {
	Stock string //rename this to something more general
	Asks  []Order
	Bids  []Order
	// add volume and last price here?
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
	o, _ := CreateOrder(1, 3, "buy", false, false)
	book := OrderBook{Stock: "AAPL"}
	book.SubmitOrder(o)
	fmt.Println(book)
}
