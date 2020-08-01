package fronius

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Fronius .
type Fronius struct {
	netClient *http.Client
	Host      string
	hostAPI   string
}

// APIVersion is the api version of the fronius inverter.
// @TODO: actually do something if the version changes, currently only have seen v1 anywhere.
type APIVersion struct {
	APIVersion         int    `json:"APIVersion"`
	BaseURL            string `json:"BaseURL"`
	CompatibilityRange string `json:"CompatibilityRange"`
}

// New creates a new fronius api client.
func New(host string) (*Fronius, error) {
	netClientTimeout := 10
	var netClient = &http.Client{
		Timeout: time.Second * time.Duration(netClientTimeout),
	}
	client := &Fronius{
		netClient: netClient,
		Host:      host,
	}
	// get the API baseURL
	api, err := client.GetAPIVersion()
	if err != nil {
		return client, err
	}
	// add it to the client
	client.hostAPI = api.BaseURL
	return client, nil
}

// Request creates and executes the http request to the API.
func (f *Fronius) Request(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", f.Host, endpoint), bytes.NewBuffer([]byte{}))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.netClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	rBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return []byte{}, errors.New("error performing check or connecting to inverter")
	}
	return rBody, nil
}

// GetAPIVersion gets the api version of the fronius inverter.
func (f *Fronius) GetAPIVersion() (*APIVersion, error) {
	data := &APIVersion{}
	resp, err := f.Request("/solar_api/GetAPIVersion.cgi")
	if err != nil {
		return data, err
	}
	json.Unmarshal(resp, &data)
	return data, nil
}
