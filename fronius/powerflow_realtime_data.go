package fronius

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetPowerFlowRealtimeData gets the realtime powerflow data from the inverter.
func (f *Fronius) GetPowerFlowRealtimeData() (*PowerFlowRealtimeData, error) {
	data := &PowerFlowRealtimeData{}
	resp, err := f.Request(fmt.Sprintf("%sGetPowerFlowRealtimeData.fcgi", f.hostAPI))
	if err != nil {
		return data, err
	}
	json.Unmarshal(resp, &data)
	return data, nil
}

// PowerFlowRealtimeData http://192.168.1.50/solar_api/v1/GetPowerFlowRealtimeData.fcgi
type PowerFlowRealtimeData struct {
	Body struct {
		Data struct {
			Inverters struct {
				Num1 struct {
					DT     int     `json:"DT"`
					EDay   float64 `json:"E_Day"`
					ETotal int     `json:"E_Total"`
					EYear  int     `json:"E_Year"`
					P      int     `json:"P"`
				} `json:"1"`
			} `json:"Inverters"`
			Site struct {
				EDay               float64     `json:"E_Day"`
				ETotal             int         `json:"E_Total"`
				EYear              int         `json:"E_Year"`
				MeterLocation      string      `json:"Meter_Location"`
				Mode               string      `json:"Mode"`
				PAkku              interface{} `json:"P_Akku"`
				PGrid              float64     `json:"P_Grid"`
				PLoad              float64     `json:"P_Load"`
				PPV                float64     `json:"P_PV"`
				RelAutonomy        float64     `json:"rel_Autonomy"`
				RelSelfConsumption int         `json:"rel_SelfConsumption"`
			} `json:"Site"`
			Version string `json:"Version"`
		} `json:"Data"`
	} `json:"Body"`
	Head struct {
		RequestArguments struct {
		} `json:"RequestArguments"`
		Status struct {
			Code        int    `json:"Code"`
			Reason      string `json:"Reason"`
			UserMessage string `json:"UserMessage"`
		} `json:"Status"`
		Timestamp time.Time `json:"Timestamp"`
	} `json:"Head"`
}
