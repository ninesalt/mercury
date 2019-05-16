package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestAdd(t *testing.T) {

	// Tree object to be tested
	tree := Tree{nil}

	// Using slice as control
	control := []float32{}

	// Create and insert orders
	for i := 0; i < 20; i++ {
		o, _ := CreateOrder(rand.Float32(), 3, "buy", false, false)
		tree.Add(o)
		control = append(control, o.Price)
	}
	// Get list of sorted order prices from our tree
	tmp := tree.Flatten()
	orders := []float32{}
	for i := 0; i < len(tmp); i++ {
		orders = append(orders, tmp[i].Price)
	}
	// Sort control slice
	sort.Slice(control, func(i, j int) bool { return control[i] < control[j] })

	got := orders
	want := control
	// Compare the lists of prices
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Got %v; want %v", got, want)
	}
}
