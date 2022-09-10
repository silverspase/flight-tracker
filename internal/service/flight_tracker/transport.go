package flight_tracker

import "net/http"

type Transport interface {
	GetFlightPath(w http.ResponseWriter, r *http.Request)
}
