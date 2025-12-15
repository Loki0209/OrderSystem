package controllers

import (
	"net/http"

	"ordernew/models"
	"ordernew/services"

	"github.com/gin-gonic/gin"
)

// CreateCategory handles category creation
func CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.CreateCategory(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Category created successfully",
		"data":    category.ToCategoryResponse(),
	})
}

// GetCategory handles retrieving a single category
func GetCategory(c *gin.Context) {
	categoryID := c.Param("id")

	category, err := services.GetCategoryByID(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category retrieved successfully",
		"data":    category.ToCategoryResponse(),
	})
}

// GetCategoriesByStore handles retrieving all categories for a store
func GetCategoriesByStore(c *gin.Context) {
	storeID := c.Param("storeId")

	categories, err := services.GetCategoriesByStore(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categoryResponses []models.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, category.ToCategoryResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categories retrieved successfully",
		"count":   len(categoryResponses),
		"data":    categoryResponses,
	})
}

// GetActiveCategoriesByStore handles retrieving all active categories for a store (for customers)
func GetActiveCategoriesByStore(c *gin.Context) {
	storeID := c.Param("storeId")

	categories, err := services.GetActiveCategoriesByStore(storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var categoryResponses []models.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, category.ToCategoryResponse())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Active categories retrieved successfully",
		"count":   len(categoryResponses),
		"data":    categoryResponses,
	})
}

// UpdateCategory handles updating a category
func UpdateCategory(c *gin.Context) {
	categoryID := c.Param("id")

	var req models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.UpdateCategory(categoryID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category updated successfully",
		"data":    category.ToCategoryResponse(),
	})
}

// DeleteCategory handles deleting a category
func DeleteCategory(c *gin.Context) {
	categoryID := c.Param("id")

	err := services.DeleteCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}
