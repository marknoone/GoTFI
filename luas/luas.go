package luas

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	tfi "github.com/marknoone/GoTFI"
)

const luasBaseURL = "http://luasforecasts.rpa.ie/xml/get.ashx?action=forecast&encrypt=false"

var _ tfi.Operator = Luas{}

type Luas struct{}

func (Luas) GetStopRTPI(id string) (tfi.RTResult, error) {
	_, ok := luasStops[id]
	if !ok {
		return tfi.RTResult{}, fmt.Errorf("luas stop (%s) not supported", id)
	}

	reqURL, err := url.Parse(fmt.Sprintf("%s&stop=%s", luasBaseURL, id))
	if err != nil {
		return tfi.RTResult{}, err
	}

	resp, err := http.Get(reqURL.String())
	if err != nil {
		return tfi.RTResult{}, err
	}
	defer resp.Body.Close()

	var xmlResp luasForcastingResponse
	err = xml.NewDecoder(resp.Body).Decode(&xmlResp)
	if err != nil {
		return tfi.RTResult{}, err
	}

	var results []tfi.RTResultRow
	for _, d := range xmlResp.Direction {
		for _, t := range d.Tram {
			var r string
			if s, ok := luasStops[t.Destination]; ok {
				r = s.Routes[0]
			}

			i, err := strconv.Atoi(t.Due)
			if err != nil {
				return tfi.RTResult{}, err
			}

			results = append(results, tfi.RTResultRow{
				Route:       r,
				Destination: t.Destination,
				Due:         i,
				Direction:   d.Direction,
			})
		}
	}

	return tfi.RTResult{
		Timestamp: xmlResp.CreatedAt,
		Results:   results}, nil
}

func (Luas) GetRoutes() (map[string]tfi.Route, error) {
	var res = make(map[string]tfi.Route, 2)
	for k, v := range luasStops {
		val, ok := res[v.Routes[0]]
		if !ok {
			val.ID = v.Routes[0]
		}
		val.StopIDs = append(val.StopIDs, k)
		res[v.Routes[0]] = val
	}
	return res, nil
}

func (Luas) GetStops() (map[string]tfi.Stop, error) {
	return luasStops, nil
}

func (Luas) GetStop(id string) (tfi.Stop, error) {
	s, ok := luasStops[id]
	if !ok {
		return tfi.Stop{}, errors.New("Stop not found")
	}

	return s, nil
}

func (Luas) GetRoute(id string) (tfi.Route, error) {
	var res = tfi.Route{
		ID:      id,
		StopIDs: []string{},
	}

	for k, v := range luasStops {
		if v.Routes[0] != id {
			continue
		}
		res.StopIDs = append(res.StopIDs, k)
	}
	return res, nil
}
