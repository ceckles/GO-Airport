package models

import "testing"

func TestNewPlane(t *testing.T) {
	pilot1 := NewPilot("John Doe", 30, 10, 9.5, "LIC001", "Active")
	crew1 := NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")
	passenger1 := NewPassenger("Jane Doe", 42, "12A")
	
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White", 100000, 200000, 150000, 100000, 80000,
		[]*Crew{crew1}, []*Pilot{pilot1}, 100,
		[]*Passenger{passenger1}, "London", "New York", "Landed", 50000,
	)
	
	if plane == nil {
		t.Error("NewPlane returned nil")
	}
	
	if plane.GetID() != "PLANE001" {
		t.Errorf("Expected ID 'PLANE001', got '%s'", plane.GetID())
	}
	
	if plane.GetModel() != "Boeing 747" {
		t.Errorf("Expected model 'Boeing 747', got '%s'", plane.GetModel())
	}
	
	if len(plane.GetPilots()) != 1 {
		t.Errorf("Expected 1 pilot, got %d", len(plane.GetPilots()))
	}
	
	if len(plane.GetCrew()) != 1 {
		t.Errorf("Expected 1 crew member, got %d", len(plane.GetCrew()))
	}
	
	if len(plane.GetPassengers()) != 1 {
		t.Errorf("Expected 1 passenger, got %d", len(plane.GetPassengers()))
	}
}

func TestPlane_GettersAndSetters(t *testing.T) {
	pilot1 := NewPilot("John Doe", 30, 10, 9.5, "LIC001", "Active")
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White", 100000, 200000, 150000, 100000, 80000,
		[]*Crew{}, []*Pilot{pilot1}, 100, []*Passenger{},
		"London", "New York", "Landed", 50000,
	)
	
	// Test setters
	plane.SetID("PLANE002")
	if plane.GetID() != "PLANE002" {
		t.Errorf("SetID failed")
	}
	
	plane.SetModel("Airbus A380")
	if plane.GetModel() != "Airbus A380" {
		t.Errorf("SetModel failed")
	}
	
	plane.SetYear(2025)
	if plane.GetYear() != 2025 {
		t.Errorf("SetYear failed")
	}
	
	plane.SetPrice(2000000)
	if plane.GetPrice() != 2000000 {
		t.Errorf("SetPrice failed")
	}
	
	plane.SetDestination("Paris")
	if plane.GetDestination() != "Paris" {
		t.Errorf("SetDestination failed")
	}
	
	plane.SetOrigin("Tokyo")
	if plane.GetOrigin() != "Tokyo" {
		t.Errorf("SetOrigin failed")
	}
	
	plane.SetStatus("In Flight")
	if plane.GetStatus() != "In Flight" {
		t.Errorf("SetStatus failed")
	}
}

func TestPlane_AddPassenger(t *testing.T) {
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White", 100000, 200000, 150000, 100000, 80000,
		[]*Crew{}, []*Pilot{}, 100, []*Passenger{},
		"London", "New York", "Landed", 50000,
	)
	
	passenger1 := NewPassenger("John Doe", 45, "12A")
	passenger2 := NewPassenger("Jane Doe", 42, "12B")
	
	plane.AddPassenger(passenger1)
	if len(plane.GetPassengers()) != 1 {
		t.Errorf("Expected 1 passenger after AddPassenger, got %d", len(plane.GetPassengers()))
	}
	
	plane.AddPassenger(passenger2)
	if len(plane.GetPassengers()) != 2 {
		t.Errorf("Expected 2 passengers after second AddPassenger, got %d", len(plane.GetPassengers()))
	}
}

func TestPlane_RemovePassenger(t *testing.T) {
	passenger1 := NewPassenger("John Doe", 45, "12A")
	passenger2 := NewPassenger("Jane Doe", 42, "12B")
	passenger3 := NewPassenger("Bob Smith", 35, "12C")
	
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White", 100000, 200000, 150000, 100000, 80000,
		[]*Crew{}, []*Pilot{}, 100, []*Passenger{passenger1, passenger2, passenger3},
		"London", "New York", "Landed", 50000,
	)
	
	plane.RemovePassenger(passenger2)
	
	passengers := plane.GetPassengers()
	if len(passengers) != 2 {
		t.Errorf("Expected 2 passengers after removal, got %d", len(passengers))
	}
	
	// Verify passenger2 is removed
	for _, p := range passengers {
		if p == passenger2 {
			t.Error("passenger2 should have been removed")
		}
	}
}

func TestPlane_AddCrew(t *testing.T) {
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White", 100000, 200000, 150000, 100000, 80000,
		[]*Crew{}, []*Pilot{}, 100, []*Passenger{},
		"London", "New York", "Landed", 50000,
	)
	
	crew1 := NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")
	crew2 := NewCrew("Alice Johnson", 32, 12, 9.0, "CREW002", "Active")
	
	plane.AddCrew(crew1)
	if len(plane.GetCrew()) != 1 {
		t.Errorf("Expected 1 crew member after AddCrew, got %d", len(plane.GetCrew()))
	}
	
	plane.AddCrew(crew2)
	if len(plane.GetCrew()) != 2 {
		t.Errorf("Expected 2 crew members after second AddCrew, got %d", len(plane.GetCrew()))
	}
}

func TestPlane_RemoveCrew(t *testing.T) {
	crew1 := NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")
	crew2 := NewCrew("Alice Johnson", 32, 12, 9.0, "CREW002", "Active")
	crew3 := NewCrew("Charlie Brown", 26, 6, 7.5, "CREW003", "Active")
	
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White", 100000, 200000, 150000, 100000, 80000,
		[]*Crew{crew1, crew2, crew3}, []*Pilot{}, 100, []*Passenger{},
		"London", "New York", "Landed", 50000,
	)
	
	plane.RemoveCrew(crew2)
	
	crew := plane.GetCrew()
	if len(crew) != 2 {
		t.Errorf("Expected 2 crew members after removal, got %d", len(crew))
	}
	
	// Verify crew2 is removed
	for _, c := range crew {
		if c == crew2 {
			t.Error("crew2 should have been removed")
		}
	}
}

func TestPlane_GetCalculatedWeight(t *testing.T) {
	pilot1 := NewPilot("John Doe", 30, 10, 9.5, "LIC001", "Active")
	
	passenger1 := NewPassenger("John Doe", 45, "12A")
	passenger1.SetPassangerWeight(80.0)
	passenger1.SetBaggageWeight(25.5)
	
	passenger2 := NewPassenger("Jane Doe", 42, "12B")
	passenger2.SetPassangerWeight(65.0)
	passenger2.SetBaggageWeight(18.0)
	
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White",
		100000, // Weight
		200000, // MaxWeight
		150000, // CurrentWeight
		100000, // FuelCapacity
		80000,  // FuelLevel
		[]*Crew{}, []*Pilot{pilot1}, 100,
		[]*Passenger{passenger1, passenger2},
		"London", "New York", "Landed",
		50000, // CargoWeight
	)
	
	// Expected: 100000 (Weight) + 150000 (CurrentWeight) + 50000 (CargoWeight) + 80.0 + 25.5 + 65.0 + 18.0
	expectedWeight := 100000.0 + 150000.0 + 50000.0 + 80.0 + 25.5 + 65.0 + 18.0
	calculatedWeight := plane.GetCalculatedWeight()
	
	if calculatedWeight != expectedWeight {
		t.Errorf("Expected calculated weight %.2f, got %.2f", expectedWeight, calculatedWeight)
	}
	
	// Test with no passengers
	plane2 := NewPlane(
		"PLANE002", "Cargo", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White",
		100000, 200000, 150000, 100000, 80000,
		[]*Crew{}, []*Pilot{}, 0, []*Passenger{},
		"London", "New York", "Landed", 50000,
	)
	
	expectedWeight2 := 100000.0 + 150000.0 + 50000.0
	calculatedWeight2 := plane2.GetCalculatedWeight()
	
	if calculatedWeight2 != expectedWeight2 {
		t.Errorf("Expected calculated weight %.2f for plane with no passengers, got %.2f", expectedWeight2, calculatedWeight2)
	}
}

func TestPlane_WeightFields(t *testing.T) {
	plane := NewPlane(
		"PLANE001", "Passenger", "Boeing 747", 2024, "Boeing", "Owner",
		1000000, "White", 100000, 200000, 150000, 100000, 80000,
		[]*Crew{}, []*Pilot{}, 100, []*Passenger{},
		"London", "New York", "Landed", 50000,
	)
	
	plane.SetWeight(110000)
	if plane.GetWeight() != 110000 {
		t.Errorf("SetWeight/GetWeight failed")
	}
	
	plane.SetMaxWeight(250000)
	if plane.GetMaxWeight() != 250000 {
		t.Errorf("SetMaxWeight/GetMaxWeight failed")
	}
	
	plane.SetCurrentWeight(160000)
	if plane.GetCurrentWeight() != 160000 {
		t.Errorf("SetCurrentWeight/GetCurrentWeight failed")
	}
	
	plane.SetCargoWeight(60000)
	if plane.GetCargoWeight() != 60000 {
		t.Errorf("SetCargoWeight/GetCargoWeight failed")
	}
}

