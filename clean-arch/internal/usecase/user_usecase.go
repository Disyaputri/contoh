package usecase

import (
	"clean-arch/internal/domain"
	"clean-arch/internal/model/entity"
	"context"
	"time"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

var setTimeout = 3 * time.Second
func (u *userUsecase) Create(name, email string) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()
	user := &entity.User{
		Name: name,
		Email: email,
	}

	return u.repo.Create(ctx, user)
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()
	return u.repo.FindAll(ctx)
}

func (u *userUsecase) GetByID(id int) (*entity.User, error) {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()

	return u.repo.FindByID(ctx, id)
}

func (u *userUsecase) Update(id int, name, email string) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()

	user := &entity.User{
		ID: id,
		Name: name,
		Email: email,
	}

	return u.repo.Update(ctx, user)
}

func (u *userUsecase) Delete(id int) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()

	return u.repo.Delete(ctx, id)
}