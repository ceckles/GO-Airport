package models

import "testing"

func TestNewPilot(t *testing.T) {
	pilot := NewPilot("John Doe", 30, 10, 9.5, "LIC123", "Active")

	if pilot == nil {
		t.Error("NewPilot returned nil")
	}

	if pilot.GetName() != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", pilot.GetName())
	}

	if pilot.GetAge() != 30 {
		t.Errorf("Expected age 30, got %d", pilot.GetAge())
	}

	if pilot.GetExperience() != 10 {
		t.Errorf("Expected experience 10, got %d", pilot.GetExperience())
	}

	if pilot.GetRating() != 9.5 {
		t.Errorf("Expected rating 9.5, got %.2f", pilot.GetRating())
	}

	if pilot.GetLicense() != "LIC123" {
		t.Errorf("Expected license 'LIC123', got '%s'", pilot.GetLicense())
	}

	if pilot.GetStatus() != "Active" {
		t.Errorf("Expected status 'Active', got '%s'", pilot.GetStatus())
	}
}

func TestPilot_GettersAndSetters(t *testing.T) {
	pilot := NewPilot("John Doe", 30, 10, 9.5, "LIC123", "Active")

	// Test setters
	pilot.SetName("Jane Doe")
	if pilot.GetName() != "Jane Doe" {
		t.Errorf("SetName failed, expected 'Jane Doe', got '%s'", pilot.GetName())
	}

	pilot.SetAge(35)
	if pilot.GetAge() != 35 {
		t.Errorf("SetAge failed, expected 35, got %d", pilot.GetAge())
	}

	pilot.SetExperience(15)
	if pilot.GetExperience() != 15 {
		t.Errorf("SetExperience failed, expected 15, got %d", pilot.GetExperience())
	}

	pilot.SetRating(9.8)
	if pilot.GetRating() != 9.8 {
		t.Errorf("SetRating failed, expected 9.8, got %.2f", pilot.GetRating())
	}

	pilot.SetLicense("LIC456")
	if pilot.GetLicense() != "LIC456" {
		t.Errorf("SetLicense failed, expected 'LIC456', got '%s'", pilot.GetLicense())
	}

	pilot.SetStatus("Inactive")
	if pilot.GetStatus() != "Inactive" {
		t.Errorf("SetStatus failed, expected 'Inactive', got '%s'", pilot.GetStatus())
	}
}

func TestPilot_GetPilot(t *testing.T) {
	pilot := NewPilot("John Doe", 30, 10, 9.5, "LIC123", "Active")

	returnedPilot := pilot.GetPilot()
	if returnedPilot != pilot {
		t.Error("GetPilot should return the same pilot instance")
	}
}
