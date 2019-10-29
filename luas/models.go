package luas

import (
	"encoding/xml"
)

type luasForcastingResponse struct {
	XMLName   xml.Name    `xml:"stopInfo"`
	CreatedAt string      `xml:"created,attr"`
	Message   string      `xml:"message"`
	Direction []direction `xml:"direction"`
}

type direction struct {
	XMLName   xml.Name `xml:"direction"`
	Direction string   `xml:"name,attr"`
	Tram      []tram   `xml:"tram"`
}

type tram struct {
	Due         string `xml:"dueMins,attr"`
	Destination string `xml:"destination,attr"`
}
