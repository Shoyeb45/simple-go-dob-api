package tests

import (
	"testing"
	"time"
	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
)

func TestSimpleDob(t *testing.T) {
	t.Run("Testing simple case", func(t *testing.T) {
		// the date of birth is 1st january, 2000
		dob := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
		
		// We'll try to calculate age till 1st Jan, 2015. It will be 15 years 
		tillDate := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
		expectedAge := 15
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})
}

func TestAgeCalculationEdgeCases(t *testing.T) {
	t.Run("Birthday hasn't occurred this year yet", func(t *testing.T) {
		// Born on Dec 31, 2000
		dob := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)
		// Calculate age on Jan 1, 2024
		tillDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
		expectedAge := 23 // Still 23, not 24 yet
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})

	t.Run("Exact birthday - should include completed year", func(t *testing.T) {
		dob := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
		tillDate := time.Date(2020, 5, 15, 0, 0, 0, 0, time.UTC)
		expectedAge := 30
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})

	t.Run("One day before birthday", func(t *testing.T) {
		dob := time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)
		tillDate := time.Date(2020, 5, 14, 0, 0, 0, 0, time.UTC)
		expectedAge := 29 // Still 29, not 30 yet
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})
}

func TestLeapYearScenarios(t *testing.T) {
	t.Run("Born on leap day - check on leap year", func(t *testing.T) {
		// Born on Feb 29, 2000
		dob := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
		// Check on Feb 29, 2024 (leap year)
		tillDate := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
		expectedAge := 24
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})

	t.Run("Born on leap day - check on non-leap year Feb 28", func(t *testing.T) {
		// Born on Feb 29, 2000
		dob := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
		// Check on Feb 28, 2023 (non-leap year)
		tillDate := time.Date(2023, 2, 28, 0, 0, 0, 0, time.UTC)
		expectedAge := 22 // Birthday hasn't occurred yet
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})

	t.Run("Born on leap day - check on non-leap year March 1", func(t *testing.T) {
		// Born on Feb 29, 2000
		dob := time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC)
		// Check on March 1, 2023 (non-leap year)
		tillDate := time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)
		expectedAge := 23 // Birthday considered passed
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})
}

func TestMonthEndScenarios(t *testing.T) {
	t.Run("Born on Jan 31 - check on Feb 28", func(t *testing.T) {
		dob := time.Date(2000, 1, 31, 0, 0, 0, 0, time.UTC)
		tillDate := time.Date(2020, 2, 28, 0, 0, 0, 0, time.UTC)
		expectedAge := 20 // Birthday passed
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})

	t.Run("Born on March 31 - check on April 30", func(t *testing.T) {
		dob := time.Date(2000, 3, 31, 0, 0, 0, 0, time.UTC)
		tillDate := time.Date(2020, 4, 30, 0, 0, 0, 0, time.UTC)
		expectedAge := 20
		
		calculatedAge := core.CalculateAge(dob, &tillDate)
		if expectedAge != calculatedAge {
			t.Errorf("expected %v, got %v", expectedAge, calculatedAge)
		}
	})
}
