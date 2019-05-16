package main

import (
	"errors"
)

type OrderBook struct {
	Stock   string //rename this to something more general
	Asks    []Order
	Bids    []Order
	BidTree Tree
	AskTree Tree
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
		ob.BidTree.Add(o.Price, o)
	}

	if t == "sell" {
		ob.Asks = append(ob.Asks, o)
		ob.AskTree.Add(o.Price, o)
	}

	return nil
}

func (ob *OrderBook) FillOrder(bid Order, ask Order) {
	//TODO
}

func (*OrderBook) CalculateSpread() int {
	return 0
}

//CreateOrderBook creates the order book instances
// and creates the mappings required
func CreateOrderBook(stock string) OrderBook {
	return OrderBook{Stock: stock}
}

func main() {
	o, _ := CreateOrder(12.41, 3, "buy", false, false)
	o2, _ := CreateOrder(12.41, 3, "sell", false, false)
	book := CreateOrderBook("AAPL")
	book.SubmitOrder(o)
	book.SubmitOrder(o2)
}
