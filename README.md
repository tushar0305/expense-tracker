# Expense Tracker API

## About the Project
This is an expense tracker API built using Golang and Gin. The API allows users to create, retrieve, update, and delete expenses while ensuring authentication using JWT.

## Project URL
[Expense Tracker Repository](https://github.com/tushar0305/expense-tracker.git)

## Prerequisites
- Go 1.18+
- SQLite (or another configured database)
- VS Code or any code editor
- REST Client extension in VS Code (for testing API requests)

## Project Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/tushar0305/expense-tracker.git
   cd expense-tracker
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints

### 1. User Authentication
#### Sign Up
```http
POST http://localhost:8080/signup
Content-Type: application/json

{
  "email": "testuser@gmail.com",
  "password": "password123"
}
```

#### Login
```http
POST http://localhost:8080/login
Content-Type: application/json

{
  "email": "testuser@gmail.com",
  "password": "password123"
}
```
**Response:**
```json
{
  "token": "your_jwt_token"
}
```

### 2. Expense Management (Requires Authorization Header)
#### Create Expense
```http
POST http://localhost:8080/expense
Authorization: your_jwt_token
Content-Type: application/json

{
  "amount": 500,
  "category": "Food",
  "description": "Lunch at restaurant",
  "date": "2025-03-02T00:00:00Z"
}
```

#### Get Expenses
```http
GET http://localhost:8080/expenses?range=week
Authorization: your_jwt_token
```

#### Update Expense
```http
PUT http://localhost:8080/expenses/1
Authorization: your_jwt_token
Content-Type: application/json

{
  "amount": 700,
  "category": "Food",
  "description": "Dinner at restaurant",
  "date": "2025-03-02T00:00:00Z"
}
```

#### Delete Expense
```http
DELETE http://localhost:8080/expenses/1
Authorization: your_jwt_token
```

## Running and Testing in VS Code
- The project includes a folder with `.http` files for testing API requests.
- Open the folder after running `go run main.go` and use the **REST Client** extension in VS Code to send requests.

