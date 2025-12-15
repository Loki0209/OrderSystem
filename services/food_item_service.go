package services

import (
	"context"
	"errors"
	"time"

	"ordernew/config"
	"ordernew/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodItemCollection *mongo.Collection

// InitFoodItemCollection initializes the food item collection
func InitFoodItemCollection() {
	foodItemCollection = config.GetCollection("food_items")
}

// CreateFoodItem creates a new food item
func CreateFoodItem(req models.CreateFoodItemRequest) (*models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storeID, err := primitive.ObjectIDFromHex(req.StoreID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	categoryID, err := primitive.ObjectIDFromHex(req.CategoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}

	// Verify store exists
	_, err = GetStoreByID(req.StoreID)
	if err != nil {
		return nil, errors.New("store not found")
	}

	// Verify category exists
	_, err = GetCategoryByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	foodItem := &models.FoodItem{
		StoreID:      storeID,
		CategoryID:   categoryID,
		Name:         req.Name,
		Description:  req.Description,
		Price:        req.Price,
		Image:        req.Image,
		IsVeg:        req.IsVeg,
		IsAvailable:  true,
		IsActive:     true,
		PrepTime:     req.PrepTime,
		DisplayOrder: req.DisplayOrder,
		Tags:         req.Tags,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result, err := foodItemCollection.InsertOne(ctx, foodItem)
	if err != nil {
		return nil, err
	}

	foodItem.ID = result.InsertedID.(primitive.ObjectID)
	return foodItem, nil
}

// GetFoodItemByID retrieves a food item by ID
func GetFoodItemByID(foodItemID string) (*models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(foodItemID)
	if err != nil {
		return nil, errors.New("invalid food item ID")
	}

	var foodItem models.FoodItem
	err = foodItemCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&foodItem)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("food item not found")
		}
		return nil, err
	}

	return &foodItem, nil
}

// GetFoodItemsByStore retrieves all food items for a specific store
func GetFoodItemsByStore(storeID string) ([]models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	opts := options.Find().SetSort(bson.D{{Key: "display_order", Value: 1}})
	cursor, err := foodItemCollection.Find(ctx, bson.M{"store_id": objectID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var foodItems []models.FoodItem
	if err = cursor.All(ctx, &foodItems); err != nil {
		return nil, err
	}

	return foodItems, nil
}

// GetFoodItemsByCategory retrieves all food items for a specific category
func GetFoodItemsByCategory(categoryID string) ([]models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}

	opts := options.Find().SetSort(bson.D{{Key: "display_order", Value: 1}})
	cursor, err := foodItemCollection.Find(ctx, bson.M{"category_id": objectID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var foodItems []models.FoodItem
	if err = cursor.All(ctx, &foodItems); err != nil {
		return nil, err
	}

	return foodItems, nil
}

// GetAvailableFoodItemsByStore retrieves all available food items for a specific store
func GetAvailableFoodItemsByStore(storeID string) ([]models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	opts := options.Find().SetSort(bson.D{{Key: "display_order", Value: 1}})
	cursor, err := foodItemCollection.Find(ctx, bson.M{
		"store_id":     objectID,
		"is_available": true,
		"is_active":    true,
	}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var foodItems []models.FoodItem
	if err = cursor.All(ctx, &foodItems); err != nil {
		return nil, err
	}

	return foodItems, nil
}

// GetAvailableFoodItemsByCategory retrieves all available food items for a specific category
func GetAvailableFoodItemsByCategory(categoryID string) ([]models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}

	opts := options.Find().SetSort(bson.D{{Key: "display_order", Value: 1}})
	cursor, err := foodItemCollection.Find(ctx, bson.M{
		"category_id":  objectID,
		"is_available": true,
		"is_active":    true,
	}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var foodItems []models.FoodItem
	if err = cursor.All(ctx, &foodItems); err != nil {
		return nil, err
	}

	return foodItems, nil
}

// UpdateFoodItem updates an existing food item
func UpdateFoodItem(foodItemID string, req models.UpdateFoodItemRequest) (*models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(foodItemID)
	if err != nil {
		return nil, errors.New("invalid food item ID")
	}

	// Build update document
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if req.CategoryID != "" {
		categoryID, err := primitive.ObjectIDFromHex(req.CategoryID)
		if err != nil {
			return nil, errors.New("invalid category ID")
		}
		update["$set"].(bson.M)["category_id"] = categoryID
	}
	if req.Name != "" {
		update["$set"].(bson.M)["name"] = req.Name
	}
	if req.Description != "" {
		update["$set"].(bson.M)["description"] = req.Description
	}
	if req.Price != nil {
		update["$set"].(bson.M)["price"] = *req.Price
	}
	if req.Image != "" {
		update["$set"].(bson.M)["image"] = req.Image
	}
	if req.IsVeg != nil {
		update["$set"].(bson.M)["is_veg"] = *req.IsVeg
	}
	if req.IsAvailable != nil {
		update["$set"].(bson.M)["is_available"] = *req.IsAvailable
	}
	if req.IsActive != nil {
		update["$set"].(bson.M)["is_active"] = *req.IsActive
	}
	if req.PrepTime != nil {
		update["$set"].(bson.M)["prep_time"] = *req.PrepTime
	}
	if req.DisplayOrder != nil {
		update["$set"].(bson.M)["display_order"] = *req.DisplayOrder
	}
	if req.Tags != nil {
		update["$set"].(bson.M)["tags"] = req.Tags
	}

	_, err = foodItemCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return GetFoodItemByID(foodItemID)
}

// DeleteFoodItem deletes a food item
func DeleteFoodItem(foodItemID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(foodItemID)
	if err != nil {
		return errors.New("invalid food item ID")
	}

	result, err := foodItemCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("food item not found")
	}

	return nil
}

// ToggleFoodItemAvailability toggles the is_available status of a food item
func ToggleFoodItemAvailability(foodItemID string) (*models.FoodItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(foodItemID)
	if err != nil {
		return nil, errors.New("invalid food item ID")
	}

	// Get current status
	foodItem, err := GetFoodItemByID(foodItemID)
	if err != nil {
		return nil, err
	}

	// Toggle availability
	update := bson.M{
		"$set": bson.M{
			"is_available": !foodItem.IsAvailable,
			"updated_at":   time.Now(),
		},
	}

	_, err = foodItemCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return GetFoodItemByID(foodItemID)
}
