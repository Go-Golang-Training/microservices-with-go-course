package model

import "errors"

// ErrNotFound is returned when a movie is not found
var ErrNotFound = errors.New("not found")

// Metadata defines the movie metadata domain model
type Metadata struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
}
