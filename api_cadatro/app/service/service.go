package service

import (
	"context"
	"fmt"
)

type Request struct {
	Email string
	Name  string
	Phone string
}
type valuesDb struct {
	Email string
	Name  string
	Phone string
}

type creator interface {
	Create(ctx context.Context, r *valuesDb) error
}
type CreatorService struct {
	ct creator
}

func (c *CreatorService) Create(ctx context.Context, r Request) error {
	re := &valuesDb{
		Email: r.Email,
		Name:  r.Name,
		Phone: r.Phone,
	}

	err := c.ct.Create(ctx, re)
	if err != nil {
		return fmt.Errorf("error on creating creator: %v", err)
	}

	return nil
}
