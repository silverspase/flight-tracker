package flight_tracker

type UseCase interface {
	GetFlightPath([][]string) ([]string, error)
}
