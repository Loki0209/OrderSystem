package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FoodItem represents a food item in the system
type FoodItem struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StoreID     primitive.ObjectID `json:"store_id" bson:"store_id" binding:"required"`
	CategoryID  primitive.ObjectID `json:"category_id" bson:"category_id" binding:"required"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price" binding:"required,gt=0"`
	Image       string             `json:"image" bson:"image"`
	IsVeg       bool               `json:"is_veg" bson:"is_veg"`
	IsAvailable bool               `json:"is_available" bson:"is_available"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	PrepTime    int                `json:"prep_time" bson:"prep_time"` // in minutes
	DisplayOrder int               `json:"display_order" bson:"display_order"`
	Tags        []string           `json:"tags" bson:"tags"` // e.g., "spicy", "bestseller", "new"
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// CreateFoodItemRequest represents data for creating a food item
type CreateFoodItemRequest struct {
	StoreID      string   `json:"store_id" binding:"required"`
	CategoryID   string   `json:"category_id" binding:"required"`
	Name         string   `json:"name" binding:"required"`
	Description  string   `json:"description"`
	Price        float64  `json:"price" binding:"required,gt=0"`
	Image        string   `json:"image"`
	IsVeg        bool     `json:"is_veg"`
	PrepTime     int      `json:"prep_time"`
	DisplayOrder int      `json:"display_order"`
	Tags         []string `json:"tags"`
}

// UpdateFoodItemRequest represents data for updating a food item
type UpdateFoodItemRequest struct {
	CategoryID   string    `json:"category_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        *float64  `json:"price" binding:"omitempty,gt=0"`
	Image        string    `json:"image"`
	IsVeg        *bool     `json:"is_veg"`
	IsAvailable  *bool     `json:"is_available"`
	IsActive     *bool     `json:"is_active"`
	PrepTime     *int      `json:"prep_time"`
	DisplayOrder *int      `json:"display_order"`
	Tags         []string  `json:"tags"`
}

// FoodItemResponse represents the food item data sent in responses
type FoodItemResponse struct {
	ID          primitive.ObjectID `json:"id"`
	StoreID     primitive.ObjectID `json:"store_id"`
	CategoryID  primitive.ObjectID `json:"category_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Image       string             `json:"image"`
	IsVeg       bool               `json:"is_veg"`
	IsAvailable bool               `json:"is_available"`
	IsActive    bool               `json:"is_active"`
	PrepTime    int                `json:"prep_time"`
	DisplayOrder int               `json:"display_order"`
	Tags        []string           `json:"tags"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// ToFoodItemResponse converts FoodItem to FoodItemResponse
func (f *FoodItem) ToFoodItemResponse() FoodItemResponse {
	return FoodItemResponse{
		ID:          f.ID,
		StoreID:     f.StoreID,
		CategoryID:  f.CategoryID,
		Name:        f.Name,
		Description: f.Description,
		Price:       f.Price,
		Image:       f.Image,
		IsVeg:       f.IsVeg,
		IsAvailable: f.IsAvailable,
		IsActive:    f.IsActive,
		PrepTime:    f.PrepTime,
		DisplayOrder: f.DisplayOrder,
		Tags:        f.Tags,
		CreatedAt:   f.CreatedAt,
		UpdatedAt:   f.UpdatedAt,
	}
}
