package models

// Crew represents a crew member in the airport
type Crew struct {
	Name       string
	Age        int
	Experience int
	Rating     float64
	License    string
	Status     string
}

// NewCrew creates a new Crew instance
func NewCrew(name string, age int, experience int, rating float64, license string, status string) *Crew {
	return &Crew{Name: name, Age: age, Experience: experience, Rating: rating, License: license, Status: status}
}

// GetCrew returns the Crew instance
func (c *Crew) GetCrew() *Crew {
	return c
}

// GetName returns the crew member's name
func (c *Crew) GetName() string {
	return c.Name
}

// GetAge returns the crew member's age
func (c *Crew) GetAge() int {
	return c.Age
}

// GetExperience returns the crew member's experience
func (c *Crew) GetExperience() int {
	return c.Experience
}

// GetRating returns the crew member's rating
func (c *Crew) GetRating() float64 {
	return c.Rating
}

// GetLicense returns the crew member's license
func (c *Crew) GetLicense() string {
	return c.License
}

// GetStatus returns the crew member's status
func (c *Crew) GetStatus() string {
	return c.Status
}

// SetName sets the crew member's name
func (c *Crew) SetName(name string) {
	c.Name = name
}

// SetAge sets the crew member's age
func (c *Crew) SetAge(age int) {
	c.Age = age
}

// SetExperience sets the crew member's experience
func (c *Crew) SetExperience(experience int) {
	c.Experience = experience
}

// SetRating sets the crew member's rating
func (c *Crew) SetRating(rating float64) {
	c.Rating = rating
}

// SetLicense sets the crew member's license
func (c *Crew) SetLicense(license string) {
	c.License = license
}

// SetStatus sets the crew member's status
func (c *Crew) SetStatus(status string) {
	c.Status = status
}

// GetCrewList returns the crew member's list
func (c *Crew) GetCrewList() []*Crew {
	return []*Crew{c}
}
