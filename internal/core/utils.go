package core

import "time"

func CalculateAge(dob time.Time) int {
	now := time.Now();
	age := now.Year() - dob.Year();

	// Check if birthday hasn't occurred yet this year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--;
	}

	return age;
}
