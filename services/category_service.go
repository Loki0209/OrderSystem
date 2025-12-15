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

var categoryCollection *mongo.Collection

// InitCategoryCollection initializes the category collection
func InitCategoryCollection() {
	categoryCollection = config.GetCollection("categories")
}

// CreateCategory creates a new category
func CreateCategory(req models.CreateCategoryRequest) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	storeID, err := primitive.ObjectIDFromHex(req.StoreID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	// Verify store exists
	_, err = GetStoreByID(req.StoreID)
	if err != nil {
		return nil, errors.New("store not found")
	}

	category := &models.Category{
		StoreID:      storeID,
		Name:         req.Name,
		Description:  req.Description,
		Image:        req.Image,
		DisplayOrder: req.DisplayOrder,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result, err := categoryCollection.InsertOne(ctx, category)
	if err != nil {
		return nil, err
	}

	category.ID = result.InsertedID.(primitive.ObjectID)
	return category, nil
}

// GetCategoryByID retrieves a category by ID
func GetCategoryByID(categoryID string) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}

	var category models.Category
	err = categoryCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &category, nil
}

// GetCategoriesByStore retrieves all categories for a specific store
func GetCategoriesByStore(storeID string) ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	// Sort by display_order
	opts := options.Find().SetSort(bson.D{{Key: "display_order", Value: 1}})
	cursor, err := categoryCollection.Find(ctx, bson.M{"store_id": objectID}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []models.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// GetActiveCategoriesByStore retrieves all active categories for a specific store
func GetActiveCategoriesByStore(storeID string) ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	// Sort by display_order
	opts := options.Find().SetSort(bson.D{{Key: "display_order", Value: 1}})
	cursor, err := categoryCollection.Find(ctx, bson.M{
		"store_id":  objectID,
		"is_active": true,
	}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []models.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

// UpdateCategory updates an existing category
func UpdateCategory(categoryID string, req models.UpdateCategoryRequest) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return nil, errors.New("invalid category ID")
	}

	// Build update document
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if req.Name != "" {
		update["$set"].(bson.M)["name"] = req.Name
	}
	if req.Description != "" {
		update["$set"].(bson.M)["description"] = req.Description
	}
	if req.Image != "" {
		update["$set"].(bson.M)["image"] = req.Image
	}
	if req.DisplayOrder != nil {
		update["$set"].(bson.M)["display_order"] = *req.DisplayOrder
	}
	if req.IsActive != nil {
		update["$set"].(bson.M)["is_active"] = *req.IsActive
	}

	_, err = categoryCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return GetCategoryByID(categoryID)
}

// DeleteCategory deletes a category
func DeleteCategory(categoryID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(categoryID)
	if err != nil {
		return errors.New("invalid category ID")
	}

	result, err := categoryCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("category not found")
	}

	return nil
}
