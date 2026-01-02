package spending

import "time"

// Spending represents a user's spending record
type Spending struct {
	ID            string    `bson:"_id,omitempty" json:"id,omitempty"`
	Date          time.Time `bson:"date" json:"date"`                     // Date of spending
	Amount        float64   `bson:"amount" json:"amount"`                 // Amount spent
	Currency      string    `bson:"currency" json:"currency"`             // Currency code (USD, EUR, etc.)
	Type          string    `bson:"type" json:"type"`                     // Spending type/category
	Description   string    `bson:"description,omitempty" json:"description,omitempty"`
	PaymentMethod string    `bson:"payment_method,omitempty" json:"payment_method,omitempty"`
	Tags          []string  `bson:"tags,omitempty" json:"tags,omitempty"`
	Recurring     bool      `bson:"recurring,omitempty" json:"recurring,omitempty"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
}
