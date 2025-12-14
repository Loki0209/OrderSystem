package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product represents a product in the system
type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price" binding:"required,gt=0"`
	Quantity    int                `json:"quantity" bson:"quantity" binding:"required,gte=0"`
	Category    string             `json:"category" bson:"category"`
	SKU         string             `json:"sku" bson:"sku"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	CreatedBy   string             `json:"created_by" bson:"created_by"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// CreateProductRequest represents data for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Quantity    int     `json:"quantity" binding:"required,gte=0"`
	Category    string  `json:"category"`
	SKU         string  `json:"sku" binding:"required"`
}

// UpdateProductRequest represents data for updating a product
type UpdateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
	Quantity    *int     `json:"quantity" binding:"omitempty,gte=0"`
	Category    string   `json:"category"`
	SKU         string   `json:"sku"`
	IsActive    *bool    `json:"is_active"`
}

// ProductResponse represents the product data sent in responses
type ProductResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Quantity    int                `json:"quantity"`
	Category    string             `json:"category"`
	SKU         string             `json:"sku"`
	IsActive    bool               `json:"is_active"`
	CreatedBy   string             `json:"created_by"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// ToProductResponse converts Product to ProductResponse
func (p *Product) ToProductResponse() ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Category:    p.Category,
		SKU:         p.SKU,
		IsActive:    p.IsActive,
		CreatedBy:   p.CreatedBy,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
