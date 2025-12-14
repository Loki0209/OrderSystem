package controllers

import (
	"net/http"

	"ordernew/models"
	"ordernew/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *services.ProductService
}

// NewProductController creates a new product controller instance
func NewProductController() *ProductController {
	return &ProductController{
		productService: services.NewProductService(),
	}
}

// CreateProduct creates a new product
// @Summary Create a new product
// @Description Create a new product in the database
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.CreateProductRequest true "Product data"
// @Success 201 {object} models.ProductResponse
// @Router /products [post]
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req models.CreateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := ctx.Get("user_id")
	if !exists {
		userID = "system"
	}

	product, err := c.productService.CreateProduct(req, userID.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to create product",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    product,
	})
}

// GetProductByID retrieves a product by ID
// @Summary Get product by ID
// @Description Get product details by product ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.ProductResponse
// @Router /products/{id} [get]
func (c *ProductController) GetProductByID(ctx *gin.Context) {
	productID := ctx.Param("id")

	product, err := c.productService.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Product not found",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

// GetAllProducts retrieves all products
// @Summary Get all products
// @Description Get list of all products with optional filters
// @Tags products
// @Produce json
// @Param category query string false "Filter by category"
// @Param is_active query bool false "Filter by active status"
// @Success 200 {array} models.ProductResponse
// @Router /products [get]
func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	category := ctx.Query("category")
	isActiveStr := ctx.Query("is_active")

	var isActive *bool
	if isActiveStr != "" {
		val := isActiveStr == "true"
		isActive = &val
	}

	products, err := c.productService.GetAllProducts(category, isActive)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch products",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Products retrieved successfully",
		"count":   len(products),
		"data":    products,
	})
}

// UpdateProduct updates product information (PUT - complete update)
// @Summary Update product (PUT)
// @Description Update product details - replaces entire resource
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body models.UpdateProductRequest true "Product update data"
// @Success 200 {object} models.ProductResponse
// @Router /products/{id} [put]
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	productID := ctx.Param("id")
	var req models.UpdateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	product, err := c.productService.UpdateProduct(productID, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Update failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"data":    product,
	})
}

// PatchProduct partially updates product information (PATCH - partial update)
// @Summary Patch product (PATCH)
// @Description Partially update product details
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.ProductResponse
// @Router /products/{id} [patch]
func (c *ProductController) PatchProduct(ctx *gin.Context) {
	productID := ctx.Param("id")

	var updates map[string]interface{}
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	product, err := c.productService.PatchProduct(productID, updates)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Patch failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product patched successfully",
		"data":    product,
	})
}

// DeleteProduct deletes a product
// @Summary Delete product
// @Description Delete a product by ID
// @Tags products
// @Param id path string true "Product ID"
// @Success 200 {object} map[string]string
// @Router /products/{id} [delete]
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	productID := ctx.Param("id")

	err := c.productService.DeleteProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Delete failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

// UpdateProductQuantity updates product quantity
// @Summary Update product quantity
// @Description Update only the quantity of a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.ProductResponse
// @Router /products/{id}/quantity [put]
func (c *ProductController) UpdateProductQuantity(ctx *gin.Context) {
	productID := ctx.Param("id")

	var req struct {
		Quantity int `json:"quantity" binding:"required,gte=0"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": err.Error(),
		})
		return
	}

	product, err := c.productService.UpdateProductQuantity(productID, req.Quantity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Update failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product quantity updated successfully",
		"data":    product,
	})
}

// GetProductsByCategory retrieves products by category
// @Summary Get products by category
// @Description Get all products in a specific category
// @Tags products
// @Produce json
// @Param category path string true "Category name"
// @Success 200 {array} models.ProductResponse
// @Router /products/category/{category} [get]
func (c *ProductController) GetProductsByCategory(ctx *gin.Context) {
	category := ctx.Param("category")

	products, err := c.productService.GetProductsByCategory(category)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch products",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Products retrieved successfully",
		"count":   len(products),
		"data":    products,
	})
}

// SearchProducts searches products
// @Summary Search products
// @Description Search products by name or description
// @Tags products
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} models.ProductResponse
// @Router /products/search [get]
func (c *ProductController) SearchProducts(ctx *gin.Context) {
	query := ctx.Query("q")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"message": "Search query is required",
		})
		return
	}

	products, err := c.productService.SearchProducts(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Search failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Search completed successfully",
		"count":   len(products),
		"data":    products,
	})
}
