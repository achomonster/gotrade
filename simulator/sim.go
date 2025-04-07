// simulator/sim.go
package simulator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/achomonster/gotrade/models"
)

var traderIDs = []string{"alice", "bob", "carol", "dave", "emma"}

func StartSimulation() {
	go func() {
		for {
			order := generateRandomOrder()
			sendOrder(order)

			time.Sleep(time.Second * 2) // adjust frequency here
		}
	}()
}

func generateRandomOrder() *models.Order {
	orderType := models.Buy
	if rand.Intn(2) == 0 {
		orderType = models.Sell
	}

	price := float64(100 + rand.Intn(10)) // random price 100–109
	quantity := float64(rand.Intn(5)+1)   // random qty 1–5

	order := &models.Order{
		ID:       fmt.Sprintf("%d", rand.Intn(100000)),
		TraderID: traderIDs[rand.Intn(len(traderIDs))],
		Type:     orderType,
		Price:    price,
		Quantity: quantity,
	}

	return order
}

func sendOrder(order *models.Order) {
	body, _ := json.Marshal(order)
	resp, err := http.Post("http://localhost:8080/order", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("❌ Error sending order:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("📤 Sent order: %+v\n", order)
}