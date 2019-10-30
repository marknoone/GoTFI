package rtpi

type operatorResponse struct {
	ErrCode         string        `json:"errorcode"`       // Error code of response
	ErrMessage      string        `json:"errormessage"`    // Additional error details
	NumberOfResults int           `json:"numberofresults"` // Number of results
	Timestamp       string        `json:"timestamp"`       // Date/time of informaiton returned
	Results         []OperatorRef `json:"results"`
}

type OperatorRef struct {
	Reference   string `json:"operatorreference"`   // Operator reference code
	Name        string `json:"operatorname"`        // Operator name
	Description string `json:"operatordescription"` // Operator description
}

type routeResponse struct {
	ErrCode         string  `json:"errorcode"`       // Error code of response
	ErrMessage      string  `json:"errormessage"`    // Additional error details
	NumberOfResults int     `json:"numberofresults"` // Number of results
	Timestamp       string  `json:"timestamp"`       // Date/time of informaiton returned
	Route           string  `json:"route"`           // The stop's route
	Results         []Route `json:"results"`
}

type Route struct {
	Operator             string `json:"operator"`             // Servicing operator
	Origin               string `json:"origin"`               // Origin name
	OriginLocalized      string `json:"originlocalized"`      // Translated origin name
	Destination          string `json:"destination"`          // Destination  name
	DestinationLocalized string `json:"destinationlocalized"` // Translated destination  name
	LastUpdated          string `json:"lastupdated"`          // Last updated date/time for this record
	Stops                []Stop `json:"stops"`
}

type stopResponse struct {
	ErrCode         string `json:"errorcode"`       // Error code of response
	ErrMessage      string `json:"errormessage"`    // Additional error details
	NumberOfResults int    `json:"numberofresults"` // Number of results
	Timestamp       string `json:"timestamp"`       // Date/time of informaiton returned
	Results         []Stop `json:"results"`
}

type Stop struct {
	ID                 string     `json:"stopid"`             // Stop ID
	DisplayID          string     `json:"displaystopid"`      // Stop identifier to display for the end user
	ShortName          string     `json:"shortname"`          // Stop name
	ShortNameLoclaized string     `json:"shortnamelocalized"` // Translated stop name
	FullName           string     `json:"fullname"`           // Full name
	FullNameLocalized  string     `json:"fullnamelocalized"`  // Translated full name
	Lat                string     `json:"latitude"`           // Latitude
	Lng                string     `json:"longitude"`          // Longitude
	Operators          []Operator `json:"operators"`
}

type Operator struct {
	Name   string   `json:"name"`         // Operator name
	Type   int      `json:"operatortype"` // Operator type
	Routes []string `json:"routes"`       // Serviced routes
}

type routelistResponse struct {
	ErrCode         string          `json:"errorcode"`       // Error code of response
	ErrMessage      string          `json:"errormessage"`    // Additional error details
	NumberOfResults int             `json:"numberofresults"` // Number of results
	Timestamp       string          `json:"timestamp"`       // Date/time of informaiton returned
	Results         []RouteOverview `json:"results"`
}

type RouteOverview struct {
	Operator     string `json:"operator"`     // Operator reference
	OperatorType int    `json:"operatortype"` // Operator type
	Route        string `json:"route"`        // Route
}

type timetableResponse struct {
	ErrCode         string `json:"errorcode"`       // Error code of response
	ErrMessage      string `json:"errormessage"`    // Additional error details
	NumberOfResults int    `json:"numberofresults"` // Number of results
	Timestamp       string `json:"timestamp"`       // Date/time of informaiton returned
	Route           string `json:"route"`           // The stop's route
	StopID          string `json:"stopid"`          // The stop's ID
	Results         []Stop `json:"results"`
}

type Timetable struct {
	StartDayOfWeek       string   `json:"startdayofweek"`       // Day of week (0 - Sunday, 1 - Monday ...)
	EndDayOfWeek         string   `json:"enddayofweek"`         // Day of week (0 - Sunday, 1 - Monday ...)
	Destination          string   `json:"destination"`          // Service destination name
	DestinationLocalized string   `json:"destinationlocalized"` // Translated service destination name
	LastUpdated          string   `json:"lastupdated"`          // Last update date/time for this record
	Departures           []string `json:"departures"`           // Departures time in dd/MM/yyyy HH:mm:ss format
}

type rtResponse struct {
	ErrCode         string   `json:"errorcode"`       // Error code of response
	ErrMessage      string   `json:"errormessage"`    // Additional error details
	NumberOfResults int      `json:"numberofresults"` // Number of results
	Timestamp       string   `json:"timestamp"`       // Date/time of informaiton returned
	StopID          string   `json:"stopid"`          // The stop's ID
	Results         []RTData `json:"results"`
}

type RTData struct {
	ArrivalDateTime            string `json:"arrivaldatetime"`            // Arrival time in dd/MM/yyyy HH:mm:ss format
	DueTime                    string `json:"duetime"`                    // Arrival due time in minutes
	DepartureDateTime          string `json:"departuredatetime"`          // Departure time in dd/MM/yyyy HH:mm:ss format
	DepartureDueTime           string `json:"departureduetime"`           // Departure due time in minutes
	ScheduledArrivalDateTime   string `json:"scheduledarrivaldatetime"`   // Scheduled arrival time in dd/MM/yyyy HH:mm:ss format
	ScheduledDepartureDateTime string `json:"scheduleddeparturedatetime"` // Scheduled departure time in dd/MM/yyyy HH:mm:ss format
	Destination                string `json:"destination"`                // Service destination name
	DestinationLocalized       string `json:"destinationlocalized"`       // Translated service destination name
	Origin                     string `json:"origin"`                     // Service origin name
	OriginLocalized            string `json:"originlocalized"`            // Translated service origin name
	Direction                  string `json:"direction"`                  // Service direction
	Operator                   string `json:"operator"`                   // Bus service operator name
	OperatorType               string `json:"operatortype"`               // Bus service operator type
	AdditionalInformation      string `json:"additionalinformation"`      // Additional information
	LowFloorStatus             string `json:"lowfloorstatus"`             // Bus low floor status (yes or no)
	Route                      string `json:"route"`                      // Bus route
	SourceTimestamp            string `json:"sourcetimestamp"`            // Timestamp in dd/MM/yyyy HH:mm:ss format
	Monitored                  string `json:"monitored"`
}
