package models

// Plane represents a plane in the airport
type Plane struct {
	ID             string
	Role           string
	Model          string
	Year           int
	Manufacturer   string
	Owner          string
	Price          float64
	Color          string
	Weight         float64
	MaxWeight      float64
	CurrentWeight  float64
	FuelCapacity   float64
	FuelLevel      float64
	Crew           []*Crew
	Pilots         []*Pilot
	PassengerSeats int
	Passengers     []*Passenger
	Destination    string
	Origin         string
	Status         string
	CargoWeight    float64
}

// NewPlane creates a new plane
func NewPlane(id string, role string, model string, year int, manufacturer string, owner string, price float64, color string, weight float64, maxWeight float64, currentWeight float64, fuelCapacity float64, fuelLevel float64, crew []*Crew, pilots []*Pilot, passengerSeats int, passengers []*Passenger, destination string, origin string, status string, cargoWeight float64) *Plane {
	return &Plane{ID: id, Role: role, Model: model, Year: year, Manufacturer: manufacturer, Owner: owner, Price: price, Color: color, Weight: weight, MaxWeight: maxWeight, CurrentWeight: currentWeight, FuelCapacity: fuelCapacity, FuelLevel: fuelLevel, Crew: crew, Pilots: pilots, PassengerSeats: passengerSeats, Passengers: passengers, Destination: destination, Origin: origin, Status: status, CargoWeight: cargoWeight}
}

// GetPlane returns the plane
func (p *Plane) GetPlane() *Plane {
	return p
}

// GetID returns the plane's ID
func (p *Plane) GetID() string {
	return p.ID
}

// GetRole returns the plane's role
func (p *Plane) GetRole() string {
	return p.Role
}

// GetModel returns the plane's model
func (p *Plane) GetModel() string {
	return p.Model
}

// GetYear returns the plane's year
func (p *Plane) GetYear() int {
	return p.Year
}

// GetManufacturer returns the plane's manufacturer
func (p *Plane) GetManufacturer() string {
	return p.Manufacturer
}

// GetOwner returns the plane's owner
func (p *Plane) GetOwner() string {
	return p.Owner
}

// GetPrice returns the plane's price
func (p *Plane) GetPrice() float64 {
	return p.Price
}

// GetColor returns the plane's color
func (p *Plane) GetColor() string {
	return p.Color
}

// GetWeight returns the plane's weight
func (p *Plane) GetWeight() float64 {
	return p.Weight
}

// GetMaxWeight returns the plane's max weight
func (p *Plane) GetMaxWeight() float64 {
	return p.MaxWeight
}

// GetCurrentWeight returns the plane's current weight
func (p *Plane) GetCurrentWeight() float64 {
	return p.CurrentWeight
}

// GetFuelCapacity returns the plane's fuel capacity
func (p *Plane) GetFuelCapacity() float64 {
	return p.FuelCapacity
}

// GetFuelLevel returns the plane's fuel level
func (p *Plane) GetFuelLevel() float64 {
	return p.FuelLevel
}

// GetCrew returns the plane's crew
func (p *Plane) GetCrew() []*Crew {
	return p.Crew
}

// GetPilots returns the plane's pilots
func (p *Plane) GetPilots() []*Pilot {
	return p.Pilots
}

// GetPassengerSeats returns the plane's passenger seats
func (p *Plane) GetPassengerSeats() int {
	return p.PassengerSeats
}

// GetPassengers returns the plane's passengers list
func (p *Plane) GetPassengers() []*Passenger {
	return p.Passengers
}

// GetPlaneList returns the plane's list
func (p *Plane) GetPlaneList() []Plane {
	return []Plane{*p}
}

// GetDestination returns the plane's destination
func (p *Plane) GetDestination() string {
	return p.Destination
}

// GetOrigin returns the plane's origin
func (p *Plane) GetOrigin() string {
	return p.Origin
}

// GetStatus returns the plane's status
func (p *Plane) GetStatus() string {
	return p.Status
}

// SetID sets the plane's ID
func (p *Plane) SetID(id string) {
	p.ID = id
}

// SetRole sets the plane's role
func (p *Plane) SetRole(role string) {
	p.Role = role
}

// SetModel sets the plane's model
func (p *Plane) SetModel(model string) {
	p.Model = model
}

// SetYear sets the plane's year
func (p *Plane) SetYear(year int) {
	p.Year = year
}

// SetManufacturer sets the plane's manufacturer
func (p *Plane) SetManufacturer(manufacturer string) {
	p.Manufacturer = manufacturer
}

// SetOwner sets the plane's owner
func (p *Plane) SetOwner(owner string) {
	p.Owner = owner
}

// SetPrice sets the plane's price
func (p *Plane) SetPrice(price float64) {
	p.Price = price
}

// SetColor sets the plane's color
func (p *Plane) SetColor(color string) {
	p.Color = color
}

// SetWeight sets the plane's weight
func (p *Plane) SetWeight(weight float64) {
	p.Weight = weight
}

// SetMaxWeight sets the plane's max weight
func (p *Plane) SetMaxWeight(maxWeight float64) {
	p.MaxWeight = maxWeight
}

// SetCurrentWeight sets the plane's current weight
func (p *Plane) SetCurrentWeight(currentWeight float64) {
	p.CurrentWeight = currentWeight
}

// SetFuelCapacity sets the plane's fuel capacity
func (p *Plane) SetFuelCapacity(fuelCapacity float64) {
	p.FuelCapacity = fuelCapacity
}

// SetFuelLevel sets the plane's fuel level
func (p *Plane) SetFuelLevel(fuelLevel float64) {
	p.FuelLevel = fuelLevel
}

// SetCrew sets the plane's crew
func (p *Plane) SetCrew(crew []*Crew) {
	p.Crew = crew
}

// SetPilots sets the plane's pilots
func (p *Plane) SetPilots(pilots []*Pilot) {
	p.Pilots = pilots
}

// SetPassengerSeats sets the plane's passenger seats
func (p *Plane) SetPassengerSeats(passengerSeats int) {
	p.PassengerSeats = passengerSeats
}

// SetPassengers sets the plane's passengers list
func (p *Plane) SetPassengers(passengers []*Passenger) {
	p.Passengers = passengers
}

// AddPassenger adds a passenger to the plane's passengers list
func (p *Plane) AddPassenger(passenger *Passenger) {
	p.Passengers = append(p.Passengers, passenger)
}

// RemovePassenger removes a passenger from the plane's passengers list
func (p *Plane) RemovePassenger(passenger *Passenger) {
	for i, pass := range p.Passengers {
		if pass == passenger {
			p.Passengers = append(p.Passengers[:i], p.Passengers[i+1:]...)
			return
		}
	}
}

// AddCrew adds a crew member to the plane's crew list
func (p *Plane) AddCrew(crew *Crew) {
	p.Crew = append(p.Crew, crew)
}

// RemoveCrew removes a crew member from the plane's crew list
func (p *Plane) RemoveCrew(crew *Crew) {
	for i, c := range p.Crew {
		if c == crew {
			p.Crew = append(p.Crew[:i], p.Crew[i+1:]...)
			return
		}
	}
}

// SetDestination sets the plane's destination
func (p *Plane) SetDestination(destination string) {
	p.Destination = destination
}

// SetOrigin sets the plane's origin
func (p *Plane) SetOrigin(origin string) {
	p.Origin = origin
}

// SetStatus sets the plane's status
func (p *Plane) SetStatus(status string) {
	p.Status = status
}

// GetCargoWeight returns the plane's cargo weight
func (p *Plane) GetCargoWeight() float64 {
	return p.CargoWeight
}

// SetCargoWeight sets the plane's cargo weight
func (p *Plane) SetCargoWeight(cargoWeight float64) {
	p.CargoWeight = cargoWeight
}

// GetCalculatedWeight returns the plane's calculated weight
// Includes: base weight + current weight + cargo weight + all passenger weights + all passenger baggage weights
func (p *Plane) GetCalculatedWeight() float64 {
	totalWeight := p.Weight + p.CurrentWeight + p.CargoWeight

	// Add passenger weights and baggage weights
	for _, passenger := range p.Passengers {
		if passenger != nil {
			totalWeight += passenger.PassangerWeight
			totalWeight += passenger.BaggageWeight
		}
	}

	return totalWeight
}
