package models

// Passenger represents a passenger in the airport
type Passenger struct {
	Name            string
	Age             int
	Seat            string
	Status          string
	Baggage         int
	BaggageWeight   float64
	PassangerWeight float64
	TicketNumber    string
}

// NewPassenger creates a new Passenger instance
func NewPassenger(name string, age int, seat string) *Passenger {
	return &Passenger{Name: name, Age: age, Seat: seat}
}

// GetPassenger returns the Passenger instance
func (p *Passenger) GetPassenger() *Passenger {
	return p
}

// GetName returns the passenger's name
func (p *Passenger) GetName() string {
	return p.Name
}

// GetAge returns the passenger's age
func (p *Passenger) GetAge() int {
	return p.Age
}

// GetSeat returns the passenger's seat
func (p *Passenger) GetSeat() string {
	return p.Seat
}

// GetStatus returns the passenger's status
func (p *Passenger) GetStatus() string {
	return p.Status
}

// GetBaggage returns the passenger's baggage
func (p *Passenger) GetBaggage() int {
	return p.Baggage
}

// GetBaggageWeight returns the passenger's baggage weight
func (p *Passenger) GetBaggageWeight() float64 {
	return p.BaggageWeight
}

// GetPassangerWeight returns the passenger's weight
func (p *Passenger) GetPassangerWeight() float64 {
	return p.PassangerWeight
}

// GetTicketNumber returns the passenger's ticket number
func (p *Passenger) GetTicketNumber() string {
	return p.TicketNumber
}

// SetName sets the passenger's name
func (p *Passenger) SetName(name string) {
	p.Name = name
}

// SetAge sets the passenger's age
func (p *Passenger) SetAge(age int) {
	p.Age = age
}

// SetSeat sets the passenger's seat
func (p *Passenger) SetSeat(seat string) {
	p.Seat = seat
}

// SetStatus sets the passenger's status
func (p *Passenger) SetStatus(status string) {
	p.Status = status
}

// SetBaggage sets the passenger's baggage
func (p *Passenger) SetBaggage(baggage int) {
	p.Baggage = baggage
}

// SetBaggageWeight sets the passenger's baggage weight
func (p *Passenger) SetBaggageWeight(baggageWeight float64) {
	p.BaggageWeight = baggageWeight
}

// SetPassangerWeight sets the passenger's weight
func (p *Passenger) SetPassangerWeight(passangerWeight float64) {
	p.PassangerWeight = passangerWeight
}

// SetTicketNumber sets the passenger's ticket number
func (p *Passenger) SetTicketNumber(ticketNumber string) {
	p.TicketNumber = ticketNumber
}
