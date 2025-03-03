
            Start editing…# Expense Tracker API
A simple expense tracker API built with Golang &amp; Gin, allowing users to manage their expenses with CRUD operations and JWT authentication.

## 🚀 Features
✅ User authentication (JWT-based)
✅ Create, Read, Update, Delete (CRUD) expenses
✅ Secure API endpoints
✅ MySQL database integration

## 🛠 Tech Stack

Go (Golang)
Gin Framework
MySQL
JWT Authentication
Docker (Optional for containerization)


## 🔧 Installation &amp; Setup
### 1️⃣ Clone the Repository
git clone https://github.com/your-username/expense-tracker.git
cd expense-tracker

### 2️⃣ Install Dependencies
Ensure you have Go 1.19+ installed. Then, run:
go mod tidy

### 3️⃣ Configure Environment Variables
Create a .env file in the project root and add:
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=expense_tracker
JWT_SECRET=your_jwt_secret_key

### 4️⃣ Set Up MySQL Database
Run the following SQL to create the database and expenses table:
CREATE DATABASE expense_tracker;

USE expense_tracker;

CREATE TABLE expenses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    amount INT NOT NULL,
    category VARCHAR(255) NOT NULL,
    description TEXT,
    date DATE NOT NULL,
    userId INT NOT NULL
);

### 5️⃣ Run the Server
go run main.go

The server will start at [http://localhost:8080](http://localhost:8080/).

## 🔑 Authentication
The API uses JWT-based authentication. You need to include the token in the Authorization header for protected routes.

## 📌 API Endpoints
### 🔹 User Authentication



Method
Endpoint
Description




POST
/signup
Register a new user


POST
/login
Authenticate and get a JWT token



### 🔹 Expense Management



Method
Endpoint
Description




POST
/expense
Create a new expense


GET
/expenses
Get expenses (filtered by date range)


PUT
/expenses/:id
Update an existing expense


DELETE
/expenses/:id
Delete an expense



#### Filters for GET /expenses
You can filter expenses by date range:

Last 7 days: /expenses?range=week
Last 30 days: /expenses?range=month
Last 3 months: /expenses?range=3months
Custom range: /expenses?start=YYYY-MM-DD&amp;end=YYYY-MM-DD


## 🧪 Testing the API
You can test the API using Postman or the provided expense.http file:

Create a User &amp; Get Token

POST http://localhost:8080/signup
Content-Type: application/json

{
  "email": "testuser@gmail.com",
  "password": "password123"
}


Login

POST http://localhost:8080/login
Content-Type: application/json

{
  "email": "testuser@gmail.com",
  "password": "password123"
}

Response: { "token": "your_jwt_token" }

Create Expense (Use token from login response)

POST http://localhost:8080/expense
Authorization: Bearer your_jwt_token
Content-Type: application/json

{
  "amount": 1000,
  "category": "Food",
  "description": "Lunch at restaurant",
  "date": "2025-03-03"
}


Get Expenses

GET http://localhost:8080/expenses?range=week
Authorization: Bearer your_jwt_token


## 🐳 Run with Docker
If you prefer running the API using Docker, follow these steps:

Build the Docker image

docker build -t expense-tracker .


Run the container

docker run -p 8080:8080 --env-file .env expense-tracker


## 🤝 Contributing
Feel free to submit issues or open pull requests for improvements!

## 📜 License
This project is licensed under the MIT License.