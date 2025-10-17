package metadata

import (
	"context"
	"errors"

	"movieexample.com/metadata/pkg/model"
)

// ErrNotFound is a domain-specific error
var ErrNotFound = errors.New("not found")

// metadataRepository interface demonstrates dependency inversion principle
type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller contains business logic - clean architecture separation
type Controller struct {
	repo metadataRepository
}

// New is a constructor that accepts dependencies - Dependency Injection
func New(repo metadataRepository) *Controller {
	return &Controller{repo: repo}
}

// Get demonstrates context propagation and proper error wrapping
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
