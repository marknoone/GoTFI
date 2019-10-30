package irishrail

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

func (Irishrail) makeRequest(uri string, dest interface{}) error {
	reqURL, err := url.Parse(uri)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return xml.NewDecoder(resp.Body).Decode(dest)
	default:
		return fmt.Errorf(
			"request error occured: recieved http status code (%d)",
			resp.StatusCode)
	}
}
