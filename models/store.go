package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Store represents a restaurant/cafe in the system
type Store struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Description string             `json:"description" bson:"description"`
	Address     Address            `json:"address" bson:"address"`
	Phone       string             `json:"phone" bson:"phone" binding:"required"`
	Email       string             `json:"email" bson:"email" binding:"omitempty,email"`
	OwnerID     primitive.ObjectID `json:"owner_id" bson:"owner_id"`
	Logo        string             `json:"logo" bson:"logo"`
	IsOpen      bool               `json:"is_open" bson:"is_open"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	OpeningTime string             `json:"opening_time" bson:"opening_time"`
	ClosingTime string             `json:"closing_time" bson:"closing_time"`
	QRCodeData  string             `json:"qr_code_data" bson:"qr_code_data"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// Address represents store address
type Address struct {
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	ZipCode string `json:"zip_code" bson:"zip_code"`
	Country string `json:"country" bson:"country"`
}

// CreateStoreRequest represents data for creating a store
type CreateStoreRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Address     Address `json:"address"`
	Phone       string  `json:"phone" binding:"required"`
	Email       string  `json:"email" binding:"omitempty,email"`
	OpeningTime string  `json:"opening_time"`
	ClosingTime string  `json:"closing_time"`
}

// UpdateStoreRequest represents data for updating a store
type UpdateStoreRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Address     *Address `json:"address"`
	Phone       string   `json:"phone"`
	Email       string   `json:"email" binding:"omitempty,email"`
	Logo        string   `json:"logo"`
	IsOpen      *bool    `json:"is_open"`
	IsActive    *bool    `json:"is_active"`
	OpeningTime string   `json:"opening_time"`
	ClosingTime string   `json:"closing_time"`
}

// StoreResponse represents the store data sent in responses
type StoreResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Address     Address            `json:"address"`
	Phone       string             `json:"phone"`
	Email       string             `json:"email"`
	OwnerID     primitive.ObjectID `json:"owner_id"`
	Logo        string             `json:"logo"`
	IsOpen      bool               `json:"is_open"`
	IsActive    bool               `json:"is_active"`
	OpeningTime string             `json:"opening_time"`
	ClosingTime string             `json:"closing_time"`
	QRCodeData  string             `json:"qr_code_data"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// ToStoreResponse converts Store to StoreResponse
func (s *Store) ToStoreResponse() StoreResponse {
	return StoreResponse{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Address:     s.Address,
		Phone:       s.Phone,
		Email:       s.Email,
		OwnerID:     s.OwnerID,
		Logo:        s.Logo,
		IsOpen:      s.IsOpen,
		IsActive:    s.IsActive,
		OpeningTime: s.OpeningTime,
		ClosingTime: s.ClosingTime,
		QRCodeData:  s.QRCodeData,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}
