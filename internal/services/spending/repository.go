package spending

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
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

// FindAll retrieves all spending records from MongoDB
func (r *Repository) FindAll(ctx context.Context) ([]Spending, error) {
	collection := r.DB.Collection("spendings")

	cursor, err := collection.Find(ctx, bson.M{}) // empty filter = all records
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var spendings []Spending
	for cursor.Next(ctx) {
		var s Spending
		if err := cursor.Decode(&s); err != nil {
			return nil, err
		}
		spendings = append(spendings, s)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return spendings, nil
}

// FindByType retrieves all spending records with/without the given category/type
func (r *Repository) SumByCategory(ctx context.Context, category string) (float64, error) {
	collection := r.DB.Collection("spendings")

	pipeline := []bson.M{}

	if category != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"type": category}})
	}

	pipeline = append(pipeline, bson.M{
		"$group": bson.M{
			"_id":   nil,
			"total": bson.M{"$sum": "$amount"},
		},
	})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	// decode directly into a struct
	var res struct {
		Total float64 `bson:"total"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&res); err != nil {
			return 0, err
		}
		return res.Total, nil
	}

	// no records found
	return 0, nil
}
