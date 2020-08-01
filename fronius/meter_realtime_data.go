package fronius

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetMeterRealtimeDataDevice gets the realtime meter data from the inverter.
func (f *Fronius) GetMeterRealtimeDataDevice(deviceID int) (*MeterRealtimeDataDevice, error) {
	data := &MeterRealtimeDataDevice{}
	resp, err := f.Request(fmt.Sprintf("%sGetMeterRealtimeData.cgi?Scope=Device&DeviceId=%d", f.hostAPI, deviceID))
	if err != nil {
		return data, err
	}
	json.Unmarshal(resp, &data)
	return data, nil
}

// MeterRealtimeDataDevice http://192.168.1.50/solar_api/v1/GetMeterRealtimeData.cgi?Scope=Device&DeviceId=0
type MeterRealtimeDataDevice struct {
	Body struct {
		Data struct {
			CurrentACPhase1 float64 `json:"Current_AC_Phase_1"`
			CurrentACPhase2 float64 `json:"Current_AC_Phase_2"`
			CurrentACPhase3 float64 `json:"Current_AC_Phase_3"`
			Details         struct {
				Manufacturer string `json:"Manufacturer"`
				Model        string `json:"Model"`
				Serial       string `json:"Serial"`
			} `json:"Details"`
			Enable                         int     `json:"Enable"`
			EnergyReactiveVArACSumConsumed int     `json:"EnergyReactive_VArAC_Sum_Consumed"`
			EnergyReactiveVArACSumProduced int     `json:"EnergyReactive_VArAC_Sum_Produced"`
			EnergyRealWACMinusAbsolute     int     `json:"EnergyReal_WAC_Minus_Absolute"`
			EnergyRealWACPlusAbsolute      int     `json:"EnergyReal_WAC_Plus_Absolute"`
			EnergyRealWACSumConsumed       int     `json:"EnergyReal_WAC_Sum_Consumed"`
			EnergyRealWACSumProduced       int     `json:"EnergyReal_WAC_Sum_Produced"`
			FrequencyPhaseAverage          int     `json:"Frequency_Phase_Average"`
			MeterLocationCurrent           int     `json:"Meter_Location_Current"`
			PowerApparentSPhase1           float64 `json:"PowerApparent_S_Phase_1"`
			PowerApparentSPhase2           float64 `json:"PowerApparent_S_Phase_2"`
			PowerApparentSPhase3           float64 `json:"PowerApparent_S_Phase_3"`
			PowerApparentSSum              float64 `json:"PowerApparent_S_Sum"`
			PowerFactorPhase1              float64 `json:"PowerFactor_Phase_1"`
			PowerFactorPhase2              int     `json:"PowerFactor_Phase_2"`
			PowerFactorPhase3              int     `json:"PowerFactor_Phase_3"`
			PowerFactorSum                 int     `json:"PowerFactor_Sum"`
			PowerReactiveQPhase1           float64 `json:"PowerReactive_Q_Phase_1"`
			PowerReactiveQPhase2           int     `json:"PowerReactive_Q_Phase_2"`
			PowerReactiveQPhase3           float64 `json:"PowerReactive_Q_Phase_3"`
			PowerReactiveQSum              float64 `json:"PowerReactive_Q_Sum"`
			PowerRealPPhase1               int     `json:"PowerReal_P_Phase_1"`
			PowerRealPPhase2               float64 `json:"PowerReal_P_Phase_2"`
			PowerRealPPhase3               float64 `json:"PowerReal_P_Phase_3"`
			PowerRealPSum                  float64 `json:"PowerReal_P_Sum"`
			TimeStamp                      int     `json:"TimeStamp"`
			Visible                        int     `json:"Visible"`
			VoltageACPhaseToPhase12        float64 `json:"Voltage_AC_PhaseToPhase_12"`
			VoltageACPhaseToPhase23        float64 `json:"Voltage_AC_PhaseToPhase_23"`
			VoltageACPhaseToPhase31        float64 `json:"Voltage_AC_PhaseToPhase_31"`
			VoltageACPhase1                float64 `json:"Voltage_AC_Phase_1"`
			VoltageACPhase2                float64 `json:"Voltage_AC_Phase_2"`
			VoltageACPhase3                float64 `json:"Voltage_AC_Phase_3"`
		} `json:"Data"`
	} `json:"Body"`
	Head struct {
		RequestArguments struct {
			DeviceClass string `json:"DeviceClass"`
			DeviceID    string `json:"DeviceId"`
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
