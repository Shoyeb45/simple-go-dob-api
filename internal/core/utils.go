package core

import (
	"strconv"
	"time"
)

// Function to calculate age from the current time, 
// Note that the current time passed as parameter to make it strong testable.
// If some person complete 21 years and he/she is in 22nd year then this will output only 21 year.
// It outputs only completed years.
func CalculateAge(dob time.Time, now *time.Time) int {
	age := now.Year() - dob.Year();

	// Check if birthday hasn't occurred yet this year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--;
	}

	return age;
}



func ParseDob(dob string) (*time.Time, error) {
	parsedDob, err := time.Parse("2006-01-02", dob);

	if err != nil {
		return nil, err;
	}
	return &parsedDob, nil;
}


func ConvertIdToi64(id *string) (int64, error) {
	idInNum, err := strconv.ParseInt(*id, 10, 64);

	if err != nil {
		return 0, NewBadRequestError("Not valid id given").WithInternal(err).WithDetails("id", id);
	}

	return idInNum, nil;
}