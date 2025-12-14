# Order Management System - Go REST API

A production-ready monolithic REST API built with Go, Gin framework, and MongoDB.

## 🏗️ Project Structure

```
ordernew/
├── config/             # Configuration and database setup
│   ├── config.go      # Environment configuration
│   └── database.go    # MongoDB connection
├── controllers/        # HTTP request handlers
│   └── user_controller.go
├── middleware/         # HTTP middleware
│   ├── auth_middleware.go
│   └── cors_middleware.go
├── models/            # Data models and schemas
│   └── user.go
├── routes/            # API route definitions
│   └── routes.go
├── services/          # Business logic layer
│   └── user_service.go
├── utils/             # Utility functions
│   ├── jwt.go        # JWT token utilities
│   └── password.go   # Password hashing utilities
├── .env              # Environment variables
├── main.go           # Application entry point
├── go.mod            # Go module dependencies
└── README.md         # Project documentation
```

## 🚀 Features

- ✅ RESTful API with Gin framework
- ✅ MongoDB database integration
- ✅ JWT-based authentication
- ✅ Password hashing with bcrypt
- ✅ CRUD operations for users
- ✅ Middleware for authentication and CORS
- ✅ Clean architecture with separation of concerns
- ✅ Environment-based configuration
- ✅ Graceful shutdown handling

## 📋 Prerequisites

Before running this project, make sure you have:

- **Go** (version 1.20 or higher) - [Download](https://golang.org/dl/)
- **MongoDB** (version 4.4 or higher) - [Download](https://www.mongodb.com/try/download/community)

## 🔧 Installation

### 1. Clone or Navigate to Project

```bash
cd c:\Users\acer\OneDrive\Desktop\project\ordernew
```

### 2. Install Dependencies

All dependencies are already installed:
- github.com/gin-gonic/gin
- go.mongodb.org/mongo-driver
- github.com/golang-jwt/jwt/v5
- github.com/joho/godotenv
- golang.org/x/crypto

To verify or reinstall:
```bash
go mod download
go mod tidy
```

### 3. Configure Environment

The `.env` file is already created with default values:

```env
PORT=8080
GIN_MODE=debug
MONGODB_URI=mongodb://localhost:27017
DATABASE_NAME=ordernew_db
JWT_SECRET=your-secret-key-change-this-in-production
JWT_EXPIRY=24h
API_VERSION=v1
```

**Important:** Change `JWT_SECRET` in production!

### 4. Start MongoDB

Make sure MongoDB is running on your system:

**Windows:**
```bash
# Start MongoDB service
net start MongoDB

# Or run MongoDB manually
"C:\Program Files\MongoDB\Server\7.0\bin\mongod.exe" --dbpath "C:\data\db"
```

**Linux/Mac:**
```bash
# Start MongoDB service
sudo systemctl start mongod

# Or
brew services start mongodb-community
```

## ▶️ Running the Application

### Start the Server

```bash
go run main.go
```

You should see:
```
Configuration loaded successfully
Connected to MongoDB successfully!
Server starting on port 8080...
API Documentation available at http://localhost:8080/api/v1/hello
```

### Build and Run (Production)

```bash
# Build the binary
go build -o ordernew.exe

# Run the binary
./ordernew.exe
```

## 📚 API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Public Endpoints

#### 1. Root Endpoint
```http
GET /
```
Returns API information and available endpoints.

#### 2. Hello World (Health Check)
```http
GET /api/v1/hello
```

**Response:**
```json
{
  "message": "Hello World! API is running successfully.",
  "version": "1.0.0",
  "status": "active"
}
```

#### 3. Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "message": "User registered successfully",
  "data": {
    "id": "6579a1b2c3d4e5f6g7h8i9j0",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-13T10:00:00Z",
    "updated_at": "2025-12-13T10:00:00Z"
  }
}
```

#### 4. Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "data": {
    "id": "6579a1b2c3d4e5f6g7h8i9j0",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "is_active": true
  }
}
```

### Protected Endpoints (Require Authentication)

**Note:** Include the JWT token in the Authorization header:
```
Authorization: Bearer <your-token>
```

#### 5. Get All Users
```http
GET /api/v1/users
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Users retrieved successfully",
  "count": 2,
  "data": [...]
}
```

#### 6. Get User by ID
```http
GET /api/v1/users/:id
Authorization: Bearer <token>
```

#### 7. Update User
```http
PUT /api/v1/users/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "John Doe Updated",
  "role": "admin",
  "is_active": true
}
```

#### 8. Delete User
```http
DELETE /api/v1/users/:id
Authorization: Bearer <token>
```

## 🧪 Testing the API

### Using cURL

#### Test Hello World:
```bash
curl http://localhost:8080/api/v1/hello
```

#### Register a User:
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"John Doe\",\"email\":\"john@example.com\",\"password\":\"password123\"}"
```

#### Login:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"john@example.com\",\"password\":\"password123\"}"
```

#### Get All Users (with token):
```bash
curl http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### Using the Test Script

Run the automated test script:
```bash
bash test_api.sh
```

### Using Postman

1. Import the following as a new collection
2. Set base URL: `http://localhost:8080/api/v1`
3. For protected routes, add header: `Authorization: Bearer <token>`

## 🔐 Security Features

- **Password Hashing**: Using bcrypt with default cost
- **JWT Authentication**: Secure token-based authentication
- **CORS**: Configured for cross-origin requests
- **Environment Variables**: Sensitive data in `.env` file
- **Input Validation**: Request validation using Gin binding

## 🛠️ Development

### Adding New Features

1. **Model**: Add to `models/`
2. **Service**: Add business logic to `services/`
3. **Controller**: Add HTTP handlers to `controllers/`
4. **Routes**: Register routes in `routes/routes.go`

### Database Collections

The application uses the following MongoDB collections:
- `users`: User accounts and authentication

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | Server port | 8080 |
| GIN_MODE | Gin mode (debug/release) | debug |
| MONGODB_URI | MongoDB connection string | mongodb://localhost:27017 |
| DATABASE_NAME | Database name | ordernew_db |
| JWT_SECRET | Secret key for JWT | your-secret-key |
| JWT_EXPIRY | Token expiration time | 24h |
| API_VERSION | API version | v1 |

## 🐛 Troubleshooting

### MongoDB Connection Failed
```
Failed to connect to MongoDB
```
**Solution:** Make sure MongoDB is running on `localhost:27017`

### Port Already in Use
```
Failed to start server: address already in use
```
**Solution:** Change PORT in `.env` or kill the process using port 8080

### Module Errors
```
go: cannot find module
```
**Solution:** Run `go mod tidy` and `go mod download`

## 📦 Dependencies

- **Gin**: HTTP web framework
- **MongoDB Driver**: Official MongoDB driver for Go
- **JWT**: JSON Web Token implementation
- **GoDotEnv**: Environment variable loader
- **Bcrypt**: Password hashing library

## 🚀 Deployment

### Building for Production

```bash
# Set production mode
export GIN_MODE=release

# Build binary
go build -o ordernew

# Run
./ordernew
```

### Docker (Optional)

Create a `Dockerfile`:
```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
```

## 👨‍💻 Author

Created with Claude Code

## 📄 License

This project is open source and available under the MIT License.

## 🎯 Next Steps

- [ ] Add more entities (Orders, Products, etc.)
- [ ] Implement pagination for list endpoints
- [ ] Add request logging middleware
- [ ] Add unit tests
- [ ] Add Swagger/OpenAPI documentation
- [ ] Implement rate limiting
- [ ] Add email verification
- [ ] Add password reset functionality
