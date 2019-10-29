package gotfi

// Stop contains all information about a transport stop for a
// given operator.
type Stop struct {
	ID   string
	Name string
	Lat  float32
	Lng  float32

	// Routes contains all routes served by this stop
	Routes []string
}

// Route contains all information of a given route
type Route struct {
	ID      string
	StopIDs []string
}

// RTResult is the real-time response of an operator
type RTResult struct {
	Timestamp string
	Results   []RTResultRow
}

// RTResultRow contains RealTime data of a given stop
type RTResultRow struct {
	Route       string
	Destination string
	Due         int
	Direction   string
}
