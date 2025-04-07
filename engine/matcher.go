// engine/matcher.go
package engine

import (
	"fmt"
	"time"

	"github.com/achomonster/gotrade/internal"
	"github.com/achomonster/gotrade/models"
)

type Matcher struct {
	orderBook *internal.OrderBook
	Trades    []models.Trade
}

func NewMatcher() *Matcher {
	return &Matcher{
		orderBook: internal.NewOrderBook(),
		Trades:    []models.Trade{},
	}
}

func (m *Matcher) Match(order *models.Order) {
	switch order.Type {
	case models.Buy:
		i := 0
		for i < len(m.orderBook.SellOrders) && order.Quantity > 0 {
			sell := m.orderBook.SellOrders[i]
			if order.Price >= sell.Price {
				tradeQty := min(order.Quantity, sell.Quantity)
				fmt.Printf("✅ TRADE! %s buys from %s at %.2f (Qty: %.2f)\n",
					order.ID, sell.ID, sell.Price, tradeQty,
				)

				m.recordTrade(order.ID, sell.ID, sell.Price, tradeQty)

				order.Quantity -= tradeQty
				sell.Quantity -= tradeQty

				if sell.Quantity == 0 {
					m.orderBook.SellOrders = append(m.orderBook.SellOrders[:i], m.orderBook.SellOrders[i+1:]...)
				} else {
					i++
				}
			} else {
				break
			}
		}

	case models.Sell:
		i := 0
		for i < len(m.orderBook.BuyOrders) && order.Quantity > 0 {
			buy := m.orderBook.BuyOrders[i]
			if order.Price <= buy.Price {
				tradeQty := min(order.Quantity, buy.Quantity)
				fmt.Printf("✅ TRADE! %s sells to %s at %.2f (Qty: %.2f)\n",
					order.ID, buy.ID, buy.Price, tradeQty,
				)

				m.recordTrade(buy.ID, order.ID, buy.Price, tradeQty)

				order.Quantity -= tradeQty
				buy.Quantity -= tradeQty

				if buy.Quantity == 0 {
					m.orderBook.BuyOrders = append(m.orderBook.BuyOrders[:i], m.orderBook.BuyOrders[i+1:]...)
				} else {
					i++
				}
			} else {
				break
			}
		}
	}

	if order.Quantity > 0 {
		m.orderBook.AddOrder(order)
	}
}

func (m *Matcher) recordTrade(buyID, sellID string, price, qty float64) {
	trade := models.Trade{
		BuyOrderID:  buyID,
		SellOrderID: sellID,
		Price:       price,
		Quantity:    qty,
		Timestamp:   time.Now().Format("15:04:05"),
	}
	m.Trades = append(m.Trades, trade)

	fmt.Printf("🧾 TRADE LOG: %+v\n", trade)
}

func (m *Matcher) PrintOrderBook() {
	m.orderBook.PrintOrderBook()
}

func (m *Matcher) GetTrades() []models.Trade {
	return m.Trades
}

// ✅ NEW: Return current buy/sell order slices
func (m *Matcher) GetOrderBook() (buyOrders, sellOrders []*models.Order) {
	return m.orderBook.BuyOrders, m.orderBook.SellOrders
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}