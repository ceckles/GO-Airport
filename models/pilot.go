package models

// Pilot represents a pilot in the airport
type Pilot struct {
	Name       string
	Age        int
	Experience int
	Rating     float64
	License    string
	Status     string
}

// NewPilot creates a new Pilot instance
func NewPilot(name string, age int, experience int, rating float64, license string, status string) *Pilot {
	return &Pilot{Name: name, Age: age, Experience: experience, Rating: rating, License: license, Status: status}
}

// GetPilot returns the Pilot instance
func (p *Pilot) GetPilot() *Pilot {
	return p
}

// GetName returns the pilot's name
func (p *Pilot) GetName() string {
	return p.Name
}

// GetAge returns the pilot's age
func (p *Pilot) GetAge() int {
	return p.Age
}

// GetExperience returns the pilot's experience
func (p *Pilot) GetExperience() int {
	return p.Experience
}

// GetRating returns the pilot's rating
func (p *Pilot) GetRating() float64 {
	return p.Rating
}

// GetLicense returns the pilot's license
func (p *Pilot) GetLicense() string {
	return p.License
}

// GetStatus returns the pilot's status
func (p *Pilot) GetStatus() string {
	return p.Status
}

// SetName sets the pilot's name
func (p *Pilot) SetName(name string) {
	p.Name = name
}

// SetAge sets the pilot's age
func (p *Pilot) SetAge(age int) {
	p.Age = age
}

// SetExperience sets the pilot's experience
func (p *Pilot) SetExperience(experience int) {
	p.Experience = experience
}

// SetRating sets the pilot's rating
func (p *Pilot) SetRating(rating float64) {
	p.Rating = rating
}

// SetLicense sets the pilot's license
func (p *Pilot) SetLicense(license string) {
	p.License = license
}

// SetStatus sets the pilot's status
func (p *Pilot) SetStatus(status string) {
	p.Status = status
}
