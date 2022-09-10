package dep_container

import (
	"github.com/sarulabs/di"

	"flightTracker/internal/service/flight_tracker/transport/chi_router"
	"flightTracker/internal/service/flight_tracker/usecase"
)

const flightTrackerDefName = "flight-tracker"

// RegisterFlightTrackerService registers FlightTrackerService dependency.
func RegisterFlightTrackerService(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: flightTrackerDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			return chi_router.New(usecase.New()), nil
		},
	})
}
