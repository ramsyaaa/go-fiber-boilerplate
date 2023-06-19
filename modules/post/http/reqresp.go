package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type PostDetail struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type (
	ShowRequest struct {
		ID int64 `params:"id"`
	}
	ShowResponse PostDetail
)

func (r *ShowRequest) bind(c *fiber.Ctx) error {
	return c.ParamsParser(r)
}

type (
	StoreRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	StoreResponse PostDetail
)

func (r *StoreRequest) bind(c *fiber.Ctx) error {
	return c.BodyParser(r)
}

type (
	UpdateRequest struct {
		ID      int64  `params:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	UpdateResponse PostDetail
)

func (r *UpdateRequest) bind(c *fiber.Ctx) error {
	err := c.ParamsParser(r)
	if err != nil {
		return err
	}
	err = c.BodyParser(r)
	if err != nil {
		return err
	}
	return nil
}

type (
	DeleteRequest struct {
		ID int64 `params:"id"`
	}
	DeleteResponse struct {
		Deleted bool `json:"deleted"`
	}
)

func (r *DeleteRequest) bind(c *fiber.Ctx) error {
	return c.ParamsParser(r)
}
