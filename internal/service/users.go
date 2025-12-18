package service

import (
	"context"
	"time"

	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
	"github.com/Shoyeb45/simple-go-dob-api/internal/mapper"
	"github.com/Shoyeb45/simple-go-dob-api/internal/models"
	"github.com/Shoyeb45/simple-go-dob-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	repo *repository.UserRepository
}

// Get user if the user exist with given Id.
func (s *UserService) GetUser(ctx context.Context, idStr *string) (*models.UserWithAgeResponse, error) {
	// convert the id to the int64
	id, err := core.ConvertIdToi64(idStr);

	if err != nil {
		return nil, err;
	}

	user, err := s.repo.GetByID(ctx, id);

	// Handle not found user
	if err != nil {
		return nil, core.NewNotFoundError("No user found with given ID.").WithInternal(err).WithDetails("id", id);
	}

	return mapper.UserToWithAgeResponse(user), nil;
}

// validate and create the user
func (s *UserService) CreateUser(ctx context.Context, body *models.UserCreate) (*models.UserResponse, error) {
	// parse the date of birth
	dob, err := core.ParseDob(body.Dob);

	if err != nil {
		return nil, core.NewBadRequestError("Invalid Date of Birth Provided.").WithInternal(err);
	}

	user, err := s.repo.Create(ctx, body.Name, &pgtype.Date{
		Time: *dob, Valid: true,
	});

	if err != nil {
		return nil, core.NewInternalError("Failed to create user.").WithInternal(err);
	}
	return mapper.UserToResponse(user), nil;
}

// Get all the users present in the db
func (h *UserService) GetAllUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := h.repo.GetAllUsers(ctx);

	if err != nil {
		return nil, core.NewInternalError("Failed to retrieve users.").WithInternal(err);
	}

	return mapper.UsersToResponse(users), nil
}

// Verify and Delete the user
func (h *UserService) DeleteUser(ctx context.Context, idStr *string) error {
	id, err := core.ConvertIdToi64(idStr);

	if err != nil {
		return err;
	}

	// check if the user is present or not
	_, err = h.repo.GetByID(ctx, id);

	if err != nil {
		return core.NewNotFoundError("No User found with given Id.").WithDetails("id", id).WithInternal(err);
	}

	// delte the user from the database
	err = h.repo.DeleteById(ctx, id);

	if err != nil {
		return core.NewInternalError("Failed to delete the user.").WithInternal(err);
	}

	return nil;
}

// Service function to update the user completely
func (h *UserService) UpdateUser(ctx context.Context, idStr *string, userData *models.UserCreate) (*models.UserResponse, error) {
	// convert id to int64
	id, err := core.ConvertIdToi64(idStr);

	if err != nil {
		return nil, err;
	}

	// Check if the user is present or not, if not present then output 404 error
	_, err = h.repo.GetByID(ctx, id);

	if err != nil {
		return nil, core.NewNotFoundError("No user is present with given ID.").WithDetails("id", id).WithInternal(err);
	}

	// convert the date-string into proper format
	dob, err := time.Parse("2006-01-02", userData.Dob)

	if err != nil {
		return nil, core.NewBadRequestError("Invalid date of birth format.").WithDetails("dob", userData.Dob).WithInternal(err)
	}

	updatedUser, err := h.repo.UpdateById(ctx, &userData.Name, &dob, id);

	if err != nil {
		return nil, core.NewInternalError("Failed to update the user.").WithInternal(err);
	}

	return mapper.UserToResponse(updatedUser), nil;
}


// Create new User service.	
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}