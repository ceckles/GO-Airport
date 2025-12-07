package models

import "testing"

func TestNewPassenger(t *testing.T) {
	passenger := NewPassenger("John Doe", 45, "12A")

	if passenger == nil {
		t.Error("NewPassenger returned nil")
	}

	if passenger.GetName() != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", passenger.GetName())
	}

	if passenger.GetAge() != 45 {
		t.Errorf("Expected age 45, got %d", passenger.GetAge())
	}

	if passenger.GetSeat() != "12A" {
		t.Errorf("Expected seat '12A', got '%s'", passenger.GetSeat())
	}
}

func TestPassenger_GettersAndSetters(t *testing.T) {
	passenger := NewPassenger("John Doe", 45, "12A")

	// Test setters
	passenger.SetName("Jane Doe")
	if passenger.GetName() != "Jane Doe" {
		t.Errorf("SetName failed, expected 'Jane Doe', got '%s'", passenger.GetName())
	}

	passenger.SetAge(42)
	if passenger.GetAge() != 42 {
		t.Errorf("SetAge failed, expected 42, got %d", passenger.GetAge())
	}

	passenger.SetSeat("15B")
	if passenger.GetSeat() != "15B" {
		t.Errorf("SetSeat failed, expected '15B', got '%s'", passenger.GetSeat())
	}

	passenger.SetStatus("Checked In")
	if passenger.GetStatus() != "Checked In" {
		t.Errorf("SetStatus failed, expected 'Checked In', got '%s'", passenger.GetStatus())
	}

	passenger.SetBaggage(2)
	if passenger.GetBaggage() != 2 {
		t.Errorf("SetBaggage failed, expected 2, got %d", passenger.GetBaggage())
	}

	passenger.SetBaggageWeight(25.5)
	if passenger.GetBaggageWeight() != 25.5 {
		t.Errorf("SetBaggageWeight failed, expected 25.5, got %.2f", passenger.GetBaggageWeight())
	}

	passenger.SetPassangerWeight(80.0)
	if passenger.GetPassangerWeight() != 80.0 {
		t.Errorf("SetPassangerWeight failed, expected 80.0, got %.2f", passenger.GetPassangerWeight())
	}

	passenger.SetTicketNumber("TKT001")
	if passenger.GetTicketNumber() != "TKT001" {
		t.Errorf("SetTicketNumber failed, expected 'TKT001', got '%s'", passenger.GetTicketNumber())
	}
}

func TestPassenger_GetPassenger(t *testing.T) {
	passenger := NewPassenger("John Doe", 45, "12A")

	returnedPassenger := passenger.GetPassenger()
	if returnedPassenger != passenger {
		t.Error("GetPassenger should return the same passenger instance")
	}
}
