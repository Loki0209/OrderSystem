# Restaurant Ordering System - API Documentation

## Overview
Complete REST API for a Restaurant/Cafe Ordering System built with Go, Gin framework, and MongoDB.

## Base URL
```
http://localhost:8080/api/v1
```

## MongoDB Collections Created
- **stores** - Restaurant/cafe information
- **categories** - Food categories for each store
- **food_items** - Menu items with pricing and availability
- **users** - User accounts (owners and customers)
- **products** - Legacy product collection (kept for backward compatibility)

---

## Authentication

All protected endpoints require a JWT token in the Authorization header:
```
Authorization: Bearer <your-token>
```

### Register User
**POST** `/auth/register`

Create a new user account (restaurant owner/admin).

**Request Body:**
```json
{
  "name": "Restaurant Owner",
  "email": "owner@restaurant.com",
  "password": "password123"
}
```

**Response:** `201 Created`
```json
{
  "message": "User registered successfully",
  "data": {
    "id": "675c123...",
    "name": "Restaurant Owner",
    "email": "owner@restaurant.com",
    "phone": "",
    "role": "user",
    "is_active": true,
    "created_at": "2025-12-15T10:00:00Z",
    "updated_at": "2025-12-15T10:00:00Z"
  }
}
```

### Login
**POST** `/auth/login`

Authenticate and receive JWT token.

**Request Body:**
```json
{
  "email": "owner@restaurant.com",
  "password": "password123"
}
```

**Response:** `200 OK`
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "data": {
    "id": "675c123...",
    "name": "Restaurant Owner",
    "email": "owner@restaurant.com",
    "role": "user",
    "is_active": true
  }
}
```

---

## Stores API

### Create Store
**POST** `/stores` 🔒 (Requires Authentication)

Create a new restaurant/cafe store.

**Request Body:**
```json
{
  "name": "The Gourmet Cafe",
  "description": "A premium cafe serving delicious coffee and pastries",
  "address": {
    "street": "123 Main Street",
    "city": "New York",
    "state": "NY",
    "zip_code": "10001",
    "country": "USA"
  },
  "phone": "+1-234-567-8900",
  "email": "contact@gourmetcafe.com",
  "opening_time": "08:00 AM",
  "closing_time": "10:00 PM"
}
```

**Response:** `201 Created`
```json
{
  "message": "Store created successfully",
  "data": {
    "id": "675c456...",
    "name": "The Gourmet Cafe",
    "description": "A premium cafe serving delicious coffee and pastries",
    "address": { /* address object */ },
    "phone": "+1-234-567-8900",
    "email": "contact@gourmetcafe.com",
    "owner_id": "675c123...",
    "logo": "",
    "is_open": true,
    "is_active": true,
    "opening_time": "08:00 AM",
    "closing_time": "10:00 PM",
    "qr_code_data": "store_id=675c456...",
    "created_at": "2025-12-15T10:00:00Z",
    "updated_at": "2025-12-15T10:00:00Z"
  }
}
```

### Get All Stores
**GET** `/stores` 🌐 (Public)

Retrieve all stores.

**Response:** `200 OK`
```json
{
  "message": "Stores retrieved successfully",
  "count": 2,
  "data": [ /* array of store objects */ ]
}
```

### Get Store by ID
**GET** `/stores/:id` 🌐 (Public)

Get specific store details (used for QR code scanning).

**Response:** `200 OK`
```json
{
  "message": "Store retrieved successfully",
  "data": { /* store object */ }
}
```

### Get My Stores
**GET** `/stores/my-stores` 🔒 (Requires Authentication)

Get all stores owned by the authenticated user.

**Response:** `200 OK`
```json
{
  "message": "Your stores retrieved successfully",
  "count": 1,
  "data": [ /* array of store objects */ ]
}
```

### Update Store
**PUT** `/stores/:id` 🔒 (Requires Authentication)

Update store details.

**Request Body:**
```json
{
  "name": "The Gourmet Cafe - Updated",
  "description": "A premium cafe serving delicious coffee, pastries, and sandwiches",
  "opening_time": "07:00 AM",
  "closing_time": "11:00 PM"
}
```

**Response:** `200 OK`

### Toggle Store Status
**PATCH** `/stores/:id/toggle-status` 🔒 (Requires Authentication)

Toggle store open/closed status.

**Response:** `200 OK`
```json
{
  "message": "Store status updated successfully",
  "status": "open",
  "data": { /* store object */ }
}
```

### Delete Store
**DELETE** `/stores/:id` 🔒 (Requires Authentication)

Delete a store.

**Response:** `200 OK`

---

## Categories API

### Create Category
**POST** `/categories` 🔒 (Requires Authentication)

Create a new food category for a store.

**Request Body:**
```json
{
  "store_id": "675c456...",
  "name": "Beverages",
  "description": "Hot and cold beverages",
  "image": "https://example.com/beverages.jpg",
  "display_order": 1
}
```

**Response:** `201 Created`
```json
{
  "message": "Category created successfully",
  "data": {
    "id": "675c789...",
    "store_id": "675c456...",
    "name": "Beverages",
    "description": "Hot and cold beverages",
    "image": "https://example.com/beverages.jpg",
    "display_order": 1,
    "is_active": true,
    "created_at": "2025-12-15T10:00:00Z",
    "updated_at": "2025-12-15T10:00:00Z"
  }
}
```

### Get Categories by Store
**GET** `/categories/store/:storeId` 🌐 (Public)

Get all categories for a specific store.

**Response:** `200 OK`
```json
{
  "message": "Categories retrieved successfully",
  "count": 3,
  "data": [ /* array of category objects */ ]
}
```

### Get Active Categories by Store
**GET** `/categories/store/:storeId/active` 🌐 (Public)

Get only active categories (for customer menu display).

**Response:** `200 OK`

### Get Category by ID
**GET** `/categories/:id` 🌐 (Public)

Get specific category details.

**Response:** `200 OK`

### Update Category
**PUT** `/categories/:id` 🔒 (Requires Authentication)

Update category details.

**Request Body:**
```json
{
  "name": "Hot & Cold Beverages",
  "description": "Premium hot and cold beverages",
  "display_order": 1,
  "is_active": true
}
```

**Response:** `200 OK`

### Delete Category
**DELETE** `/categories/:id` 🔒 (Requires Authentication)

Delete a category.

**Response:** `200 OK`

---

## Food Items API

### Create Food Item
**POST** `/food-items` 🔒 (Requires Authentication)

Create a new food item in a category.

**Request Body:**
```json
{
  "store_id": "675c456...",
  "category_id": "675c789...",
  "name": "Cappuccino",
  "description": "Rich espresso with steamed milk and foam",
  "price": 4.99,
  "image": "https://example.com/cappuccino.jpg",
  "is_veg": true,
  "prep_time": 5,
  "display_order": 1,
  "tags": ["bestseller", "hot"]
}
```

**Response:** `201 Created`
```json
{
  "message": "Food item created successfully",
  "data": {
    "id": "675cabc...",
    "store_id": "675c456...",
    "category_id": "675c789...",
    "name": "Cappuccino",
    "description": "Rich espresso with steamed milk and foam",
    "price": 4.99,
    "image": "https://example.com/cappuccino.jpg",
    "is_veg": true,
    "is_available": true,
    "is_active": true,
    "prep_time": 5,
    "display_order": 1,
    "tags": ["bestseller", "hot"],
    "created_at": "2025-12-15T10:00:00Z",
    "updated_at": "2025-12-15T10:00:00Z"
  }
}
```

### Get Food Items by Store
**GET** `/food-items/store/:storeId` 🌐 (Public)

Get all food items for a store.

**Response:** `200 OK`
```json
{
  "message": "Food items retrieved successfully",
  "count": 15,
  "data": [ /* array of food item objects */ ]
}
```

### Get Available Food Items by Store
**GET** `/food-items/store/:storeId/available` 🌐 (Public)

Get only available food items (for customer menu).

**Response:** `200 OK`

### Get Food Items by Category
**GET** `/food-items/category/:categoryId` 🌐 (Public)

Get all food items in a specific category.

**Response:** `200 OK`

### Get Available Food Items by Category
**GET** `/food-items/category/:categoryId/available` 🌐 (Public)

Get only available food items in a category.

**Response:** `200 OK`

### Get Food Item by ID
**GET** `/food-items/:id` 🌐 (Public)

Get specific food item details.

**Response:** `200 OK`

### Update Food Item
**PUT** `/food-items/:id` 🔒 (Requires Authentication)

Update food item details.

**Request Body:**
```json
{
  "name": "Premium Cappuccino",
  "description": "Rich espresso with steamed milk, foam, and chocolate sprinkles",
  "price": 5.49,
  "is_available": true,
  "prep_time": 6,
  "tags": ["bestseller", "hot", "premium"]
}
```

**Response:** `200 OK`

### Toggle Food Item Availability
**PATCH** `/food-items/:id/toggle-availability` 🔒 (Requires Authentication)

Toggle food item availability status.

**Response:** `200 OK`
```json
{
  "message": "Food item availability updated successfully",
  "status": "available",
  "data": { /* food item object */ }
}
```

### Delete Food Item
**DELETE** `/food-items/:id` 🔒 (Requires Authentication)

Delete a food item.

**Response:** `200 OK`

---

## Users API

### Get All Users
**GET** `/users` 🔒 (Requires Authentication)

Get all registered users.

**Response:** `200 OK`

### Get User by ID
**GET** `/users/:id` 🔒 (Requires Authentication)

Get specific user details.

**Response:** `200 OK`

### Update User
**PUT** `/users/:id` 🔒 (Requires Authentication)

Update user details.

**Request Body:**
```json
{
  "name": "Restaurant Owner - Updated",
  "role": "admin",
  "is_active": true
}
```

**Response:** `200 OK`

### Delete User
**DELETE** `/users/:id` 🔒 (Requires Authentication)

Delete a user.

**Response:** `200 OK`

---

## Testing with Postman

1. Import the `postman_collection.json` file into Postman
2. Collection variables are automatically set after API calls:
   - `auth_token` - Set after login
   - `store_id` - Set after creating a store
   - `category_id` - Set after creating a category
   - `food_item_id` - Set after creating a food item
   - `user_id` - Set after registration/login

### Recommended Testing Flow:

1. **Register User** → Creates owner account
2. **Login** → Get auth token (auto-saved to variable)
3. **Create Store** → Create your restaurant (auto-saves store_id)
4. **Create Category** → Create food category (auto-saves category_id)
5. **Create Food Item** → Add menu items (auto-saves food_item_id)
6. **Get Available Food Items by Store** → View customer menu

---

## Error Responses

All endpoints return consistent error responses:

**400 Bad Request**
```json
{
  "error": "Validation error message"
}
```

**401 Unauthorized**
```json
{
  "error": "User not authenticated"
}
```

**404 Not Found**
```json
{
  "error": "Resource not found"
}
```

**500 Internal Server Error**
```json
{
  "error": "Internal server error message"
}
```

---

## Running the Application

1. **Start MongoDB:**
   ```bash
   # Make sure MongoDB is running on localhost:27017
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```
   Or build and run:
   ```bash
   go build -o ordernew.exe
   ./ordernew.exe
   ```

3. **Verify it's running:**
   ```
   GET http://localhost:8080/api/v1/hello
   ```

---

## Database Schema

### stores
```javascript
{
  _id: ObjectId,
  name: String,
  description: String,
  address: {
    street: String,
    city: String,
    state: String,
    zip_code: String,
    country: String
  },
  phone: String,
  email: String,
  owner_id: ObjectId,
  logo: String,
  is_open: Boolean,
  is_active: Boolean,
  opening_time: String,
  closing_time: String,
  qr_code_data: String,
  created_at: Date,
  updated_at: Date
}
```

### categories
```javascript
{
  _id: ObjectId,
  store_id: ObjectId,
  name: String,
  description: String,
  image: String,
  display_order: Number,
  is_active: Boolean,
  created_at: Date,
  updated_at: Date
}
```

### food_items
```javascript
{
  _id: ObjectId,
  store_id: ObjectId,
  category_id: ObjectId,
  name: String,
  description: String,
  price: Number,
  image: String,
  is_veg: Boolean,
  is_available: Boolean,
  is_active: Boolean,
  prep_time: Number,
  display_order: Number,
  tags: [String],
  created_at: Date,
  updated_at: Date
}
```

### users
```javascript
{
  _id: ObjectId,
  name: String,
  email: String,
  phone: String,
  password: String (hashed),
  role: String,
  is_active: Boolean,
  created_at: Date,
  updated_at: Date
}
```

---

## Next Steps (Future Enhancements)

- [ ] Phone OTP authentication for customers
- [ ] Shopping cart APIs
- [ ] Order management APIs
- [ ] Payment gateway integration
- [ ] Real-time order status updates (WebSocket/SSE)
- [ ] Image upload for stores, categories, and food items
- [ ] SMS notifications
- [ ] Analytics and reporting APIs
- [ ] Customer order history
- [ ] Rating and review system

---

## Support

For issues or questions, please refer to the README.md file or contact the development team.
