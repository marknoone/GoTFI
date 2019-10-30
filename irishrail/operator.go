package irishrail

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	tfi "github.com/marknoone/GoTFI"
)

// Operator Interface Funcs
func (i Irishrail) GetStop(id string) (tfi.Stop, error) {
	stops, err := i.GetStops()
	if err != nil {
		return tfi.Stop{}, err
	}

	val, ok := stops[id]
	if !ok {
		return tfi.Stop{}, errors.New("stop not found")
	}

	return val, nil
}

func (i Irishrail) GetStops() (map[string]tfi.Stop, error) {
	var (
		resp arrayOfObjStation

		// Default from current total count of 267 stops
		stopMap = make(map[string]tfi.Stop, 280)
		uri     = fmt.Sprintf("%s/getAllStationsXML", irBaseURL)
	)

	parseFloat := func(s string) float32 {
		f, err := strconv.ParseFloat(s, 32)
		if err != nil {
			log.Println("float conversion error: number " + s)
		}
		return float32(f)
	}

	err := i.makeRequest(uri, resp)
	if err != nil {
		return nil, err
	}

	for _, s := range resp.Station {
		stopMap[s.Code] = tfi.Stop{
			ID:   s.Code,
			Name: s.Description,
			Lat:  parseFloat(s.Latitude),
			Lng:  parseFloat(s.Longitude),
		}
	}

	return stopMap, nil

}

func (i Irishrail) GetRoute(id string) (tfi.Route, error) {
	if _, ok := routeMapping[id]; !ok {
		return tfi.Route{}, errors.New("route not found")
	}

	var (
		resp  arrayOfObjStation
		route tfi.Route
		uri   = fmt.Sprintf("%s/getAllStationsXML_WithStationType?StationType=%s",
			irBaseURL, routeMapping[id])
	)

	err := i.makeRequest(uri, resp)
	if err != nil {
		return tfi.Route{}, err
	}

	var r = tfi.Route{ID: id, StopIDs: []string{}}
	for _, s := range resp.Station {
		r.StopIDs = append(r.StopIDs, s.Code)
	}

	return route, nil
}

func (i Irishrail) GetRoutes() (map[string]tfi.Route, error) {

	getRoute := func(t string) (tfi.Route, error) {
		var respObj arrayOfObjStation
		err := i.makeRequest(
			fmt.Sprintf("%s/getAllStationsXML_WithStationType?StationType=%s", irBaseURL, t),
			respObj)
		if err != nil {
			return tfi.Route{}, err
		}

		var r = tfi.Route{ID: typeMapping[t], StopIDs: []string{}}
		for _, s := range respObj.Station {
			r.StopIDs = append(r.StopIDs, s.Code)
		}
		return r, nil
	}

	dr, err := getRoute("D")
	if err != nil {
		return nil, err
	}

	sr, err := getRoute("S")
	if err != nil {
		return nil, err
	}

	mr, err := getRoute("M")
	if err != nil {
		return nil, err
	}

	return map[string]tfi.Route{
		typeMapping["D"]: dr, typeMapping["S"]: sr, typeMapping["M"]: mr}, nil
}

func (i Irishrail) GetStopRTPI(id string) (tfi.RTResult, error) {
	var (
		resp    arrayOfObjStationData
		results = []tfi.RTResultRow{}
		uri     = fmt.Sprintf("%s/getStationDataByCodeXML?StationCode=%s", irBaseURL, id)
	)

	parseInt := func(s string) int {
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			log.Println("float conversion error: number " + s)
		}
		return int(i)
	}

	err := i.makeRequest(uri, resp)
	if err != nil {
		return tfi.RTResult{}, err
	}

	for _, sd := range resp.StationData {
		results = append(results, tfi.RTResultRow{
			Destination: sd.Destination,
			Due:         parseInt(sd.DueIn),
			Direction:   sd.Direction,
		})
	}

	return tfi.RTResult{
		Timestamp: resp.StationData[0].ServerTime,
		Results:   results}, nil
}
