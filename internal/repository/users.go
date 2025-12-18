package repository

import (
	"context"
	"time"

	db "github.com/Shoyeb45/simple-go-dob-api/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	q *db.Queries
}

// Get the user by Id, you will be returned pointer to the user
func (r *UserRepository) GetByID(ctx context.Context, id int64) (*db.User, error) {
	user, err := r.q.GetUser(ctx, id);
	return &user, err
} 

// Create the user with dob and name
func (r *UserRepository) Create(ctx context.Context, name string, dob *time.Time) (*db.User, error) {
	createdUser, err := r.q.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob: pgtype.Date{
			Time: *dob,
			Valid: true,
		},
	});

	return &createdUser, err;
}

// Update the entire user or update the user partially
func (r *UserRepository) UpdateById(ctx context.Context, name *string, dob *time.Time, id int64) (*db.User, error) {
	params := db.UpdateUserPartialParams{
		ID: id,
	};

	if name != nil {
		params.Name = *name;
	}

	if dob != nil {
		params.Dob = pgtype.Date{
			Time: *dob,
			Valid: true,
		}
	}

	updatedUser, err := r.q.UpdateUserPartial(ctx, params);
	return &updatedUser, err;
} 

// Delete the user by Id
func (r *UserRepository) DeleteById(ctx context.Context, id int64) error {
	return r.q.DeleteUser(ctx, id);
}