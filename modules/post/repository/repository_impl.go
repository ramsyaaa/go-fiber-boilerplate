package repository

import (
	"context"
	"errors"

	"go-fiber-modular/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &repository{db: db}
}

func (r *repository) FindByID(ctx context.Context, id int64) (*models.Post, error) {
	var post *models.Post
	err := r.db.WithContext(ctx).First(&post, id).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return post, err
}

func (r *repository) Create(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Create(post).Error
}

func (r *repository) Update(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Save(post).Error
}

func (r *repository) Delete(ctx context.Context, post *models.Post) error {
	return r.db.WithContext(ctx).Delete(post).Error
}
