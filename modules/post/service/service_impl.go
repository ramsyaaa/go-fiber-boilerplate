package service

import (
	"context"

	"go-fiber-modular/models"
	"go-fiber-modular/modules/post/repository"
)

type service struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &service{repo: repo}
}

func (s *service) FindByID(ctx context.Context, id int64) (*models.Post, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) Create(ctx context.Context, data CreatePostData) (*models.Post, error) {
	post := &models.Post{
		Title:   data.Title,
		Content: data.Content,
	}

	err := s.repo.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *service) Update(ctx context.Context, data UpdatePostData) (*models.Post, error) {
	post, err := s.repo.FindByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, err
	}

	post.Title = data.Title
	post.Content = data.Content

	err = s.repo.Update(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *service) DeleteByID(ctx context.Context, id int64) (*models.Post, error) {
	post, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, err
	}

	err = s.repo.Delete(ctx, post)

	return post, err
}
