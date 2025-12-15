package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category represents a food category in the system
type Category struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StoreID     primitive.ObjectID `json:"store_id" bson:"store_id" binding:"required"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Description string             `json:"description" bson:"description"`
	Image       string             `json:"image" bson:"image"`
	DisplayOrder int               `json:"display_order" bson:"display_order"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// CreateCategoryRequest represents data for creating a category
type CreateCategoryRequest struct {
	StoreID      string `json:"store_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	DisplayOrder int    `json:"display_order"`
}

// UpdateCategoryRequest represents data for updating a category
type UpdateCategoryRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Image        string `json:"image"`
	DisplayOrder *int   `json:"display_order"`
	IsActive     *bool  `json:"is_active"`
}

// CategoryResponse represents the category data sent in responses
type CategoryResponse struct {
	ID          primitive.ObjectID `json:"id"`
	StoreID     primitive.ObjectID `json:"store_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Image       string             `json:"image"`
	DisplayOrder int               `json:"display_order"`
	IsActive    bool               `json:"is_active"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// ToCategoryResponse converts Category to CategoryResponse
func (c *Category) ToCategoryResponse() CategoryResponse {
	return CategoryResponse{
		ID:          c.ID,
		StoreID:     c.StoreID,
		Name:        c.Name,
		Description: c.Description,
		Image:       c.Image,
		DisplayOrder: c.DisplayOrder,
		IsActive:    c.IsActive,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
