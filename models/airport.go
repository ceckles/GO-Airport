package models

// airports is a package-level slice that stores all airport instances
var airports []*Airport

// Airport represents an airport with a name and a list of planes
type Airport struct {
	name      string
	planeList []string // Using string for plane identifier, can be changed to a Plane struct if needed
}

// NewAirport creates a new Airport instance and adds it to the airports slice
func NewAirport(name string) *Airport {
	airport := &Airport{
		name:      name,
		planeList: make([]string, 0),
	}
	airports = append(airports, airport)
	return airport
}

// AddPlane adds a plane to the airport's plane list (only if not already present)
func (a *Airport) AddPlane(plane string) {
	// Check if plane already exists
	for _, p := range a.planeList {
		if p == plane {
			return // Plane already in list
		}
	}
	a.planeList = append(a.planeList, plane)
}

// Land adds a plane to the airport's plane list (only if not already present)
func (a *Airport) Land(plane string) {
	// Check if plane already exists
	for _, p := range a.planeList {
		if p == plane {
			return // Plane already in list
		}
	}
	a.planeList = append(a.planeList, plane)
}

// Takeoff removes a plane from the airport's plane list
func (a *Airport) Takeoff(plane string) {
	for i, p := range a.planeList {
		if p == plane {
			// Remove the plane by creating a new slice without it
			a.planeList = append(a.planeList[:i], a.planeList[i+1:]...)
			return
		}
	}
}

// GetAirports returns all airport instances
func GetAirports() []*Airport {
	return airports
}

// GetName returns the airport's name
func (a *Airport) GetName() string {
	return a.name
}

// GetPlaneList returns the airport's plane list
func (a *Airport) GetPlaneList() []string {
	return a.planeList
}
