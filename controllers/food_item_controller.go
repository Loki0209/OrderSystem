package controllers

import (
	"net/http"

	"ordernew/models"
	"ordernew/services"

	"github.com/gin-gonic/gin"
)

// CreateFoodItem handles food item creation
func CreateFoodItem(c *gin.Context) {
	var req models.CreateFoodItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foodItem, err := services.CreateFoodItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Food item created successfully",
		"data":    foodItem.ToFoodItemResponse(),
	})
}

// GetFoodItem handles retrieving a single food item
func GetFoodItem(c *gin.Context) {
	foodItemID := c.Param("id")

	foodItem, err := services.GetFoodItemByID(foodItemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food item retrieved successfully",
		"data":    foodItem.ToFoodItemResponse(),
	})
}

// GetFoodItemsByStore handles retrieving all food items for a store
func GetFoodItemsByStore(c *gin.Context) {
	storeID := c.Param("storeId")

	foodItems, err := services.GetFoodItemsByStore(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var foodItemResponses []models.FoodItemResponse
	for _, foodItem := range foodItems {
		foodItemResponses = append(foodItemResponses, foodItem.ToFoodItemResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food items retrieved successfully",
		"count":   len(foodItemResponses),
		"data":    foodItemResponses,
	})
}

// GetFoodItemsByCategory handles retrieving all food items for a category
func GetFoodItemsByCategory(c *gin.Context) {
	categoryID := c.Param("categoryId")

	foodItems, err := services.GetFoodItemsByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var foodItemResponses []models.FoodItemResponse
	for _, foodItem := range foodItems {
		foodItemResponses = append(foodItemResponses, foodItem.ToFoodItemResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food items retrieved successfully",
		"count":   len(foodItemResponses),
		"data":    foodItemResponses,
	})
}

// GetAvailableFoodItemsByStore handles retrieving all available food items for a store (for customers)
func GetAvailableFoodItemsByStore(c *gin.Context) {
	storeID := c.Param("storeId")

	foodItems, err := services.GetAvailableFoodItemsByStore(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var foodItemResponses []models.FoodItemResponse
	for _, foodItem := range foodItems {
		foodItemResponses = append(foodItemResponses, foodItem.ToFoodItemResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Available food items retrieved successfully",
		"count":   len(foodItemResponses),
		"data":    foodItemResponses,
	})
}

// GetAvailableFoodItemsByCategory handles retrieving all available food items for a category (for customers)
func GetAvailableFoodItemsByCategory(c *gin.Context) {
	categoryID := c.Param("categoryId")

	foodItems, err := services.GetAvailableFoodItemsByCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var foodItemResponses []models.FoodItemResponse
	for _, foodItem := range foodItems {
		foodItemResponses = append(foodItemResponses, foodItem.ToFoodItemResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Available food items retrieved successfully",
		"count":   len(foodItemResponses),
		"data":    foodItemResponses,
	})
}

// UpdateFoodItem handles updating a food item
func UpdateFoodItem(c *gin.Context) {
	foodItemID := c.Param("id")

	var req models.UpdateFoodItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foodItem, err := services.UpdateFoodItem(foodItemID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food item updated successfully",
		"data":    foodItem.ToFoodItemResponse(),
	})
}

// DeleteFoodItem handles deleting a food item
func DeleteFoodItem(c *gin.Context) {
	foodItemID := c.Param("id")

	err := services.DeleteFoodItem(foodItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food item deleted successfully",
	})
}

// ToggleFoodItemAvailability handles toggling food item availability
func ToggleFoodItemAvailability(c *gin.Context) {
	foodItemID := c.Param("id")

	foodItem, err := services.ToggleFoodItemAvailability(foodItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status := "unavailable"
	if foodItem.IsAvailable {
		status = "available"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Food item availability updated successfully",
		"status":  status,
		"data":    foodItem.ToFoodItemResponse(),
	})
}
