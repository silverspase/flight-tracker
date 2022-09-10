package dep_container

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sarulabs/di"
	"go.uber.org/zap"

	"flightTracker/internal/config"                 // nolint:typecheck
	"flightTracker/internal/service/flight_tracker" // nolint:typecheck
)

const (
	httpServerDefName     = "http-server"
	GetFlightPathEndpoint = "get-flight-path"
)

// RegisterHTTPServer registers HTTP Server dependency.
func RegisterHTTPServer(builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: httpServerDefName,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(configDefName).(*config.Config)
			flightTrackerService := ctn.Get(flightTrackerDefName).(flight_tracker.Transport)

			router := chi.NewRouter()
			router.HandleFunc(fmt.Sprintf("/%s", GetFlightPathEndpoint), flightTrackerService.GetFlightPath)

			zap.L().Info("started http server",
				zap.String("address", fmt.Sprintf("http://localhost:%s/", cfg.Port)))
			zap.S().Fatal("", zap.Error(http.ListenAndServe(":"+cfg.Port, router)))

			return router, nil
		},
	})
}

// RunHTTPServer runs HTTP Server dependency.
func (c Container) RunHTTPServer() {
	c.container.Get(httpServerDefName)
}
