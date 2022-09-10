package chi_router

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"flightTracker/internal/service/flight_tracker"
)

type controller struct {
	useCase flight_tracker.UseCase
}

func New(useCase flight_tracker.UseCase) flight_tracker.Transport {
	return &controller{
		useCase: useCase,
	}
}

// GetFlightPath handles http requests and call useCase layer
func (c controller) GetFlightPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var flightList [][]string
	err := json.NewDecoder(r.Body).Decode(&flightList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := c.useCase.GetFlightPath(flightList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		zap.L().Error("failed to marshal json", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
