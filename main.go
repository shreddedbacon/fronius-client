package main

import (
	"flag"
	"fmt"

	"github.com/shreddedbacon/fronius-client/fronius"
)

/*
	This is an example usage of the fronius client
*/
func main() {
	var inverter string
	flag.StringVar(&inverter, "inverter", "http://192.168.1.50", "URL or IP for the host")
	flag.Parse()
	f, _ := fronius.New(inverter)

	// Get the realtime data from the metering device 0
	m, _ := f.GetMeterRealtimeDataDevice(0)
	fmt.Println(m)

	// Get the realtime data from the interver system
	i, _ := f.GetInverterRealtimeDataSystem()
	fmt.Println(i)

	// Get the realtime powerflow data
	p, _ := f.GetPowerFlowRealtimeData()
	fmt.Println(p)

	// Just run any request against the inverter
	d, _ := f.Request("/solar_api/v1/GetActiveDeviceInfo.cgi?DeviceClass=Inverter")
	fmt.Println(string(d))

}
