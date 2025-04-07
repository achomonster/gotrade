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

### POST `/order`

Send a new order.

```json
{
  "id": "123",
  "traderId": "alice",
  "type": "buy",
  "price": 101,
  "quantity": 2
}