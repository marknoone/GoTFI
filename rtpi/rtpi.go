package rtpi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	tfi "github.com/marknoone/GoTFI"
)

const rtpiBaseURL = "https://data.smartdublin.ie/cgi-bin/rtpi/realtimebusinformation"

var OpCodes = struct {
	DublinBus  opcode
	GoAhead    opcode
	BusEireann opcode
	KildareBus opcode
}{
	DublinBus:  opcode("bac"),
	GoAhead:    opcode("GAD"),
	BusEireann: opcode("BE"),
	KildareBus: opcode("KB"),
}

type opcode string
type rtpiOperator struct {
	opCode string
}

// NewRtpiOperator creates a new operator object to interact
// with  RTPI services.
func NewRtpiOperator(o opcode) rtpiOperator {
	return rtpiOperator{string(o)}
}

// GetStop returns all of the information of a given stop
// operated by the RTPI operator
func (r rtpiOperator) GetStop(id string) (tfi.Stop, error) {
	var (
		resp stopResponse
		uri  = fmt.Sprintf(
			"%s/busstopinformation?stopid=%s&operator=%s&format=json", rtpiBaseURL, id, r)
	)
	err := r.makeRequest(uri, resp)
	if err != nil {
		return tfi.Stop{}, err
	}

	var routes []string
	st := resp.Results[0]
	for _, op := range st.Operators {
		if op.Name == r.opCode {
			routes = op.Routes
			break
		}
	}

	return tfi.Stop{
		ID:     st.ID,
		Name:   st.ShortName,
		Lat:    parseFloat(st.Lat),
		Lng:    parseFloat(st.Lng),
		Routes: routes,
	}, nil
}

// GetStops returns all stops ran by the given operator
func (r rtpiOperator) GetStops() (map[string]tfi.Stop, error) {
	var (
		resp stopResponse
		uri  = fmt.Sprintf(
			"%s/busstopinformation?operator=%s&format=json", rtpiBaseURL, r)
	)
	err := r.makeRequest(uri, resp)
	if err != nil {
		return nil, err
	}

	stopMap := make(map[string]tfi.Stop, len(resp.Results))
	for _, stop := range resp.Results {
		var routes []string
		for _, op := range stop.Operators {
			if op.Name == r.opCode {
				routes = op.Routes
				break
			}
		}

		stopMap[stop.ID] = tfi.Stop{
			ID:     stop.ID,
			Name:   stop.ShortName,
			Lat:    parseFloat(stop.Lat),
			Lng:    parseFloat(stop.Lng),
			Routes: routes,
		}
	}

	return stopMap, nil
}

// GetRoute returns all of the information of a given route
// operated by the RTPI operator
func (r rtpiOperator) GetRoute(id string) (tfi.Route, error) {
	var (
		resp routeResponse
		uri  = fmt.Sprintf(
			"%s/routeinformation?route=%s&operator=%s&format=json", rtpiBaseURL, id, r)
	)
	err := r.makeRequest(uri, resp)
	if err != nil {
		return tfi.Route{}, err
	}

	var stops = make([]string, len(resp.Results[0].Stops))
	for i, stop := range resp.Results[0].Stops {
		stops[i] = stop.ID
	}

	return tfi.Route{
		ID:      resp.Route,
		StopIDs: stops,
	}, nil
}

// GetRoutes returns all routes ran by the given operator
func (r rtpiOperator) GetRoutes() (map[string]tfi.Route, error) {
	var (
		resp routelistResponse
		uri  = fmt.Sprintf(
			"%s/routelistinformation?operator=%s&format=json", rtpiBaseURL, r)
	)
	err := r.makeRequest(uri, resp)
	if err != nil {
		return nil, err
	}

	routeMap := make(map[string]tfi.Route, len(resp.Results))
	for _, route := range resp.Results {
		routeMap[route.Route] = tfi.Route{ID: route.Route}
	}

	return routeMap, nil
}

// GetStopRTPI retrieves all of the Real-time Information of a given stop
func (r rtpiOperator) GetStopRTPI(id string) (tfi.RTResult, error) {
	var (
		resp rtResponse
		uri  = fmt.Sprintf(
			"%s/realtimebusinformation?stopid=%s&operator=%s&format=json", rtpiBaseURL, id, r)
	)

	parseInt := func(s string) int {
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			log.Println("float conversion error: number " + s)
		}
		return int(i)
	}

	err := r.makeRequest(uri, resp)
	if err != nil {
		return tfi.RTResult{}, err
	}

	var results = make([]tfi.RTResultRow, len(resp.Results))
	for i, res := range resp.Results {
		results[i] = tfi.RTResultRow{
			Route:       res.Route,
			Destination: res.Destination,
			Due:         parseInt(res.DueTime),
			Direction:   res.Direction,
		}
	}

	return tfi.RTResult{
		Timestamp: resp.Timestamp,
		Results:   results,
	}, nil
}

// GetOperators retrieves all operator references used by the RTPI API
func (r rtpiOperator) GetOperators() ([]OperatorRef, error) {
	var (
		resp operatorResponse
		uri  = fmt.Sprintf("%s/operatorinformation?format=json", rtpiBaseURL)
	)
	err := r.makeRequest(uri, resp)
	return resp.Results, err
}

func (rtpiOperator) makeRequest(uri string, dest interface{}) error {
	reqURL, err := url.Parse(uri)
	if err != nil {
		return err
	}

	resp, err := http.Get(reqURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return json.NewDecoder(resp.Body).Decode(dest)
	default:
		return fmt.Errorf(
			"request error occured: recieved http status code (%d)",
			resp.StatusCode)
	}
}

func parseFloat(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Println("float conversion error: number " + s)
	}
	return float32(f)
}
