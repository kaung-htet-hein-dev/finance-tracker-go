#!/bin/bash

# API Test Script for Finance Tracker
set -e

API_URL="http://localhost:8080"
TOKEN=""

echo "🚀 Starting Finance Tracker API Tests..."

# Clean up any existing database and running processes
rm -f finance_tracker.db
pkill -f "go run cmd/server/main.go" 2>/dev/null || true
sleep 1

# Start server in background
echo "Starting server..."
go run cmd/server/main.go &
SERVER_PID=$!

# Wait for server to start
sleep 3

# Function to cleanup on exit
cleanup() {
    echo "Cleaning up..."
    kill $SERVER_PID 2>/dev/null || true
    rm -f finance_tracker.db
}
trap cleanup EXIT

# Test 1: Health Check
echo "✅ Testing Health Check..."
response=$(curl -s -w "\n%{http_code}" "$API_URL/health")
status_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n -1)

if [ "$status_code" -ne 200 ]; then
    echo "❌ Health check failed. Status: $status_code"
    exit 1
fi

echo "✅ Health check passed: $body"

# Test 2: User Registration
echo "✅ Testing User Registration..."
response=$(curl -s -w "\n%{http_code}" -X POST "$API_URL/api/v1/auth/register" \
    -H "Content-Type: application/json" \
    -d '{
        "username": "testuser",
        "email": "test@example.com",
        "password": "password123"
    }')

status_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n -1)

if [ "$status_code" -ne 201 ]; then
    echo "❌ User registration failed. Status: $status_code, Body: $body"
    exit 1
fi

# Extract token
TOKEN=$(echo "$body" | python3 -c "import sys, json; print(json.load(sys.stdin)['token'])")
echo "✅ User registration passed. Token obtained."

# Test 3: User Login
echo "✅ Testing User Login..."
response=$(curl -s -w "\n%{http_code}" -X POST "$API_URL/api/v1/auth/login" \
    -H "Content-Type: application/json" \
    -d '{
        "username": "testuser",
        "password": "password123"
    }')

status_code=$(echo "$response" | tail -n1)
if [ "$status_code" -ne 200 ]; then
    echo "❌ User login failed. Status: $status_code"
    exit 1
fi

echo "✅ User login passed."

# Test 4: Create Income Transaction
echo "✅ Testing Create Income Transaction..."
response=$(curl -s -w "\n%{http_code}" -X POST "$API_URL/api/v1/transactions" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
        "type": "income",
        "amount": 3000.00,
        "category": "salary",
        "description": "Monthly salary",
        "date": "2024-01-01T09:00:00Z"
    }')

status_code=$(echo "$response" | tail -n1)
if [ "$status_code" -ne 201 ]; then
    echo "❌ Create income transaction failed. Status: $status_code"
    exit 1
fi

echo "✅ Create income transaction passed."

# Test 5: Create Expense Transaction
echo "✅ Testing Create Expense Transaction..."
response=$(curl -s -w "\n%{http_code}" -X POST "$API_URL/api/v1/transactions" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
        "type": "expense",
        "amount": 150.75,
        "category": "groceries",
        "description": "Weekly grocery shopping",
        "date": "2024-01-02T10:30:00Z"
    }')

status_code=$(echo "$response" | tail -n1)
if [ "$status_code" -ne 201 ]; then
    echo "❌ Create expense transaction failed. Status: $status_code"
    exit 1
fi

echo "✅ Create expense transaction passed."

# Test 6: Get Transactions
echo "✅ Testing Get Transactions..."
response=$(curl -s -w "\n%{http_code}" -X GET "$API_URL/api/v1/transactions" \
    -H "Authorization: Bearer $TOKEN")

status_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n -1)

if [ "$status_code" -ne 200 ]; then
    echo "❌ Get transactions failed. Status: $status_code"
    exit 1
fi

# Check if we have transactions
transaction_count=$(echo "$body" | python3 -c "import sys, json; print(len(json.load(sys.stdin)['transactions']))")
if [ "$transaction_count" -lt 2 ]; then
    echo "❌ Expected at least 2 transactions, got $transaction_count"
    exit 1
fi

echo "✅ Get transactions passed. Found $transaction_count transactions."

# Test 7: Get Insights
echo "✅ Testing Get Insights..."
response=$(curl -s -w "\n%{http_code}" -X GET "$API_URL/api/v1/insights" \
    -H "Authorization: Bearer $TOKEN")

status_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n -1)

if [ "$status_code" -ne 200 ]; then
    echo "❌ Get insights failed. Status: $status_code"
    exit 1
fi

# Check insights data
total_income=$(echo "$body" | python3 -c "import sys, json; print(json.load(sys.stdin)['total_income'])")
total_expenses=$(echo "$body" | python3 -c "import sys, json; print(json.load(sys.stdin)['total_expenses'])")

if (( $(echo "$total_income < 3000" | bc -l) )); then
    echo "❌ Expected total income >= 3000, got $total_income"
    exit 1
fi

if (( $(echo "$total_expenses < 150" | bc -l) )); then
    echo "❌ Expected total expenses >= 150, got $total_expenses"
    exit 1
fi

echo "✅ Get insights passed. Income: $total_income, Expenses: $total_expenses"

# Test 8: Authentication Required
echo "✅ Testing Authentication Required..."
response=$(curl -s -w "\n%{http_code}" -X GET "$API_URL/api/v1/transactions")

status_code=$(echo "$response" | tail -n1)
if [ "$status_code" -ne 401 ]; then
    echo "❌ Expected 401 for unauthenticated request, got $status_code"
    exit 1
fi

echo "✅ Authentication requirement test passed."

echo "🎉 All API tests passed successfully!"