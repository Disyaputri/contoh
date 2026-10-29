package domain

import (
	"clean-arch/internal/model/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindAll(ctx context.Context, ) ([]entity.User, error)
	FindByID(ctx context.Context, id int) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int) error
}

type UserUsecase interface {
	Create(name, email string) error
	GetAll() ([]entity.User, error)
	GetByID(id int) (*entity.User, error)
	Update(id int, name, email string) error
	Delete(id int) error
}

type UserHandler interface {
	Run()
}