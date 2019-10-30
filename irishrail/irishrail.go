package irishrail

import (
	"errors"
	"fmt"
	"time"

	tfi "github.com/marknoone/GoTFI"
)

const irBaseURL = "http://api.irishrail.ie/realtime/realtime.asmx"

var (
	_            tfi.Operator = Irishrail{}
	typeMapping               = map[string]string{"D": "Dart", "S": "Suburban", "M": "Mainline", "A": "All"}
	routeMapping              = map[string]string{"Dart": "D", "Suburban": "S", "Mainline": "M", "All": "A"}
)

// Irishrail struct to wrap all associated functions
type Irishrail struct{}

// GetStationsFilter retiurieves all stations that contained the passed in string
func (i Irishrail) GetStationsFilter(f string) ([]StationFilter, error) {
	var (
		resp arrayOfObjStationFilter
		uri  = fmt.Sprintf("%s/getStationsFilterXML?StationText=%s", irBaseURL, f)
	)
	err := i.makeRequest(uri, resp)
	return resp.StationFilters, err
}

// GetTrainMovements will retrieve a given trians historical movements
func (i Irishrail) GetTrainMovements(id string, date time.Time) ([]TrainMovements, error) {
	var (
		resp arrayOfObjTrainMovements
		d    = fmt.Sprintf("%d %s %d", date.Day(), date.Month().String(), date.Year())
		uri  = fmt.Sprintf(
			"%s/getTrainMovementsXML?TrainId=%s&TrainDate=%s", irBaseURL, id, d)
	)
	err := i.makeRequest(uri, resp)
	return resp.TrainMovements, err
}

// GetCurrentTrains will retirive the TrainPoisitons for all currently operating trains
func (i Irishrail) GetCurrentTrains() ([]TrainPositions, error) {
	var (
		resp arrayOfObjTrainPositions
		uri  = fmt.Sprintf("%s/getCurrentTrainsXML", irBaseURL)
	)
	err := i.makeRequest(uri, resp)
	return resp.TrainPositions, err
}

// GetCurrentTrainsByRoute will retirive the TrainPoisitons for all currently
// operating trains on a given route
func (i Irishrail) GetCurrentTrainsByRoute(r string) ([]TrainPositions, error) {
	v, ok := routeMapping[r]
	if !ok {
		return nil, errors.New("current trains request error: route not found")
	}

	var (
		resp arrayOfObjTrainPositions
		uri  = fmt.Sprintf("%s/getCurrentTrainsXMLL_WithTrainType?TrainType=%s", irBaseURL, v)
	)
	err := i.makeRequest(uri, resp)
	return resp.TrainPositions, err
}
