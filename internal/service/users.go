package service

import (
	"context"
	"time"

	db "github.com/Shoyeb45/simple-go-dob-api/db/sqlc"
	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
	"github.com/Shoyeb45/simple-go-dob-api/internal/models"
	"github.com/Shoyeb45/simple-go-dob-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *repository.UserRepository
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*db.User, error) {
	return s.repo.GetByID(ctx, id)
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, name string, dob time.Time) (*db.User, error) {
	return s.repo.Create(ctx, name, &pgtype.Date{
		Time: dob, Valid: true,
	})
}

func (h *UserService) GetAllUsers(ctx context.Context) ([]db.User, error) {
	return h.repo.GetAllUsers(ctx)
}

func (h *UserService) DeleteUser(ctx context.Context, id int64) error {
	return h.repo.DeleteById(ctx, id)
}

func (h *UserService) UpdateUser(ctx context.Context, id int64, userDate models.UserCreate) (*db.User, error) {
	// convert the date-string into proper format
	dob, err := time.Parse("2006-01-02", userDate.Dob)

	if err != nil {
		return nil, core.NewBadRequestError("Invalid date of birth format.").WithDetails("dob", userDate.Dob).WithInternal(err)
	}

	return h.repo.UpdateById(ctx, &userDate.Name, &dob, id)
}
