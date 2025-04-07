# GoTrade - Matching Engine & Trade Simulator 💸⚙️

GoTrade is a real-time order matching engine written in Go. It simulates a live exchange backend where buy and sell orders are processed, matched, and tracked using clean matching logic, APIs, and in-memory storage.

Built for learning, performance, and demonstration — perfect for fintech, backend, and trading-focused portfolios.

---

## 🚀 Features

- ✅ Limit order matching (buy/sell)
- ✅ Partial fills support
- ✅ Price-time priority logic
- ✅ Real-time `/order` API (POST)
- ✅ Live `/trades` API (GET)
- ✅ Automated market simulator
- ✅ In-memory trade ledger with timestamps
- ✅ Built in Go with zero dependencies

---

## 🧠 Why This Matters

This project shows off:

- Backend design & clean Go architecture
- Understanding of real-world matching logic
- API development with `net/http`
- Simulated concurrency and state management
- Real-time system thinking

---

## 📡 API Endpoints

### POST /order
Submit a new order to the exchange.

Request body (JSON):
{
  "id": "123",
  "traderId": "alice",
  "type": "buy",
  "price": 101,
  "quantity": 2
}

### GET /trades
Returns a list of all executed trades.

Response (JSON):
[
  {
    "BuyOrderID": "123",
    "SellOrderID": "456",
    "Price": 105,
    "Quantity": 1,
    "Timestamp": "10:32:12"
  }
]

### GET /orderbook
Returns the current state of the order book.

Response (JSON):
{
  "buyOrders": [
    {
      "ID": "1001",
      "TraderID": "bob",
      "Type": "buy",
      "Price": 105,
      "Quantity": 3
    }
  ],
  "sellOrders": [
    {
      "ID": "1002",
      "TraderID": "carol",
      "Type": "sell",
      "Price": 108,
      "Quantity": 2
    }
  ]
}

## 📄 License

Licensed under the MIT License.