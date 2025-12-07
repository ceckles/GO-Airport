# GO-AirPort

![Go Version](https://img.shields.io/badge/Go-1.22.2-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Build Status](https://github.com/ceckles/GO-Airport/workflows/Build/badge.svg)
![Test Status](https://github.com/ceckles/GO-Airport/workflows/Test/badge.svg)

A Go-based airport management system that simulates airport operations including planes, pilots, crew, passengers, and airports. this was a project I did while in a Code bootcamp back in 2021, I wanted to port it to Go.

## Tech Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![JSON](https://img.shields.io/badge/JSON-000000?style=for-the-badge&logo=json&logoColor=white)
![Testing](https://img.shields.io/badge/Testing-Go%20Testing-00ADD8?style=for-the-badge&logo=go&logoColor=white)

> **Note**: The Build and Test badges above are dynamically updated by GitHub Actions. They reflect the actual status of the latest CI/CD pipeline runs.

## Features

- **Airport Management**: Create and manage multiple airports
- **Plane Management**: Track planes with detailed information (model, weight, fuel, status, etc.)
- **Pilot Management**: Manage pilots with ratings, experience, and licenses
- **Crew Management**: Handle flight crew members
- **Passenger Management**: Track passengers with baggage and weight information
- **Weight Calculations**: Automatic calculation of total plane weight including passengers and baggage
- **JSON Output**: Pretty-printed JSON output for plane details

## Project Structure

```
GO-AirPort/
├── models/
│   ├── airport.go      # Airport model and operations
│   ├── plane.go        # Plane model and operations
│   ├── pilot.go        # Pilot model
│   ├── crew.go         # Crew member model
│   ├── passenger.go    # Passenger model
│   └── *_test.go       # Test files for each model
├── main.go             # Main application entry point
├── go.mod              # Go module file
└── README.md           # This file
```

## Prerequisites

- Go 1.22.2 or higher
- Git (optional, for cloning)

## Setup

1. **Clone the repository** (if applicable):
   ```bash
   git clone <repository-url>
   cd GO-AirPort
   ```

2. **Install dependencies** (if any):
   ```bash
   go mod download
   ```

3. **Verify installation**:
   ```bash
   go version
   ```

## Running the Application

### Run the main application:

```bash
go run main.go
```

Or build and run:

```bash
go build
./GO-AirPort
```

### Expected Output

The application will:
- Create airports (JFK, Heathrow)
- Create pilots, crew members, and passengers
- Create a plane with all associated personnel
- Display plane details in JSON format
- Calculate and display total plane weight
- Show airport plane lists

## Running Tests

### Run all tests:

```bash
go test ./models/... -v
```

### Run tests for a specific package:

```bash
go test ./models/ -v
```

### Run tests with coverage:

```bash
go test ./models/... -cover
```

### Run tests with detailed coverage report:

```bash
go test ./models/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Run a specific test:

```bash
go test ./models/ -run TestNewAirport -v
```

## Test Coverage

The project includes comprehensive test coverage for all models:

- **airport_test.go**: Tests for airport creation, plane management (add, land, takeoff)
- **plane_test.go**: Tests for plane creation, weight calculations, passenger/crew management
- **pilot_test.go**: Tests for pilot creation and getters/setters
- **crew_test.go**: Tests for crew member creation and getters/setters
- **passenger_test.go**: Tests for passenger creation and getters/setters

## Usage Examples

### Creating an Airport

```go
airport := models.NewAirport("JFK")
```

### Creating a Pilot

```go
pilot := models.NewPilot("John Doe", 30, 10, 9.5, "LIC123", "Active")
```

### Creating a Crew Member

```go
crew := models.NewCrew("Bob Smith", 28, 8, 8.0, "CREW001", "Active")
```

### Creating a Passenger

```go
passenger := models.NewPassenger("Jane Doe", 42, "12A")
passenger.SetStatus("Checked In")
passenger.SetBaggage(2)
passenger.SetBaggageWeight(25.5)
passenger.SetPassangerWeight(80.0)
```

### Creating a Plane

```go
plane := models.NewPlane(
    "PLANE001",           // ID
    "Passenger",          // Role
    "Boeing 747",         // Model
    2024,                 // Year
    "Boeing",            // Manufacturer
    "Owner Name",         // Owner
    1000000,              // Price
    "White",              // Color
    100000,               // Weight
    200000,               // MaxWeight
    150000,               // CurrentWeight
    100000,               // FuelCapacity
    80000,                // FuelLevel
    []*models.Crew{crew}, // Crew
    []*models.Pilot{pilot}, // Pilots
    100,                  // PassengerSeats
    []*models.Passenger{passenger}, // Passengers
    "London",             // Destination
    "New York",           // Origin
    "Landed",             // Status
    50000,                // CargoWeight
)
```

### Calculating Plane Weight

The `GetCalculatedWeight()` method automatically calculates:
- Base plane weight
- Current weight
- Cargo weight
- All passenger weights
- All passenger baggage weights

```go
totalWeight := plane.GetCalculatedWeight()
```

### Airport Operations

```go
// Add plane to airport
airport.AddPlane(plane.GetID())

// Land a plane (prevents duplicates)
airport.Land(plane.GetID())

// Takeoff a plane
airport.Takeoff(plane.GetID())

// Get all planes at airport
planes := airport.GetPlaneList()
```

## Building

### Build the application:

```bash
go build
```

This creates an executable file named `GO-AirPort` (or `GO-AirPort.exe` on Windows).

### Build for different platforms:

```bash
# Linux
GOOS=linux GOARCH=amd64 go build

# Windows
GOOS=windows GOARCH=amd64 go build

# macOS
GOOS=darwin GOARCH=amd64 go build
```

## Code Quality

### Format code:

```bash
go fmt ./...
```

### Run linter (if golangci-lint is installed):

```bash
golangci-lint run
```

### Check for issues:

```bash
go vet ./...
```

## Model Details

### Airport
- Name
- Plane list (prevents duplicates)

### Plane
- ID, Role, Model, Year, Manufacturer, Owner
- Price, Color
- Weight, MaxWeight, CurrentWeight
- FuelCapacity, FuelLevel
- Crew (array of Crew objects)
- Pilots (array of Pilot objects)
- PassengerSeats
- Passengers (array of Passenger objects)
- Destination, Origin, Status
- CargoWeight

### Pilot
- Name, Age, Experience
- Rating, License, Status

### Crew
- Name, Age, Experience
- Rating, License, Status

### Passenger
- Name, Age, Seat
- Status, Baggage, BaggageWeight
- PassangerWeight, TicketNumber

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

