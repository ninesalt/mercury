package main

// TODO: move this to another file
type Order struct {
	Price  int
	Amount int
	Stock  string // should probably rename this to something more general
	Type   string
}

type OrderBook struct {
	Asks []Order
	Bids []Order
}

func (*OrderBook) CalculateSpread() int {
	return 0
}

func main() {

}
