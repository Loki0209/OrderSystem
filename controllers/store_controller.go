package controllers

import (
	"net/http"

	"ordernew/models"
	"ordernew/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateStore handles store creation
func CreateStore(c *gin.Context) {
	var req models.CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get owner ID from JWT token (stored in context by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	ownerID := userID.(primitive.ObjectID)

	store, err := services.CreateStore(req, ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Store created successfully",
		"data":    store.ToStoreResponse(),
	})
}

// GetStore handles retrieving a single store
func GetStore(c *gin.Context) {
	storeID := c.Param("id")

	store, err := services.GetStoreByID(storeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Store retrieved successfully",
		"data":    store.ToStoreResponse(),
	})
}

// GetAllStores handles retrieving all stores
func GetAllStores(c *gin.Context) {
	stores, err := services.GetAllStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var storeResponses []models.StoreResponse
	for _, store := range stores {
		storeResponses = append(storeResponses, store.ToStoreResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Stores retrieved successfully",
		"count":   len(storeResponses),
		"data":    storeResponses,
	})
}

// GetMyStores handles retrieving stores owned by the authenticated user
func GetMyStores(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	ownerID := userID.(primitive.ObjectID)

	stores, err := services.GetStoresByOwner(ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var storeResponses []models.StoreResponse
	for _, store := range stores {
		storeResponses = append(storeResponses, store.ToStoreResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your stores retrieved successfully",
		"count":   len(storeResponses),
		"data":    storeResponses,
	})
}

// UpdateStore handles updating a store
func UpdateStore(c *gin.Context) {
	storeID := c.Param("id")

	var req models.UpdateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store, err := services.UpdateStore(storeID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Store updated successfully",
		"data":    store.ToStoreResponse(),
	})
}

// DeleteStore handles deleting a store
func DeleteStore(c *gin.Context) {
	storeID := c.Param("id")

	err := services.DeleteStore(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Store deleted successfully",
	})
}

// ToggleStoreStatus handles toggling store open/closed status
func ToggleStoreStatus(c *gin.Context) {
	storeID := c.Param("id")

	store, err := services.ToggleStoreStatus(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status := "closed"
	if store.IsOpen {
		status = "open"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Store status updated successfully",
		"status":  status,
		"data":    store.ToStoreResponse(),
	})
}
