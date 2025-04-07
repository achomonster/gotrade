// models/order.go
package models

type OrderType string

const (
	Buy  OrderType = "buy"
	Sell OrderType = "sell"
)

type Order struct {
	ID       string
	TraderID string
	Type     OrderType
	Price    float64
	Quantity float64
}

// ➕ New struct to store trades
type Trade struct {
	BuyOrderID  string
	SellOrderID string
	Price       float64
	Quantity    float64
	Timestamp   string
}