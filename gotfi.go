package gotfi

// Operator is the global interface implemented by all
// operator subpackages
type Operator interface {
	// Stop-based informaiton
	GetStop(id string) (Stop, error)
	GetStops() (map[string]Stop, error)

	// Route-based information
	GetRoute(id string) (Route, error)
	GetRoutes() (map[string]Route, error)

	// Real-time Information
	GetStopRTPI(id string) (RTResult, error)
}
