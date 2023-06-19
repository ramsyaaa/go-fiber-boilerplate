package repository

import (
	"context"

	"go-fiber-modular/models"
)

type PostRepository interface {
	FindByID(ctx context.Context, id int64) (*models.Post, error)
	Create(ctx context.Context, post *models.Post) error
	Update(ctx context.Context, post *models.Post) error
	Delete(ctx context.Context, post *models.Post) error
}
