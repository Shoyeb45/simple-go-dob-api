package mapper

import (
	db "github.com/Shoyeb45/simple-go-dob-api/db/sqlc"
	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
	"github.com/Shoyeb45/simple-go-dob-api/internal/models"
)

func UserToResponse(u *db.User) *models.UserResponse {
	var dob string
	if u.Dob.Valid {
		dob = u.Dob.Time.Format("2006-01-02")
	}

	return &models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		Dob:  dob,
	}
}

func UsersToResponse(users []db.User) []models.UserResponse {
	res := make([]models.UserResponse, 0, len(users))

	for i := range users {
		res = append(res, *UserToResponse(&users[i]))
	}

	return res
}

func UserToWithAgeResponse(u *db.User) *models.UserWithAgeResponse {
    base := UserToResponse(u)

    age := core.CalculateAge(u.Dob.Time)

    return &models.UserWithAgeResponse{
        UserResponse: *base,
        Age:          age,
    }
}
