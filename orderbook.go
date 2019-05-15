package main

import (
	"errors"
	"fmt"
)

type OrderBook struct {
	Stock       string //rename this to something more general
	Asks        []Order
	Bids        []Order
	AsksMapping map[float32][]*Order
	BidsMapping map[float32][]*Order
	// add volume and last price here?
}

func (ob *OrderBook) SubmitOrder(o Order) error {

	t := o.Type

	if t == "" {
		return errors.New("Order type must be defined")
	}

	// if o.IsMarket {
	// 	if t == "buy" {
	// 		ob.FillOrder(o, ob.Asks[0])
	// 	}
	// 	ob.FillOrder(ob.Bids[0], o)
	// }

	// add to bids
	if t == "buy" {
		ob.Bids = append(ob.Bids, o)
		ob.BidsMapping[o.Price] = append(ob.BidsMapping[o.Price], &o)
	}

	if t == "sell" {
		ob.Asks = append(ob.Asks, o)
		ob.AsksMapping[o.Price] = append(ob.AsksMapping[o.Price], &o)
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
	return OrderBook{
		Stock:       stock,
		BidsMapping: make(map[float32][]*Order),
		AsksMapping: make(map[float32][]*Order)}
}

func main() {
	o, _ := CreateOrder(12.41, 3, "buy", false, false)
	o2, _ := CreateOrder(12.41, 3, "buy", false, false)
	book := CreateOrderBook("AAPL")
	book.SubmitOrder(o)
	book.SubmitOrder(o2)
	fmt.Println(book)
}
