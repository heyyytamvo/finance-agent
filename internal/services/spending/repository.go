package spending

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository handles MongoDB operations for Spending
type Repository struct {
	DB *mongo.Database
}

// NewRepository creates a new spending repository
func NewRepository(db *mongo.Database) *Repository {
	return &Repository{DB: db}
}

// Insert adds a new Spending record to MongoDB
func (r *Repository) Insert(ctx context.Context, s Spending) (*Spending, error) {
	now := time.Now()
	s.CreatedAt = now
	s.UpdatedAt = now
	s.ID = primitive.NewObjectID().Hex()

	collection := r.DB.Collection("spendings")
	_, err := collection.InsertOne(ctx, s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
