package types

// PointType determines whether airport is origin or destination point
type PointType string

const (
	Source      PointType = "source"
	Destination PointType = "destination"
)
