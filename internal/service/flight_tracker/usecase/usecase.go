package usecase

import (
	"errors"

	"flightTracker/internal/service/flight_tracker"
	"flightTracker/internal/service/flight_tracker/types"
)

const (
	oneFlightShouldContain2AirportsError = "one flight should contain 2 airports"
	theFlightListIsNotConsistentError    = "the flight list is not consistent"
)

type useCase struct {
	// here should be repo layer
}

func New() flight_tracker.UseCase {
	return &useCase{}
}

// GetFlightPath returns source and destination airports based on the whole flight path.
// If the flightList is consistent, the source and destination airports must be without pair.
// In the code we go over flightList and check if each airport is present in map. If no - we add it, if yes, we remove it.
// The map structure is map[string]types.PointType, where types.PointType marks whether the airport is source or destination.
// We will end up with map with two elements - and this is what we were looking for.
// The algorithm complexity is O(n)
func (uc useCase) GetFlightPath(flightList [][]string) ([]string, error) {
	pointsMap := make(map[string]types.PointType)
	var source, destination string
	var ok bool
	for _, flightPair := range flightList {
		if len(flightPair) != 2 {
			return nil, errors.New(oneFlightShouldContain2AirportsError)
		}

		// check source airport from single flight
		source = flightPair[0]
		_, ok = pointsMap[source]
		if !ok {
			pointsMap[source] = types.Source
		} else {
			delete(pointsMap, source)
		}

		// check destination airport from single flight
		destination = flightPair[1]
		_, ok = pointsMap[destination]
		if !ok {
			pointsMap[destination] = types.Destination
		} else {
			delete(pointsMap, destination)
		}
	}

	// we should end up with two keys in map, otherwise the flight list is inconsistent
	if len(pointsMap) != 2 {
		return nil, errors.New(theFlightListIsNotConsistentError)
	}

	finalPath := make([]string, 2)
	for key, val := range pointsMap {
		switch val {
		case types.Source:
			finalPath[0] = key
		case types.Destination:
			finalPath[1] = key
		}
	}

	return finalPath, nil
}
