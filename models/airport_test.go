package models

import "testing"

func TestNewAirport(t *testing.T) {
	// Clear airports slice for clean test
	airports = []*Airport{}

	airport := NewAirport("Test Airport")

	if airport == nil {
		t.Error("NewAirport returned nil")
	}

	if airport.GetName() != "Test Airport" {
		t.Errorf("Expected name 'Test Airport', got '%s'", airport.GetName())
	}

	if len(airport.GetPlaneList()) != 0 {
		t.Errorf("Expected empty plane list, got %d planes", len(airport.GetPlaneList()))
	}

	// Check if airport was added to airports slice
	if len(GetAirports()) != 1 {
		t.Errorf("Expected 1 airport in global list, got %d", len(GetAirports()))
	}
}

func TestAirport_AddPlane(t *testing.T) {
	airports = []*Airport{}
	airport := NewAirport("Test Airport")

	airport.AddPlane("PLANE001")

	planes := airport.GetPlaneList()
	if len(planes) != 1 {
		t.Errorf("Expected 1 plane, got %d", len(planes))
	}

	if planes[0] != "PLANE001" {
		t.Errorf("Expected 'PLANE001', got '%s'", planes[0])
	}

	// Test duplicate prevention
	airport.AddPlane("PLANE001")
	if len(airport.GetPlaneList()) != 1 {
		t.Errorf("Expected 1 plane after duplicate add, got %d", len(airport.GetPlaneList()))
	}

	// Add another plane
	airport.AddPlane("PLANE002")
	if len(airport.GetPlaneList()) != 2 {
		t.Errorf("Expected 2 planes, got %d", len(airport.GetPlaneList()))
	}
}

func TestAirport_Land(t *testing.T) {
	airports = []*Airport{}
	airport := NewAirport("Test Airport")

	airport.Land("PLANE001")

	planes := airport.GetPlaneList()
	if len(planes) != 1 || planes[0] != "PLANE001" {
		t.Errorf("Land failed to add plane")
	}

	// Test duplicate prevention
	airport.Land("PLANE001")
	if len(airport.GetPlaneList()) != 1 {
		t.Errorf("Land should not add duplicate plane")
	}
}

func TestAirport_Takeoff(t *testing.T) {
	airports = []*Airport{}
	airport := NewAirport("Test Airport")

	airport.AddPlane("PLANE001")
	airport.AddPlane("PLANE002")
	airport.AddPlane("PLANE003")

	airport.Takeoff("PLANE002")

	planes := airport.GetPlaneList()
	if len(planes) != 2 {
		t.Errorf("Expected 2 planes after takeoff, got %d", len(planes))
	}

	// Check that PLANE002 is removed
	for _, p := range planes {
		if p == "PLANE002" {
			t.Error("PLANE002 should have been removed")
		}
	}

	// Test taking off non-existent plane
	airport.Takeoff("PLANE999")
	if len(airport.GetPlaneList()) != 2 {
		t.Errorf("Taking off non-existent plane should not change list size")
	}
}

func TestGetAirports(t *testing.T) {
	airports = []*Airport{}

	_ = NewAirport("Airport 1")
	_ = NewAirport("Airport 2")
	_ = NewAirport("Airport 3")

	allAirports := GetAirports()
	if len(allAirports) != 3 {
		t.Errorf("Expected 3 airports, got %d", len(allAirports))
	}

	// Verify all airports are in the list
	names := make(map[string]bool)
	for _, a := range allAirports {
		names[a.GetName()] = true
	}

	if !names["Airport 1"] || !names["Airport 2"] || !names["Airport 3"] {
		t.Error("Not all airports found in GetAirports()")
	}
}
