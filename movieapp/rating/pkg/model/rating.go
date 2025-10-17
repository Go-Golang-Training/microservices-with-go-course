package model

// RecordType defines the type of record being rated
// This demonstrates Domain-Driven Design with explicit types
type RecordType string

const (
	RecordTypeMovie RecordType = "movie"
)

// UserID is a domain primitive - making user IDs explicit in the type system
type UserID string

// RecordID identifies records across the system
type RecordID string

// RatingValue represents a rating score with validation boundaries
type RatingValue int

// Rating is our core domain entity for user ratings
type Rating struct {
	RecordID   RecordID    `json:"recordId"`
	RecordType RecordType  `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
