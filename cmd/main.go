// cmd/main.go
package main

import (
	"fmt"
	"github.com/achomonster/gotrade/api"
	"github.com/achomonster/gotrade/engine"
	"github.com/achomonster/gotrade/simulator"
)

func main() {
	fmt.Println("Starting GoTrade exchange... 🚀")

	matcher := engine.NewMatcher()
	server := api.NewServer(matcher)

	// 🧪 Start market simulation
	simulator.StartSimulation()

	// 🌐 Start API server
	server.Start()
}