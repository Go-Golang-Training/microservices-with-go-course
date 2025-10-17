package memory

import (
	"context"
	"sync"

	"movieexample.com/metadata/pkg/model"
)

// Repository implements the repository pattern for movie metadata
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New is a constructor function - idiomatic Go naming
func New() *Repository {
	return &Repository{
		data: make(map[string]*model.Metadata),
	}
}

// Get demonstrates effective error handling in Go
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, model.ErrNotFound
	}
	return m, nil
}

// Put shows mutex usage for safe concurrent access
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata
	return nil
}
