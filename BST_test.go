package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"./src/user"
)

// init runs before all tests
func init() {
	// this is necessary so that random numbers aren't always the same
	rand.Seed(time.Now().UnixNano())
}

func TestAdd(t *testing.T) {

	// Tree object to be tested
	tree := Tree{nil}

	// Using slice as control
	control := []Order{}
	u := user.CreateUser("john")

	// Create and insert orders
	for i := 0; i < 20; i++ {
		order, _ := CreateOrder(u, rand.Float32(), 3, "buy", false, false)
		tree.Add(order.Price, order)
		control = append(control, order)
	}

	// Get list of in-order (sorted) orders from our tree
	orders := tree.Flatten()
	got := []Order{}
	for _, order := range orders {
		// Need type assertion to change interface{} to Order
		got = append(got, order.(Order))
	}

	// Sort control slice
	sort.Slice(control, func(i, j int) bool { return control[i].Price < control[j].Price })

	want := control
	// Compare the lists of prices
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Got %v; want %v", got, want)
	}
}
