package spending

import (
	"context"
	"errors"
)

// Service handles business logic for Spending
type Service struct {
	Repo *Repository
}

// Create validates and inserts a new spending record
func (s *Service) Create(ctx context.Context, spending Spending) (*Spending, error) {
	if spending.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	if spending.Type == "" {
		return nil, errors.New("type is required")
	}

	return s.Repo.Insert(ctx, spending)
}

// GetAll returns all spending records
func (s *Service) GetAll(ctx context.Context) ([]Spending, error) {
	return s.Repo.FindAll(ctx)
}


// GetByCategory returns all spending records for a category/type
func (s *Service) GetByCategory(ctx context.Context, category string) ([]Spending, error) {
	return s.Repo.FindByType(ctx, category)
}