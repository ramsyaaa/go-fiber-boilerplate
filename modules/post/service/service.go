package service

import (
	"context"

	"go-fiber-modular/models"
)

type PostService interface {
	FindByID(ctx context.Context, id int64) (*models.Post, error)
	Create(ctx context.Context, data CreatePostData) (*models.Post, error)
	Update(ctx context.Context, data UpdatePostData) (*models.Post, error)
	DeleteByID(ctx context.Context, id int64) (*models.Post, error)
}

type CreatePostData struct {
	Title   string
	Content string
}

type UpdatePostData struct {
	ID      int64
	Title   string
	Content string
}
