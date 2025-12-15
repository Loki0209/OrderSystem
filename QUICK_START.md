# Quick Start Guide - Restaurant Ordering System

## 🚀 Getting Started in 5 Minutes

### Prerequisites
- Go 1.20+ installed
- MongoDB running on `localhost:27017`
- Postman (for testing)

---

## Step 1: Start the Server

```bash
# Navigate to project directory
cd d:\project\ordernew

# Run the application
go run main.go
```

You should see:
```
Configuration loaded successfully
Connected to MongoDB successfully!
Server starting on port 8080...
```

---

## Step 2: Import Postman Collection

1. Open Postman
2. Click **Import** button
3. Select file: `postman_collection.json`
4. Collection "Restaurant Ordering System API" will appear

---

## Step 3: Test the APIs

### 3.1 Register & Login

1. **Authentication → Register User (Owner/Admin)**
   - Click Send
   - User created ✅
   - `user_id` automatically saved

2. **Authentication → Login**
   - Click Send
   - JWT token returned ✅
   - `auth_token` automatically saved to variables

---

### 3.2 Create Your Restaurant

3. **Stores → Create Store**
   - Click Send
   - Store created with QR code data ✅
   - `store_id` automatically saved

Example response:
```json
{
  "message": "Store created successfully",
  "data": {
    "id": "675c4d8f...",
    "name": "The Gourmet Cafe",
    "qr_code_data": "store_id=675c4d8f...",
    ...
  }
}
```

---

### 3.3 Create Food Categories

4. **Categories → Create Category**
   - Category "Beverages" created ✅
   - `category_id` automatically saved

5. **Repeat** to create more categories:
   - Edit request body → Change name to "Appetizers"
   - Edit request body → Change name to "Main Course"
   - Edit request body → Change name to "Desserts"

---

### 3.4 Add Menu Items

6. **Food Items → Create Food Item**
   - Food item "Cappuccino" created ✅
   - `food_item_id` automatically saved

7. **Repeat** to add more items:
   - Change category_id to different categories
   - Update name, description, price
   - Try different tags: ["bestseller", "new", "spicy"]

---

### 3.5 View Customer Menu (Public Endpoint)

8. **Food Items → Get Available Food Items by Store**
   - No authentication needed! ✅
   - Returns all available menu items
   - This is what customers see after scanning QR code

---

## 🎯 Key Features to Test

### Toggle Store Status
**Stores → Toggle Store Status (Open/Close)**
- Opens or closes your restaurant
- Response shows current status

### Toggle Item Availability
**Food Items → Toggle Food Item Availability**
- Marks items as available/unavailable
- Use when items are out of stock

### Get My Stores
**Stores → Get My Stores**
- Shows all stores owned by you
- Useful for multi-store management

---

## 📋 Collection Variables (Auto-Saved)

These are automatically set after each API call:
- `{{auth_token}}` - JWT token from login
- `{{store_id}}` - Created store ID
- `{{category_id}}` - Created category ID
- `{{food_item_id}}` - Created food item ID
- `{{user_id}}` - Registered user ID

All subsequent requests use these variables automatically!

---

## 🌐 Public vs Protected Endpoints

### 🔓 Public (No Auth Required)
These work without `Authorization` header:
```
GET /stores/:id                               - Store details
GET /categories/store/:storeId/active         - Active categories
GET /food-items/store/:storeId/available      - Available menu items
```

### 🔒 Protected (Auth Required)
These need `Authorization: Bearer {{auth_token}}`:
```
POST   /stores                    - Create store
PUT    /stores/:id                - Update store
DELETE /stores/:id                - Delete store
POST   /categories                - Create category
POST   /food-items                - Create food item
PATCH  /food-items/:id/toggle...  - Toggle availability
```

---

## 🧪 Testing Workflow

### For Restaurant Owner:
```
1. Register → Login
2. Create Store
3. Create Categories
4. Add Food Items
5. Toggle availability as needed
6. Toggle store open/closed
```

### For Customer (Simulated):
```
1. Get Store by ID (QR code scan)
2. Get Active Categories
3. Get Available Food Items
4. Browse and select items
```

---

## 📊 Sample Data

### Sample Store:
```json
{
  "name": "The Gourmet Cafe",
  "description": "A premium cafe serving delicious coffee and pastries",
  "phone": "+1-234-567-8900",
  "email": "contact@gourmetcafe.com"
}
```

### Sample Categories:
- Beverages (Hot & Cold drinks)
- Appetizers (Starters)
- Main Course (Primary dishes)
- Desserts (Sweets)

### Sample Food Items:
- Cappuccino ($4.99) - Beverages
- French Fries ($3.99) - Appetizers
- Grilled Chicken ($12.99) - Main Course
- Chocolate Cake ($5.99) - Desserts

---

## 🐛 Common Issues

### 1. "Failed to connect to database"
**Solution:** Make sure MongoDB is running
```bash
# Windows
net start MongoDB

# Or check if it's running
tasklist | findstr mongod
```

### 2. "User not authenticated"
**Solution:** Make sure you ran the Login request first
- The auth_token should be automatically saved
- Check Collection Variables to verify token exists

### 3. "Store not found"
**Solution:** Make sure you created a store first
- Run "Create Store" request
- Verify store_id is saved in collection variables

### 4. Go command not found
**Solution:** Make sure Go is installed and in PATH
- Download from: https://golang.org/dl/
- Verify: `go version`

---

## 🎉 Success Indicators

You've successfully set up the system when:
- ✅ Server starts without errors
- ✅ Login returns a JWT token
- ✅ Store creation returns QR code data
- ✅ Public endpoints work without authentication
- ✅ Protected endpoints work with authentication
- ✅ MongoDB collections are created automatically

---

## 📖 Next Steps

1. **Read full documentation:** `API_DOCUMENTATION.md`
2. **Review implementation:** `IMPLEMENTATION_SUMMARY.md`
3. **Add more menu items** to your store
4. **Test all endpoints** in the Postman collection
5. **Ready to build frontend** or continue with backend features

---

## 🆘 Need Help?

- Check `API_DOCUMENTATION.md` for complete API reference
- Check `IMPLEMENTATION_SUMMARY.md` for architecture details
- Check `README.md` for project overview
- Examine Postman request examples for correct format

---

## 🎯 Quick Testing Checklist

- [ ] Server starts successfully
- [ ] Register user works
- [ ] Login returns token
- [ ] Create store works
- [ ] Create category works
- [ ] Create food item works
- [ ] Public endpoints accessible without auth
- [ ] Protected endpoints require auth
- [ ] Toggle functions work
- [ ] Get requests return proper data

**If all checkboxes are ticked, you're ready to go! 🚀**
