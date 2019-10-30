package irishrail

import "encoding/xml"

type arrayOfObjStationData struct {
	XMLName     xml.Name      `xml:"ArrayOfObjStationData"`
	StationData []StationData `xml:"objStationData"`
}

// StationData contains all of a stations real-time data
// retrieved from the irishrail real-time service
type StationData struct {
	XMLName           xml.Name `xml:"objStationData"`
	ServerTime        string   `xml:"Servertime"`
	TrainCode         string   `xml:"Traincode"`
	StationFullName   string   `xml:"Stationfullname"`
	StationCode       string   `xml:"Stationcode"`
	QueryTime         string   `xml:"Querytime"`
	TrainDate         string   `xml:"Traindate"`
	Origin            string   `xml:"Origin"`
	Destination       string   `xml:"Destination"`
	OriginTime        string   `xml:"Origintime"`
	DestinationTime   string   `xml:"Destinationtime"`
	Status            string   `xml:"Status"`
	LastLocation      string   `xml:"Lastlocation"`
	DueIn             string   `xml:"Duein"`
	Late              string   `xml:"Late"`
	ExpectedArrival   string   `xml:"Exparrival"`
	ExpectedDeparture string   `xml:"Expdepart"`
	ScheduledArrival  string   `xml:"Scharrival"`
	ScheduledDepart   string   `xml:"Schdepart"`
	Direction         string   `xml:"Direction"`
	TrainType         string   `xml:"Traintype"`
	LocationType      string   `xml:"Locationtype"`
}

type arrayOfObjStation struct {
	XMLName xml.Name  `xml:"ArrayOfObjStation"`
	Station []Station `xml:"objStation"`
}

// Station contains all of a stations static
type Station struct {
	XMLName     xml.Name `xml:"objStation"`
	Description string   `xml:"StationDesc"`
	Alias       string   `xml:"StationAlias"`
	Latitude    string   `xml:"StationLatitude"`
	Longitude   string   `xml:"StationLongitude"`
	Code        string   `xml:"StationCode"`
	ID          string   `xml:"StationId"`
}

type arrayOfObjTrainPositions struct {
	XMLName        xml.Name         `xml:"ArrayOfObjStationFilter"`
	TrainPositions []TrainPositions `xml:"objTrainPositions"`
}

// TrainPositions contains a trains current position on the network
type TrainPositions struct {
	XMLName       xml.Name `xml:"objTrainPositions"`
	Status        string   `xml:"TrainStatus"`
	Latitude      string   `xml:"TrainLatitude"`
	Longitude     string   `xml:"TrainLongitude"`
	Code          string   `xml:"TrainCode"`
	Date          string   `xml:"TrainDate"`
	PublicMessage string   `xml:"PublicMessage"`
	Direction     string   `xml:"Direction"`
}

type arrayOfObjStationFilter struct {
	XMLName        xml.Name        `xml:"ArrayOfObjStationFilter"`
	StationFilters []StationFilter `xml:"objStationFilter"`
}

// StationFilter provides a single result from Irishrail's station
// filter API
type StationFilter struct {
	XMLName       xml.Name `xml:"objStationFilter"`
	DescriptionSP string   `xml:"StationDesc_sp"`
	Description   string   `xml:"StationDesc"`
	Code          string   `xml:"StationCode"`
}

type arrayOfObjTrainMovements struct {
	XMLName        xml.Name         `xml:"ArrayOfObjTrainMovements"`
	TrainMovements []TrainMovements `xml:"objTrainMovements"`
}

// TrainMovements is the data container for a single result of
// Irishrail's train movement API
type TrainMovements struct {
	XMLName            xml.Name `xml:"objStationFilter"`
	TrainCode          string   `xml:"TrainCode"`
	TrainDate          string   `xml:"TrainDate"`
	LocationCode       string   `xml:"LocationCode"`
	LocationFullName   string   `xml:"LocationFullName"`
	LocationOrder      string   `xml:"LocationOrder"`
	LocationType       string   `xml:"LocationType"`
	TrainOrigin        string   `xml:"TrainOrigin"`
	TrainDestination   string   `xml:"TrainDestination"`
	ScheduledArrival   string   `xml:"ScheduledArrival"`
	ScheduledDeparture string   `xml:"ScheduledDeparture"`
	ExpectedArrival    string   `xml:"ExpectedArrival"`
	ExpectedDeparture  string   `xml:"ExpectedDeparture"`
	Arrival            string   `xml:"Arrival"`
	Departure          string   `xml:"Departure"`
	AutoArrival        string   `xml:"AutoArrival"`
	AutoDepart         string   `xml:"AutoDepart"`
	StopType           string   `xml:"StopType"`
}
