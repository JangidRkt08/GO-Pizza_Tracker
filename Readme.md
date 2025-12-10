# 🍕 Go Pizza Tracker – Real-Time Order Tracking System

A full-stack **Pizza Order Tracking System** built with **Go (Gin), GORM, SQLite, Server-Sent Events (SSE)** and **TailwindCSS**.  
This project allows customers to place pizza orders and track them **live**, while admins can manage and update order statuses in **real-time**.

---

## Screens


## 🚀 Features

### ✅ Customer Side
- Place pizza orders with:
  - Multiple pizzas per order
  - Size, type & special instructions
- Real-time **order tracking**
- Live **status progress bar**
- Automatic updates using **Server-Sent Events (SSE)**
- Clean, modern UI with TailwindCSS

### ✅ Admin Dashboard
- Secure **Admin Login**
- View all active orders in a table
- Update order status:
  - Order Placed → Preparing → Baking → Quality Check → Ready
- Delete orders
- **Live updates** when new orders are created (no refresh required)
- Session-based authentication

### ✅ Backend
- RESTful architecture using **Gin**
- **SQLite + GORM** for database
- **SSE-based real-time notifications**
- Secure cookie-based sessions
- Proper validation & error handling
- Clean MVC-style structure

---

## 🛠 Tech Stack

| Layer        | Technology |
|-------------|------------|
| Backend     | Go (Gin Framework) |
| Database    | SQLite + GORM |
| Frontend    | HTML Templates + TailwindCSS |
| Realtime    | Server-Sent Events (SSE) |
| Auth        | Gin Sessions |
| Config      | Environment Variables (.env) |

---

## 📁 Project Structure

```
GO-Pizza_Tracker/
├── cmd/
│   ├── main.go
│   ├── routes.go
│   ├── handlers.go
│   ├── admin.go
│   ├── customer.go
│   ├── events.go
│   └── notification.go
├── internal/
│   └── models/
├── templates/
│   ├── admin.tmpl
│   ├── customer.tmpl
│   ├── order.tmpl
│   ├── login.tmpl
│   └── static/
├── db/
│   └── orders.db
├── .env
├── go.mod
└── go.sum
```


---

##  Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/JangidRkt08/GO-Pizza_Tracker.git
cd GO-Pizza_Tracker
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Create .env File

Create a .env file in the root:

```bash
PORT=8080
DB_PATH=./db/orders.db
```

### 4. Run the Application.

```bash
go run ./cmd/
```
Server will start at:
http://localhost:8080

---

## Admin Login


Create the admin user directly in the database or via seed logic.

After login:

http://localhost:8080/admin

---

## Testing the Project
View Database:
```bash
sqlite3 ./db/orders.db
```

Inside sqlite:
```sql
.tables
SELECT * FROM orders;
SELECT * FROM order_items;
.quit
```

## 🔥 Real-Time System (SSE)

This project uses Server-Sent Events for:

- New order notifications to admin

- Live status updates to customer

### SSE Endpoints:

- /admin/notifications

- /notifications?orderId=XYZ

This avoids WebSocket complexity while keeping everything real-time.

