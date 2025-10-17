package model

import "movieexample.com/metadata/pkg/model"

// MovieDetails is an aggregate root that combines multiple domain entities
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
