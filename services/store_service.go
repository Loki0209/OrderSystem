package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"ordernew/config"
	"ordernew/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var storeCollection *mongo.Collection

func init() {
	// This will be set after database connection
	// storeCollection will be initialized in InitCollections
}

// InitStoreCollection initializes the store collection
func InitStoreCollection() {
	storeCollection = config.GetCollection("stores")
}

// CreateStore creates a new store
func CreateStore(req models.CreateStoreRequest, ownerID primitive.ObjectID) (*models.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create store
	store := &models.Store{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Phone:       req.Phone,
		Email:       req.Email,
		OwnerID:     ownerID,
		IsOpen:      true,
		IsActive:    true,
		OpeningTime: req.OpeningTime,
		ClosingTime: req.ClosingTime,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Generate QR code data (storeID will be set after insert)
	result, err := storeCollection.InsertOne(ctx, store)
	if err != nil {
		return nil, err
	}

	// Set the ID and generate QR code data
	store.ID = result.InsertedID.(primitive.ObjectID)
	store.QRCodeData = fmt.Sprintf("store_id=%s", store.ID.Hex())

	// Update with QR code data
	update := bson.M{
		"$set": bson.M{
			"qr_code_data": store.QRCodeData,
		},
	}
	_, err = storeCollection.UpdateOne(ctx, bson.M{"_id": store.ID}, update)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// GetStoreByID retrieves a store by ID
func GetStoreByID(storeID string) (*models.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	var store models.Store
	err = storeCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&store)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("store not found")
		}
		return nil, err
	}

	return &store, nil
}

// GetAllStores retrieves all stores
func GetAllStores() ([]models.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := storeCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stores []models.Store
	if err = cursor.All(ctx, &stores); err != nil {
		return nil, err
	}

	return stores, nil
}

// GetStoresByOwner retrieves all stores owned by a specific user
func GetStoresByOwner(ownerID primitive.ObjectID) ([]models.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := storeCollection.Find(ctx, bson.M{"owner_id": ownerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stores []models.Store
	if err = cursor.All(ctx, &stores); err != nil {
		return nil, err
	}

	return stores, nil
}

// UpdateStore updates an existing store
func UpdateStore(storeID string, req models.UpdateStoreRequest) (*models.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("invalid store ID")
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
	if req.Address != nil {
		update["$set"].(bson.M)["address"] = req.Address
	}
	if req.Phone != "" {
		update["$set"].(bson.M)["phone"] = req.Phone
	}
	if req.Email != "" {
		update["$set"].(bson.M)["email"] = req.Email
	}
	if req.Logo != "" {
		update["$set"].(bson.M)["logo"] = req.Logo
	}
	if req.IsOpen != nil {
		update["$set"].(bson.M)["is_open"] = *req.IsOpen
	}
	if req.IsActive != nil {
		update["$set"].(bson.M)["is_active"] = *req.IsActive
	}
	if req.OpeningTime != "" {
		update["$set"].(bson.M)["opening_time"] = req.OpeningTime
	}
	if req.ClosingTime != "" {
		update["$set"].(bson.M)["closing_time"] = req.ClosingTime
	}

	_, err = storeCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return GetStoreByID(storeID)
}

// DeleteStore deletes a store
func DeleteStore(storeID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.New("invalid store ID")
	}

	result, err := storeCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("store not found")
	}

	return nil
}

// ToggleStoreStatus toggles the is_open status of a store
func ToggleStoreStatus(storeID string) (*models.Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("invalid store ID")
	}

	// Get current status
	store, err := GetStoreByID(storeID)
	if err != nil {
		return nil, err
	}

	// Toggle status
	update := bson.M{
		"$set": bson.M{
			"is_open":    !store.IsOpen,
			"updated_at": time.Now(),
		},
	}

	_, err = storeCollection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}

	return GetStoreByID(storeID)
}
