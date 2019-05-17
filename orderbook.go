package main

import (
	"errors"
	"fmt"

	"./src/user"
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

func (ob *OrderBook) Monitor() {
	// ticks := time.Tick(2 * time.Second)
	// for now := range ticks {
	// 	fmt.Println(now)
	// }
}

func (ob *OrderBook) FillOrder(bid *Order, ask *Order) {
	seller := ask.User
	stock := ob.Stock
	buyer := bid.User
	amount := ask.Amount

	//TODO: make this part atomic
	seller.SubtractFromBalance(stock, amount)
	buyer.AddToBalance(stock, amount)
	bid.AmountFilled += amount
	ask.AmountFilled += amount

	if bid.AmountFilled == bid.Amount {
		// delete bid from OB
	}

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
	a := user.CreateUser("adam")
	e := user.CreateUser("eve")

	o, _ := CreateOrder(a, 12.41, 3, "buy", false, false)
	o2, _ := CreateOrder(e, 12.41, 3, "sell", false, false)

	fmt.Println(o, o2)
	book := CreateOrderBook("AAPL")
	book.FillOrder(&o, &o2)
	fmt.Println(o, o2)
	fmt.Println(book)
	book.Monitor()
	// book.SubmitOrder(o)
	// book.SubmitOrder(o2)
	// fmt.Println(book)
}
