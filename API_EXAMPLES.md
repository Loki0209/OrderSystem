# API Testing Examples - Quick Reference

## 🚀 All API Endpoints with Examples

---

## 📋 Table of Contents
1. [Authentication](#authentication)
2. [User Management](#user-management)
3. [Product Management](#product-management)

---

## Authentication

### 1. Register New User

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

**JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/v1/auth/register', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: 'John Doe',
    email: 'john@example.com',
    password: 'password123'
  })
})
.then(res => res.json())
.then(data => console.log(data));
```

**Response:**
```json
{
  "message": "User registered successfully",
  "data": {
    "id": "693d5ab5b95fa28643bad0d8",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "is_active": true
  }
}
```

---

### 2. Login

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**JavaScript (Fetch):**
```javascript
fetch('http://localhost:8080/api/v1/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    email: 'john@example.com',
    password: 'password123'
  })
})
.then(res => res.json())
.then(data => {
  console.log('Token:', data.token);
  localStorage.setItem('authToken', data.token);
});
```

**Response:**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "data": {
    "id": "693d5ab5b95fa28643bad0d8",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user"
  }
}
```

**Save this token! You'll need it for all protected endpoints.**

---

## User Management

### 3. Get All Users (Protected)

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const token = localStorage.getItem('authToken');

fetch('http://localhost:8080/api/v1/users', {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

**Response:**
```json
{
  "message": "Users retrieved successfully",
  "count": 2,
  "data": [
    {
      "id": "693d5ab5b95fa28643bad0d8",
      "name": "John Doe",
      "email": "john@example.com",
      "role": "user",
      "is_active": true
    }
  ]
}
```

---

### 4. Get User by ID (Protected)

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/users/693d5ab5b95fa28643bad0d8 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const userId = '693d5ab5b95fa28643bad0d8';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/users/${userId}`, {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 5. Update User (PUT - Protected)

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/693d5ab5b95fa28643bad0d8 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Updated"
  }'
```

**JavaScript (Fetch):**
```javascript
const userId = '693d5ab5b95fa28643bad0d8';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/users/${userId}`, {
  method: 'PUT',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    name: 'John Doe Updated'
  })
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 6. Delete User (Protected)

**cURL Example:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/693d5ab5b95fa28643bad0d8 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const userId = '693d5ab5b95fa28643bad0d8';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/users/${userId}`, {
  method: 'DELETE',
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

---

## Product Management

### 7. Create Product (POST - Protected)

**cURL Example:**
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop",
    "description": "High performance laptop",
    "price": 1200.50,
    "quantity": 10,
    "category": "Electronics",
    "sku": "LAP-001"
  }'
```

**JavaScript (Fetch):**
```javascript
const token = localStorage.getItem('authToken');

fetch('http://localhost:8080/api/v1/products', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    name: 'Laptop',
    description: 'High performance laptop',
    price: 1200.50,
    quantity: 10,
    category: 'Electronics',
    sku: 'LAP-001'
  })
})
.then(res => res.json())
.then(data => console.log(data));
```

**Response:**
```json
{
  "message": "Product created successfully",
  "data": {
    "id": "693d65cf2664008c4c7795c0",
    "name": "Laptop",
    "description": "High performance laptop",
    "price": 1200.5,
    "quantity": 10,
    "category": "Electronics",
    "sku": "LAP-001",
    "is_active": true,
    "created_at": "2025-12-13T18:40:39.963Z"
  }
}
```

---

### 8. Get All Products (GET - Protected)

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**With Filters:**
```bash
# Filter by category
curl -X GET "http://localhost:8080/api/v1/products?category=Electronics" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# Filter by active status
curl -X GET "http://localhost:8080/api/v1/products?is_active=true" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# Multiple filters
curl -X GET "http://localhost:8080/api/v1/products?category=Electronics&is_active=true" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const token = localStorage.getItem('authToken');

// Get all products
fetch('http://localhost:8080/api/v1/products', {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));

// With filters
fetch('http://localhost:8080/api/v1/products?category=Electronics&is_active=true', {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 9. Get Product by ID (GET - Protected)

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/products/693d65cf2664008c4c7795c0 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const productId = '693d65cf2664008c4c7795c0';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/products/${productId}`, {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 10. Update Product - Full Update (PUT - Protected)

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/products/693d65cf2664008c4c7795c0 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Laptop",
    "description": "High performance gaming laptop with RTX 4090",
    "price": 2500.99,
    "quantity": 5
  }'
```

**JavaScript (Fetch):**
```javascript
const productId = '693d65cf2664008c4c7795c0';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/products/${productId}`, {
  method: 'PUT',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    name: 'Gaming Laptop',
    description: 'High performance gaming laptop with RTX 4090',
    price: 2500.99,
    quantity: 5
  })
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 11. Update Product - Partial Update (PATCH - Protected)

**cURL Example:**
```bash
curl -X PATCH http://localhost:8080/api/v1/products/693d65cf2664008c4c7795c0 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "price": 2399.99
  }'
```

**JavaScript (Fetch):**
```javascript
const productId = '693d65cf2664008c4c7795c0';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/products/${productId}`, {
  method: 'PATCH',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    price: 2399.99
  })
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 12. Update Product Quantity (PUT - Protected)

**cURL Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/products/693d65cf2664008c4c7795c0/quantity \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "quantity": 100
  }'
```

**JavaScript (Fetch):**
```javascript
const productId = '693d65cf2664008c4c7795c0';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/products/${productId}/quantity`, {
  method: 'PUT',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    quantity: 100
  })
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 13. Delete Product (DELETE - Protected)

**cURL Example:**
```bash
curl -X DELETE http://localhost:8080/api/v1/products/693d65cf2664008c4c7795c0 \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const productId = '693d65cf2664008c4c7795c0';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/products/${productId}`, {
  method: 'DELETE',
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 14. Search Products (GET - Protected)

**cURL Example:**
```bash
curl -X GET "http://localhost:8080/api/v1/products/search?q=laptop" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const searchQuery = 'laptop';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/products/search?q=${searchQuery}`, {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

---

### 15. Get Products by Category (GET - Protected)

**cURL Example:**
```bash
curl -X GET http://localhost:8080/api/v1/products/category/Electronics \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

**JavaScript (Fetch):**
```javascript
const category = 'Electronics';
const token = localStorage.getItem('authToken');

fetch(`http://localhost:8080/api/v1/products/category/${category}`, {
  headers: {
    'Authorization': `Bearer ${token}`
  }
})
.then(res => res.json())
.then(data => console.log(data));
```

---

## 🔄 Complete Workflow Example

### Frontend Login Flow

```javascript
// 1. Login
async function loginWorkflow() {
  // Step 1: Login
  const loginResponse = await fetch('http://localhost:8080/api/v1/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({
      email: 'john@example.com',
      password: 'password123'
    })
  });

  const loginData = await loginResponse.json();

  // Step 2: Store token
  localStorage.setItem('authToken', loginData.token);
  localStorage.setItem('userData', JSON.stringify(loginData.data));

  console.log('Logged in successfully!');
  console.log('User:', loginData.data);
  console.log('Token:', loginData.token);

  // Step 3: Fetch user's products
  const token = loginData.token;

  const productsResponse = await fetch('http://localhost:8080/api/v1/products', {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  });

  const productsData = await productsResponse.json();
  console.log('Products:', productsData);

  return { user: loginData.data, products: productsData.data };
}

// Run the workflow
loginWorkflow()
  .then(data => console.log('Workflow complete:', data))
  .catch(error => console.error('Error:', error));
```

---

## 📱 React Native Example

```javascript
import AsyncStorage from '@react-native-async-storage/async-storage';

// Login
async function login(email, password) {
  try {
    const response = await fetch('http://localhost:8080/api/v1/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });

    const data = await response.json();

    if (response.ok) {
      await AsyncStorage.setItem('authToken', data.token);
      await AsyncStorage.setItem('userData', JSON.stringify(data.data));
      return data;
    } else {
      throw new Error(data.message);
    }
  } catch (error) {
    console.error('Login error:', error);
    throw error;
  }
}

// Get products
async function getProducts() {
  try {
    const token = await AsyncStorage.getItem('authToken');

    const response = await fetch('http://localhost:8080/api/v1/products', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });

    return await response.json();
  } catch (error) {
    console.error('Get products error:', error);
    throw error;
  }
}
```

---

## 🎯 Quick Tips

1. **Always check the response status code**
   ```javascript
   if (!response.ok) {
     throw new Error(data.message || 'Request failed');
   }
   ```

2. **Handle 401 (Unauthorized) globally**
   ```javascript
   if (response.status === 401) {
     localStorage.removeItem('authToken');
     window.location.href = '/login';
   }
   ```

3. **Use async/await for cleaner code**
   ```javascript
   async function fetchData() {
     try {
       const response = await fetch(url);
       const data = await response.json();
       return data;
     } catch (error) {
       console.error(error);
     }
   }
   ```

---

## 📞 Need Help?

- **API Base URL:** `http://localhost:8080/api/v1`
- **Health Check:** `GET http://localhost:8080/api/v1/hello`
- **Documentation:** See FRONTEND_API_GUIDE.md for detailed integration guide

---

**Last Updated:** December 13, 2025
