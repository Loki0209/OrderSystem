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

type ProductService struct {
	collection *mongo.Collection
}

// NewProductService creates a new product service instance
func NewProductService() *ProductService {
	return &ProductService{
		collection: config.GetCollection("products"),
	}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(req models.CreateProductRequest, userID string) (*models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if SKU already exists
	var existingProduct models.Product
	err := s.collection.FindOne(ctx, bson.M{"sku": req.SKU}).Decode(&existingProduct)
	if err == nil {
		return nil, errors.New("product with this SKU already exists")
	}

	// Create product
	product := models.Product{
		ID:          primitive.NewObjectID(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Category:    req.Category,
		SKU:         req.SKU,
		IsActive:    true,
		CreatedBy:   userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = s.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, errors.New("failed to create product")
	}

	response := product.ToProductResponse()
	return &response, nil
}

// GetProductByID retrieves a product by ID
func (s *ProductService) GetProductByID(productID string) (*models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product models.Product
	err = s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		return nil, errors.New("product not found")
	}

	response := product.ToProductResponse()
	return &response, nil
}

// GetAllProducts retrieves all products with optional filtering
func (s *ProductService) GetAllProducts(category string, isActive *bool) ([]models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build filter
	filter := bson.M{}
	if category != "" {
		filter["category"] = category
	}
	if isActive != nil {
		filter["is_active"] = *isActive
	}

	// Set sort options (newest first)
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := s.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.New("failed to fetch products")
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, errors.New("failed to decode products")
	}

	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, product.ToProductResponse())
	}

	return responses, nil
}

// UpdateProduct updates product information (PUT - full update)
func (s *ProductService) UpdateProduct(productID string, req models.UpdateProductRequest) (*models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	// Build update document
	update := bson.M{
		"updated_at": time.Now(),
	}

	if req.Name != "" {
		update["name"] = req.Name
	}
	if req.Description != "" {
		update["description"] = req.Description
	}
	if req.Price != nil {
		update["price"] = *req.Price
	}
	if req.Quantity != nil {
		update["quantity"] = *req.Quantity
	}
	if req.Category != "" {
		update["category"] = req.Category
	}
	if req.SKU != "" {
		// Check if SKU is already used by another product
		var existingProduct models.Product
		err := s.collection.FindOne(ctx, bson.M{
			"sku": req.SKU,
			"_id": bson.M{"$ne": objectID},
		}).Decode(&existingProduct)
		if err == nil {
			return nil, errors.New("SKU already used by another product")
		}
		update["sku"] = req.SKU
	}
	if req.IsActive != nil {
		update["is_active"] = *req.IsActive
	}

	// Update product
	result := s.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if result.Err() != nil {
		return nil, errors.New("product not found")
	}

	var updatedProduct models.Product
	if err := result.Decode(&updatedProduct); err != nil {
		return nil, errors.New("failed to decode updated product")
	}

	response := updatedProduct.ToProductResponse()
	return &response, nil
}

// PatchProduct partially updates product information (PATCH - partial update)
func (s *ProductService) PatchProduct(productID string, updates map[string]interface{}) (*models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	// Add updated_at to updates
	updates["updated_at"] = time.Now()

	// Update product
	result := s.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updates},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if result.Err() != nil {
		return nil, errors.New("product not found")
	}

	var updatedProduct models.Product
	if err := result.Decode(&updatedProduct); err != nil {
		return nil, errors.New("failed to decode updated product")
	}

	response := updatedProduct.ToProductResponse()
	return &response, nil
}

// DeleteProduct deletes a product by ID
func (s *ProductService) DeleteProduct(productID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return errors.New("invalid product ID")
	}

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return errors.New("failed to delete product")
	}

	if result.DeletedCount == 0 {
		return errors.New("product not found")
	}

	return nil
}

// UpdateProductQuantity updates only the quantity (useful for inventory management)
func (s *ProductService) UpdateProductQuantity(productID string, quantity int) (*models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	update := bson.M{
		"quantity":   quantity,
		"updated_at": time.Now(),
	}

	result := s.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if result.Err() != nil {
		return nil, errors.New("product not found")
	}

	var updatedProduct models.Product
	if err := result.Decode(&updatedProduct); err != nil {
		return nil, errors.New("failed to decode updated product")
	}

	response := updatedProduct.ToProductResponse()
	return &response, nil
}

// GetProductsByCategory retrieves products by category
func (s *ProductService) GetProductsByCategory(category string) ([]models.ProductResponse, error) {
	return s.GetAllProducts(category, nil)
}

// SearchProducts searches products by name or description
func (s *ProductService) SearchProducts(query string) ([]models.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Case-insensitive search in name and description
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("failed to search products")
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, errors.New("failed to decode products")
	}

	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, product.ToProductResponse())
	}

	return responses, nil
}
