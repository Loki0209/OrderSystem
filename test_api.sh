#!/bin/bash

# API Test Script for Order Management System
BASE_URL="http://localhost:8080/api/v1"

echo "================================"
echo "Testing Order Management API"
echo "================================"
echo ""

# Test 1: Hello World Endpoint
echo "1. Testing Hello World Endpoint..."
curl -X GET "$BASE_URL/hello"
echo -e "\n"

# Test 2: Register a new user
echo "2. Testing User Registration..."
REGISTER_RESPONSE=$(curl -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }')
echo $REGISTER_RESPONSE
echo -e "\n"

# Test 3: Login
echo "3. Testing User Login..."
LOGIN_RESPONSE=$(curl -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }')
echo $LOGIN_RESPONSE

# Extract token from login response
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)
echo "Token: $TOKEN"
echo -e "\n"

# Test 4: Get All Users (Protected Route)
echo "4. Testing Get All Users (Protected)..."
curl -X GET "$BASE_URL/users" \
  -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# Test 5: Get User by ID (Extract ID from login response)
USER_ID=$(echo $LOGIN_RESPONSE | grep -o '"id":"[^"]*' | cut -d'"' -f4)
echo "5. Testing Get User by ID..."
echo "User ID: $USER_ID"
curl -X GET "$BASE_URL/users/$USER_ID" \
  -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# Test 6: Update User
echo "6. Testing Update User..."
curl -X PUT "$BASE_URL/users/$USER_ID" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Updated"
  }'
echo -e "\n"

# Test 7: Create another user for testing
echo "7. Creating another test user..."
curl -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com",
    "password": "password456"
  }'
echo -e "\n"

# Test 8: Get All Users Again
echo "8. Testing Get All Users Again..."
curl -X GET "$BASE_URL/users" \
  -H "Authorization: Bearer $TOKEN"
echo -e "\n"

# Test 9: Delete User (Optional - uncomment to test)
# echo "9. Testing Delete User..."
# curl -X DELETE "$BASE_URL/users/$USER_ID" \
#   -H "Authorization: Bearer $TOKEN"
# echo -e "\n"

echo "================================"
echo "API Testing Complete!"
echo "================================"
