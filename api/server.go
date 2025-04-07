// api/server.go
package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/achomonster/gotrade/engine"
	"github.com/achomonster/gotrade/models"
)

type Server struct {
	matcher *engine.Matcher
}

func NewServer(matcher *engine.Matcher) *Server {
	return &Server{matcher: matcher}
}

func (s *Server) Start() {
	http.HandleFunc("/order", s.handleOrder)
	http.HandleFunc("/trades", s.handleGetTrades)
	http.HandleFunc("/orderbook", s.handleGetOrderBook)

	fmt.Println("🌐 API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) handleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	s.matcher.Match(&order)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("✅ Order received and processed"))
}

func (s *Server) handleGetTrades(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	trades := s.matcher.GetTrades()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trades)
}

// ✅ NEW: /orderbook endpoint
func (s *Server) handleGetOrderBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	buys, sells := s.matcher.GetOrderBook()

	data := map[string]interface{}{
		"buyOrders":  buys,
		"sellOrders": sells,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}