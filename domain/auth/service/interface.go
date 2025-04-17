package service

import (
	"context"

	"example-go-api/domain/auth/entity"
	"example-go-api/domain/auth/request"
)

type AuthService interface {
	Create(ctx context.Context, req *request.CreateAuthRequest) (*entity.User, error)
	Get(ctx context.Context, email string) (*entity.User, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
}
