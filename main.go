// main go file
package main

import (
	"encoding/json"
	"fmt"

	"github.com/chad/GO-AirPort/models"
)

func main() {
	// Create airports
	jfk := models.NewAirport("JFK")
	heathrow := models.NewAirport("Heathrow")

	// Create pilots
	pilot1 := models.NewPilot("John Doe", 30, 10, 9.5, "1234567890", "Active")
	pilot2 := models.NewPilot("Jane Doe", 25, 5, 8.5, "0987654321", "Active")

	// Create crew members
	crew1 := models.NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")
	crew2 := models.NewCrew("Alice Johnson", 32, 12, 9.0, "CREW002", "Active")
	crew3 := models.NewCrew("Charlie Brown", 26, 6, 7.5, "CREW003", "Active")

	// Create passengers
	passenger1 := models.NewPassenger("John Doe", 45, "12A")
	passenger1.SetStatus("Checked In")
	passenger1.SetBaggage(2)
	passenger1.SetBaggageWeight(25.5)
	passenger1.SetPassangerWeight(80.0)
	passenger1.SetTicketNumber("TKT001")

	passenger2 := models.NewPassenger("Jane Doe", 42, "12B")
	passenger2.SetStatus("Checked In")
	passenger2.SetBaggage(1)
	passenger2.SetBaggageWeight(18.0)
	passenger2.SetPassangerWeight(65.0)
	passenger2.SetTicketNumber("TKT002")

	// Create plane
	plane1 := models.NewPlane("1", "Passenger", "Boeing 747", 2024, "Boeing", "John Doe", 1000000, "White", 100000, 200000, 150000, 100000, 80000, []*models.Crew{crew1, crew2, crew3}, []*models.Pilot{pilot1, pilot2}, 100, []*models.Passenger{passenger1, passenger2}, "London", "New York", "Landed", 50000)

	// Display created plane info in JSON format
	planeJSON, err := json.MarshalIndent(plane1.GetPlane(), "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling plane: %v\n", err)
	} else {
		fmt.Println("Plane Details (JSON):")
		fmt.Println(string(planeJSON))
		fmt.Println()
	}

	// Display plane calculated weight
	fmt.Printf("\nPlane Calculated Weight: %.2f kg\n", plane1.GetCalculatedWeight())

	// Add plane to airport
	jfk.AddPlane(plane1.GetID())

	// Land a plane (will check for duplicates)
	jfk.Land(plane1.GetID())

	// Takeoff a plane
	//jfk.Takeoff(plane1.GetID())

	// Display airports
	fmt.Printf("\nAirport: %s, Planes: %v\n", jfk.GetName(), jfk.GetPlaneList())
	fmt.Printf("Airport: %s, Planes: %v\n", heathrow.GetName(), heathrow.GetPlaneList())
	fmt.Printf("Total airports: %d\n", len(models.GetAirports()))
}
