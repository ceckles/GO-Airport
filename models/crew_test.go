package models

import "testing"

func TestNewCrew(t *testing.T) {
	crew := NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")

	if crew == nil {
		t.Error("NewCrew returned nil")
	}

	if crew.GetName() != "Bob Smith" {
		t.Errorf("Expected name 'Bob Smith', got '%s'", crew.GetName())
	}

	if crew.GetAge() != 28 {
		t.Errorf("Expected age 28, got %d", crew.GetAge())
	}

	if crew.GetExperience() != 8 {
		t.Errorf("Expected experience 8, got %d", crew.GetExperience())
	}

	if crew.GetRating() != 8.0 {
		t.Errorf("Expected rating 8.0, got %.2f", crew.GetRating())
	}

	if crew.GetLicense() != "CREW001" {
		t.Errorf("Expected license 'CREW001', got '%s'", crew.GetLicense())
	}

	if crew.GetStatus() != "Active" {
		t.Errorf("Expected status 'Active', got '%s'", crew.GetStatus())
	}
}

func TestCrew_GettersAndSetters(t *testing.T) {
	crew := NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")

	// Test setters
	crew.SetName("Alice Johnson")
	if crew.GetName() != "Alice Johnson" {
		t.Errorf("SetName failed, expected 'Alice Johnson', got '%s'", crew.GetName())
	}

	crew.SetAge(32)
	if crew.GetAge() != 32 {
		t.Errorf("SetAge failed, expected 32, got %d", crew.GetAge())
	}

	crew.SetExperience(12)
	if crew.GetExperience() != 12 {
		t.Errorf("SetExperience failed, expected 12, got %d", crew.GetExperience())
	}

	crew.SetRating(9.0)
	if crew.GetRating() != 9.0 {
		t.Errorf("SetRating failed, expected 9.0, got %.2f", crew.GetRating())
	}

	crew.SetLicense("CREW002")
	if crew.GetLicense() != "CREW002" {
		t.Errorf("SetLicense failed, expected 'CREW002', got '%s'", crew.GetLicense())
	}

	crew.SetStatus("Inactive")
	if crew.GetStatus() != "Inactive" {
		t.Errorf("SetStatus failed, expected 'Inactive', got '%s'", crew.GetStatus())
	}
}

func TestCrew_GetCrew(t *testing.T) {
	crew := NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")

	returnedCrew := crew.GetCrew()
	if returnedCrew != crew {
		t.Error("GetCrew should return the same crew instance")
	}
}
