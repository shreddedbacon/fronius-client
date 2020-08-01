package fronius

import (
	"encoding/json"
	"fmt"
	"time"
)


// GetInverterRealtimeDataSystem gets the realtime data from the inverter
func (f *Fronius) GetInverterRealtimeDataSystem() (*InverterRealtimeDataSystem, error) {
	data := &InverterRealtimeDataSystem{}
	resp, err := f.Request(fmt.Sprintf("%sGetInverterRealtimeData.cgi?Scope=System", f.hostAPI))
	if err != nil {
		return data, err
	}
	json.Unmarshal(resp, &data)
	return data, nil
}

// InverterRealtimeDataSystem http://192.168.1.50/solar_api/v1/GetInverterRealtimeData.cgi?Scope=System
type InverterRealtimeDataSystem struct {
	Body struct {
		Data struct {
			DAYENERGY struct {
				Unit   string `json:"Unit"`
				Values struct {
					Num1 int `json:"1"`
				} `json:"Values"`
			} `json:"DAY_ENERGY"`
			PAC struct {
				Unit   string `json:"Unit"`
				Values struct {
					Num1 int `json:"1"`
				} `json:"Values"`
			} `json:"PAC"`
			TOTALENERGY struct {
				Unit   string `json:"Unit"`
				Values struct {
					Num1 int `json:"1"`
				} `json:"Values"`
			} `json:"TOTAL_ENERGY"`
			YEARENERGY struct {
				Unit   string `json:"Unit"`
				Values struct {
					Num1 int `json:"1"`
				} `json:"Values"`
			} `json:"YEAR_ENERGY"`
		} `json:"Data"`
	} `json:"Body"`
	Head struct {
		RequestArguments struct {
			DeviceClass string `json:"DeviceClass"`
			Scope       string `json:"Scope"`
		} `json:"RequestArguments"`
		Status struct {
			Code        int    `json:"Code"`
			Reason      string `json:"Reason"`
			UserMessage string `json:"UserMessage"`
		} `json:"Status"`
		Timestamp time.Time `json:"Timestamp"`
	} `json:"Head"`
}