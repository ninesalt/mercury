package main

import (
	"errors"
	"time"
)

// Order is the basic struct that wraps
// all order information
type Order struct {
	Price       float32
	Amount      float32
	Type        string // buy/sell
	PartialFill bool
	IsMarket    bool // if market or limit order
	Created     time.Time
}

// CreateOrder wraps the Order struct to create instances of Orders
func CreateOrder(price, amount float32, ordertype string, pfill bool, IsMarket bool) (Order, error) {

	if price <= 0 {
		return Order{}, errors.New("Price must be a positive integer greater than 0")
	}
	if amount <= 0 {
		return Order{}, errors.New("Amount must be a positive integer greater than 0")
	}
	timeNow := time.Now().UTC()
	return Order{price, amount, ordertype, IsMarket, pfill, timeNow}, nil
}
