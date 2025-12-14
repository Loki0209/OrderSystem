package services

import (
	"context"
	"errors"
	"time"

	"ordernew/config"
	"ordernew/models"
	"ordernew/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

// NewUserService creates a new user service instance
func NewUserService() *UserService {
	return &UserService{
		collection: config.GetCollection("users"),
	}
}

// Register creates a new user account
func (s *UserService) Register(req models.RegisterRequest) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if user already exists
	var existingUser models.User
	err := s.collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := models.User{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      "user",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	response := user.ToUserResponse()
	return &response, nil
}

// Login authenticates a user and returns a token
func (s *UserService) Login(req models.LoginRequest) (string, *models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find user by email
	var user models.User
	err := s.collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return "", nil, errors.New("user account is inactive")
	}

	// Generate token
	token, err := utils.GenerateToken(user.ID.Hex(), user.Email, user.Role)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	response := user.ToUserResponse()
	return token, &response, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(userID string) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	var user models.User
	err = s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	response := user.ToUserResponse()
	return &response, nil
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("failed to fetch users")
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, errors.New("failed to decode users")
	}

	var responses []models.UserResponse
	for _, user := range users {
		responses = append(responses, user.ToUserResponse())
	}

	return responses, nil
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(userID string, req models.UpdateUserRequest) (*models.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	// Build update document
	update := bson.M{
		"updated_at": time.Now(),
	}

	if req.Name != "" {
		update["name"] = req.Name
	}
	if req.Email != "" {
		update["email"] = req.Email
	}
	if req.Role != "" {
		update["role"] = req.Role
	}
	if req.IsActive != nil {
		update["is_active"] = *req.IsActive
	}

	// Update user
	result := s.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
	)

	if result.Err() != nil {
		return nil, errors.New("user not found")
	}

	// Get updated user
	return s.GetUserByID(userID)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return errors.New("failed to delete user")
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}
