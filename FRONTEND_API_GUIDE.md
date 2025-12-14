# Frontend API Integration Guide

## 📌 Base URL
```
http://localhost:8080/api/v1
```

---

## 🔐 Authentication Flow

### Overview
This API uses **JWT (JSON Web Token)** based authentication. The flow is:
1. User registers or logs in
2. Server returns a JWT token
3. Client stores the token (localStorage/sessionStorage)
4. Client sends token in Authorization header for protected routes

---

## 1️⃣ User Registration

### Endpoint
```
POST /api/v1/auth/register
```

### Request Headers
```
Content-Type: application/json
```

### Request Body
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### Validation Rules
- `name`: Required, string
- `email`: Required, valid email format
- `password`: Required, minimum 6 characters

### Success Response (201 Created)
```json
{
  "message": "User registered successfully",
  "data": {
    "id": "693d5ab5b95fa28643bad0d8",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-13T12:23:17.493Z",
    "updated_at": "2025-12-13T12:23:17.493Z"
  }
}
```

### Error Responses

**400 Bad Request - Validation Error**
```json
{
  "error": "Invalid request",
  "message": "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```

**400 Bad Request - Email Already Exists**
```json
{
  "error": "Registration failed",
  "message": "user with this email already exists"
}
```

### Frontend Implementation Example (JavaScript/React)

```javascript
async function registerUser(name, email, password) {
  try {
    const response = await fetch('http://localhost:8080/api/v1/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, email, password }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Registration failed');
    }

    console.log('Registration successful:', data);
    return data;
  } catch (error) {
    console.error('Registration error:', error.message);
    throw error;
  }
}

// Usage
registerUser('John Doe', 'john@example.com', 'password123')
  .then(data => {
    // Show success message
    alert('Registration successful! Please login.');
  })
  .catch(error => {
    // Show error message
    alert(error.message);
  });
```

---

## 2️⃣ User Login

### Endpoint
```
POST /api/v1/auth/login
```

### Request Headers
```
Content-Type: application/json
```

### Request Body
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

### Validation Rules
- `email`: Required, valid email format
- `password`: Required

### Success Response (200 OK)
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNjkzZDVhYjViOTVmYTI4NjQzYmFkMGQ4IiwiZW1haWwiOiJqb2huQGV4YW1wbGUuY29tIiwicm9sZSI6InVzZXIiLCJleHAiOjE3NjU3MTUwMjMsIm5iZiI6MTc2NTYyODYyMywiaWF0IjoxNzY1NjI4NjIzfQ.5ujBKlOR97CX2rKhxVAW0ZXHfu008gw46hhIfM7pDgw",
  "data": {
    "id": "693d5ab5b95fa28643bad0d8",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-13T12:23:17.493Z",
    "updated_at": "2025-12-13T12:23:17.493Z"
  }
}
```

### Error Responses

**401 Unauthorized - Invalid Credentials**
```json
{
  "error": "Login failed",
  "message": "invalid email or password"
}
```

**401 Unauthorized - Inactive Account**
```json
{
  "error": "Login failed",
  "message": "user account is inactive"
}
```

### Frontend Implementation Example (JavaScript/React)

```javascript
async function loginUser(email, password) {
  try {
    const response = await fetch('http://localhost:8080/api/v1/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Login failed');
    }

    // Store token in localStorage
    localStorage.setItem('authToken', data.token);

    // Store user data
    localStorage.setItem('userData', JSON.stringify(data.data));

    console.log('Login successful:', data);
    return data;
  } catch (error) {
    console.error('Login error:', error.message);
    throw error;
  }
}

// Usage
loginUser('john@example.com', 'password123')
  .then(data => {
    // Redirect to dashboard
    window.location.href = '/dashboard';
  })
  .catch(error => {
    // Show error message
    alert(error.message);
  });
```

---

## 3️⃣ Using Protected Routes

All endpoints except `/auth/register` and `/auth/login` require authentication.

### Authorization Header Format
```
Authorization: Bearer <your-jwt-token>
```

### Frontend Implementation Example

```javascript
// Get token from localStorage
function getAuthToken() {
  return localStorage.getItem('authToken');
}

// Make authenticated request
async function fetchProtectedData(endpoint) {
  const token = getAuthToken();

  if (!token) {
    throw new Error('No authentication token found');
  }

  try {
    const response = await fetch(`http://localhost:8080/api/v1${endpoint}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    if (response.status === 401) {
      // Token expired or invalid
      localStorage.removeItem('authToken');
      localStorage.removeItem('userData');
      window.location.href = '/login';
      throw new Error('Session expired. Please login again.');
    }

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Request failed');
    }

    return data;
  } catch (error) {
    console.error('API error:', error.message);
    throw error;
  }
}

// Usage
fetchProtectedData('/users')
  .then(data => {
    console.log('Users:', data);
  })
  .catch(error => {
    alert(error.message);
  });
```

---

## 4️⃣ Example: Get All Users (Protected)

### Endpoint
```
GET /api/v1/users
```

### Request Headers
```
Authorization: Bearer <your-jwt-token>
```

### Success Response (200 OK)
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
      "is_active": true,
      "created_at": "2025-12-13T12:23:17.493Z",
      "updated_at": "2025-12-13T12:26:30.6Z"
    }
  ]
}
```

### Error Response (401 Unauthorized)
```json
{
  "error": "Unauthorized",
  "message": "Authorization header is required"
}
```

---

## 5️⃣ Complete React Authentication Example

### Auth Context (authContext.js)

```javascript
import React, { createContext, useState, useContext, useEffect } from 'react';

const AuthContext = createContext();

export function useAuth() {
  return useContext(AuthContext);
}

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);
  const [token, setToken] = useState(null);
  const [loading, setLoading] = useState(true);

  // Check if user is logged in on mount
  useEffect(() => {
    const storedToken = localStorage.getItem('authToken');
    const storedUser = localStorage.getItem('userData');

    if (storedToken && storedUser) {
      setToken(storedToken);
      setUser(JSON.parse(storedUser));
    }
    setLoading(false);
  }, []);

  // Login function
  const login = async (email, password) => {
    try {
      const response = await fetch('http://localhost:8080/api/v1/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.message || 'Login failed');
      }

      // Save to state and localStorage
      setToken(data.token);
      setUser(data.data);
      localStorage.setItem('authToken', data.token);
      localStorage.setItem('userData', JSON.stringify(data.data));

      return data;
    } catch (error) {
      throw error;
    }
  };

  // Register function
  const register = async (name, email, password) => {
    try {
      const response = await fetch('http://localhost:8080/api/v1/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, email, password }),
      });

      const data = await response.json();

      if (!response.ok) {
        throw new Error(data.message || 'Registration failed');
      }

      return data;
    } catch (error) {
      throw error;
    }
  };

  // Logout function
  const logout = () => {
    setToken(null);
    setUser(null);
    localStorage.removeItem('authToken');
    localStorage.removeItem('userData');
  };

  // API call helper with auth
  const apiCall = async (endpoint, options = {}) => {
    const defaultHeaders = {
      'Content-Type': 'application/json',
    };

    if (token) {
      defaultHeaders['Authorization'] = `Bearer ${token}`;
    }

    const response = await fetch(`http://localhost:8080/api/v1${endpoint}`, {
      ...options,
      headers: {
        ...defaultHeaders,
        ...options.headers,
      },
    });

    if (response.status === 401) {
      logout();
      throw new Error('Session expired. Please login again.');
    }

    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.message || 'Request failed');
    }

    return data;
  };

  const value = {
    user,
    token,
    login,
    register,
    logout,
    apiCall,
    isAuthenticated: !!token,
  };

  return (
    <AuthContext.Provider value={value}>
      {!loading && children}
    </AuthContext.Provider>
  );
}
```

### Login Component (Login.js)

```javascript
import React, { useState } from 'react';
import { useAuth } from './authContext';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      await login(email, password);
      navigate('/dashboard');
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-container">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Email:</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        {error && <div className="error">{error}</div>}
        <button type="submit" disabled={loading}>
          {loading ? 'Logging in...' : 'Login'}
        </button>
      </form>
    </div>
  );
}

export default Login;
```

### Register Component (Register.js)

```javascript
import React, { useState } from 'react';
import { useAuth } from './authContext';
import { useNavigate } from 'react-router-dom';

function Register() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const { register } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

    try {
      await register(name, email, password);
      alert('Registration successful! Please login.');
      navigate('/login');
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="register-container">
      <h2>Register</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name:</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Email:</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            minLength="6"
          />
        </div>
        {error && <div className="error">{error}</div>}
        <button type="submit" disabled={loading}>
          {loading ? 'Registering...' : 'Register'}
        </button>
      </form>
    </div>
  );
}

export default Register;
```

### Protected Route Component (ProtectedRoute.js)

```javascript
import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from './authContext';

function ProtectedRoute({ children }) {
  const { isAuthenticated } = useAuth();

  if (!isAuthenticated) {
    return <Navigate to="/login" />;
  }

  return children;
}

export default ProtectedRoute;
```

### App.js Setup

```javascript
import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './authContext';
import Login from './Login';
import Register from './Register';
import Dashboard from './Dashboard';
import ProtectedRoute from './ProtectedRoute';

function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route
            path="/dashboard"
            element={
              <ProtectedRoute>
                <Dashboard />
              </ProtectedRoute>
            }
          />
        </Routes>
      </BrowserRouter>
    </AuthProvider>
  );
}

export default App;
```

---

## 6️⃣ Product API Endpoints (With Auth)

### Create Product
```javascript
const { apiCall } = useAuth();

const createProduct = async (productData) => {
  return await apiCall('/products', {
    method: 'POST',
    body: JSON.stringify(productData),
  });
};

// Usage
createProduct({
  name: 'Laptop',
  description: 'High performance laptop',
  price: 1200.50,
  quantity: 10,
  category: 'Electronics',
  sku: 'LAP-001'
});
```

### Get All Products
```javascript
const getProducts = async () => {
  return await apiCall('/products');
};
```

### Update Product (PUT)
```javascript
const updateProduct = async (productId, productData) => {
  return await apiCall(`/products/${productId}`, {
    method: 'PUT',
    body: JSON.stringify(productData),
  });
};

// Usage
updateProduct('693d65cf2664008c4c7795c0', {
  name: 'Gaming Laptop',
  price: 2500.99,
  quantity: 5
});
```

### Delete Product
```javascript
const deleteProduct = async (productId) => {
  return await apiCall(`/products/${productId}`, {
    method: 'DELETE',
  });
};
```

---

## 7️⃣ Error Handling Guide

### Common HTTP Status Codes

| Code | Meaning | Action |
|------|---------|--------|
| 200 | Success | Process response data |
| 201 | Created | Resource created successfully |
| 400 | Bad Request | Show validation errors to user |
| 401 | Unauthorized | Redirect to login, clear token |
| 403 | Forbidden | Show "Access Denied" message |
| 404 | Not Found | Show "Resource not found" |
| 500 | Server Error | Show "Server error, try again" |

### Error Handling Pattern

```javascript
async function handleApiRequest(apiFunction) {
  try {
    const data = await apiFunction();
    return { success: true, data };
  } catch (error) {
    let errorMessage = 'An error occurred';

    if (error.message) {
      errorMessage = error.message;
    }

    return { success: false, error: errorMessage };
  }
}

// Usage
const result = await handleApiRequest(() => loginUser(email, password));

if (result.success) {
  // Handle success
  console.log('Success:', result.data);
} else {
  // Handle error
  alert(result.error);
}
```

---

## 8️⃣ CORS Configuration

The API has CORS enabled with the following headers:
- `Access-Control-Allow-Origin: *`
- `Access-Control-Allow-Methods: POST, GET, PUT, DELETE, PATCH, OPTIONS`
- `Access-Control-Allow-Headers: Content-Type, Authorization, ...`

You can make requests from any domain during development.

---

## 9️⃣ Token Expiry

- JWT tokens expire after **24 hours** (configurable)
- When a token expires, API returns `401 Unauthorized`
- Frontend should detect this and redirect to login
- Always handle 401 responses by clearing auth state

---

## 🔟 Security Best Practices

### ✅ DO:
- Store tokens in `localStorage` or `httpOnly cookies`
- Always send tokens in `Authorization` header
- Clear tokens on logout
- Validate user input before sending to API
- Use HTTPS in production
- Implement token refresh mechanism (if needed)

### ❌ DON'T:
- Store passwords in localStorage
- Send tokens in URL parameters
- Log sensitive data to console in production
- Store tokens in regular cookies without httpOnly flag

---

## 📞 Support & Questions

If you encounter any issues:
1. Check the response status code and error message
2. Verify the token is being sent correctly
3. Check CORS configuration if requests fail
4. Ensure MongoDB is running on the backend

**API Base URL:** `http://localhost:8080/api/v1`

**Server Status Check:** `GET http://localhost:8080/api/v1/hello`

---

## 🎯 Quick Reference

### Authentication
```javascript
// Register
POST /api/v1/auth/register
Body: { name, email, password }

// Login
POST /api/v1/auth/login
Body: { email, password }
Response: { token, data }
```

### Protected Requests
```javascript
Headers: {
  'Authorization': 'Bearer <token>',
  'Content-Type': 'application/json'
}
```

### Store Token
```javascript
localStorage.setItem('authToken', token);
```

### Use Token
```javascript
const token = localStorage.getItem('authToken');
fetch(url, {
  headers: { 'Authorization': `Bearer ${token}` }
});
```

---

**Last Updated:** December 13, 2025
**API Version:** 1.0.0
