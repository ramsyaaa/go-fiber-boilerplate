package http

import (
	"go-fiber-modular/helper"
	"go-fiber-modular/models"
	"go-fiber-modular/modules/post/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) Show(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &ShowRequest{}
	if err := req.bind(c); err != nil {
		return err
	}

	post, err := h.service.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	// ensure post exists or return 404
	if post == nil {
		c.Status(http.StatusNotFound).JSON(helper.APIResponse("Post is not found", http.StatusNotFound, "Error", []models.Post{}))
		return nil
	}

	resp := &ShowResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Get Post", http.StatusOK, "Success", resp))
	return nil
}

func (h *PostHandler) Store(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &StoreRequest{}
	if err := req.bind(c); err != nil {
		c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to Store Post", http.StatusBadRequest, "Error", err))
		return nil
	}

	post, err := h.service.Create(ctx, service.CreatePostData{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return err
	}

	resp := &StoreResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Store Post", http.StatusOK, "Success", resp))
	return nil
}

func (h *PostHandler) Update(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &UpdateRequest{}
	if err := req.bind(c); err != nil {
		c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to Update Post", http.StatusBadRequest, "Error", err))
		return nil
	}

	post, err := h.service.Update(ctx, service.UpdatePostData{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return err
	}
	// ensure post exists or return 404
	if post == nil {
		c.Status(http.StatusNotFound).JSON(helper.APIResponse("Post is not found", http.StatusNotFound, "Error", []models.Post{}))
		return nil
	}

	resp := &UpdateResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Store Post", http.StatusOK, "Success", resp))
	return nil
}

func (h *PostHandler) Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &DeleteRequest{}
	if err := req.bind(c); err != nil {
		c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to Delete Post", http.StatusBadRequest, "Error", err))
		return nil
	}

	post, err := h.service.DeleteByID(ctx, req.ID)
	if err != nil {
		return err
	}
	// ensure post exists or return 404
	if post == nil {
		c.Status(http.StatusNotFound).JSON(helper.APIResponse("Post is not found", http.StatusNotFound, "Error", []models.Post{}))
		return nil
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Delete Post", http.StatusOK, "Success", []models.Post{}))
	return nil
}
