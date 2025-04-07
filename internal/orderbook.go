// internal/orderbook.go
package internal

import (
	"fmt"
	"sort"
	"github.com/achomonster/gotrade/models"
)

type OrderBook struct {
	BuyOrders  []*models.Order
	SellOrders []*models.Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		BuyOrders:  []*models.Order{},
		SellOrders: []*models.Order{},
	}
}

func (ob *OrderBook) AddOrder(order *models.Order) {
	switch order.Type {
	case models.Buy:
		ob.BuyOrders = append(ob.BuyOrders, order)
		// Sort buys: highest price first
		sort.SliceStable(ob.BuyOrders, func(i, j int) bool {
			return ob.BuyOrders[i].Price > ob.BuyOrders[j].Price
		})
	case models.Sell:
		ob.SellOrders = append(ob.SellOrders, order)
		// Sort sells: lowest price first
		sort.SliceStable(ob.SellOrders, func(i, j int) bool {
			return ob.SellOrders[i].Price < ob.SellOrders[j].Price
		})
	}

	fmt.Printf("Order added: %+v\n", order)
}

// 📊 PrintOrderBook shows all remaining orders in the book
func (ob *OrderBook) PrintOrderBook() {
	fmt.Println("\n=== 📘 ORDER BOOK ===")

	fmt.Println("\n🟢 BUY ORDERS:")
	if len(ob.BuyOrders) == 0 {
		fmt.Println("  (none)")
	}
	for _, o := range ob.BuyOrders {
		fmt.Printf("  ID: %s, Price: %.2f, Qty: %.2f\n", o.ID, o.Price, o.Quantity)
	}

	fmt.Println("\n🔴 SELL ORDERS:")
	if len(ob.SellOrders) == 0 {
		fmt.Println("  (none)")
	}
	for _, o := range ob.SellOrders {
		fmt.Printf("  ID: %s, Price: %.2f, Qty: %.2f\n", o.ID, o.Price, o.Quantity)
	}

	fmt.Println("======================\n")
}