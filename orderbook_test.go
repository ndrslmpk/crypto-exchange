package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 10)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderB)

	fmt.Println(l)
}

func TestOrderbook(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 10)
	buyOrderB := NewOrder(true, 2000)
	buyOrderC := NewOrder(true, 100)

	ob.PlaceOrder(100_000, buyOrderA)
	ob.PlaceOrder(110_000, buyOrderB)

	ob.PlaceOrder(100_000, buyOrderC)

	fmt.Printf("%+v", ob.Bids[0])
	// fmt.Printf("%+v", ob.Bids[2])

	// for i := 0; i < len(ob.Bids); i++ {
	// 	fmt.Printf("%+v\n", ob.Bids[i])
	// }
}
