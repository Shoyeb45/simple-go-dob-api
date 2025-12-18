package models

type UserCreate struct {
	Name string `json:"name" validate:"required" min="1" max="50"`
	Dob string `json:"dob"`
}

type UserResponse struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    Dob  string `json:"dob"`
}

type UserWithAgeResponse struct {
    UserResponse
    Age int `json:"age"`
}
