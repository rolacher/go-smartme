package smartmeapi

import (
	// "bytes"
	"encoding/base64"
	json2 "encoding/json"
	"fmt"

	// "io"
	"io/ioutil"
	"log"
	"net/http"
)

type MeterEnergyType int32

// see https://api.smart-me.com/swagger/ for the definitions
const (
	MeterTypeUnknown       MeterEnergyType = 0
	MeterTypeElectricity   MeterEnergyType = 1
	MeterTypeWater         MeterEnergyType = 2
	MeterTypeGas           MeterEnergyType = 3
	MeterTypeHeat          MeterEnergyType = 4
	MeterTypeHCA           MeterEnergyType = 5
	MeterTypeAllMeters     MeterEnergyType = 6
	MeterTypeTemperature   MeterEnergyType = 7
	MeterTypeMBusGateway   MeterEnergyType = 8
	MeterTypeRS485Gateway  MeterEnergyType = 9
	MeterTypeCustomDevice  MeterEnergyType = 10
	MeterTypeCompressedAir MeterEnergyType = 11
	MeterTypeSolarLog      MeterEnergyType = 12
	MeterTypeVirtualMeter  MeterEnergyType = 13
	MeterTypeWMBusGateway  MeterEnergyType = 14
)

type MeterSubType int32

const (
	MeterSubTypeUnknown         MeterSubType = 0
	Cold_ColdWaterMeter         MeterSubType = 1
	Heat_HotWaterMeter          MeterSubType = 2
	MeterSubTypeChargingStation MeterSubType = 3
	MeterSubTypeElectricity     MeterSubType = 4
	MeterSubTypeWater           MeterSubType = 5
	MeterSubTypeGas             MeterSubType = 6
	Electricity_HeatMeter       MeterSubType = 7
	TemperatureMeter            MeterSubType = 8
	MeterSubTypeVirtualBattery  MeterSubType = 9
)

type MeterFamilyType int32

const (
	The_Family_Type_is_unknown_all_M_BUS_Meters_S0_meters_usw     MeterFamilyType = 0
	smart_me_connect_Meter_Plugin_Power_Meter                     MeterFamilyType = 1
	smart_me_Meter_1_Phase_DIN_Rail_Meter_without_switch          MeterFamilyType = 2
	smart_me_Meter_1_Phase_DIN_Rail_Meter_with_a_Switch           MeterFamilyType = 3
	smart_me_M_BUS_Gateway_V1                                     MeterFamilyType = 4
	smart_me_RS_485_Gateway_V1                                    MeterFamilyType = 5
	MeterFamilyTypeKamstrupModule                                 MeterFamilyType = 6
	MeterFamilyTypeSmartMe3PhaseMeter80A                          MeterFamilyType = 7
	smart_me_3_Phase_Meter_32A_with_Switch                        MeterFamilyType = 8
	smart_me_3_Phase_Meter_Transformer_Edition                    MeterFamilyType = 9
	smart_me_Landis_Gyr_Module                                    MeterFamilyType = 10
	Optical_module_for_the_FNN_meters                             MeterFamilyType = 11
	smart_me_3_Phase_Meter_80A_with_the_new_WiFi_V2               MeterFamilyType = 12
	smart_me_3_Phase_Meter_80A_with_Mobile                        MeterFamilyType = 14
	smart_me_1_Phase_Meter_80A_with_the_new_WiFi_V2               MeterFamilyType = 16
	smart_me_1_Phase_Meter_32A_with_the_new_WiFi_V2               MeterFamilyType = 17
	smart_me_1_Phase_Meter_80A_with_GPRS                          MeterFamilyType = 18
	smart_me_1_Phase_Meter_32A_with_GPRS                          MeterFamilyType = 19
	smart_me_Wirless_M_BUS_Gateway_V1                             MeterFamilyType = 20
	smart_me_3_Phase_Meter_Transformer_Edition_with_mobile_module MeterFamilyType = 21
	smart_me_3_phase_Meter_Nimbus_3_point_meter                   MeterFamilyType = 65
	Mithral_hall_charging_station_Version_1                       MeterFamilyType = 70
	REST_API_Meter                                                MeterFamilyType = 1001
	Virtual_billing_Meter                                         MeterFamilyType = 1002
)

type Device struct {
	Id                   *string          `json:"id,omitempty"`
	Name                 *string          `json:"name,omitempty"`
	Serial               *int64           `json:"serial,omitempty"`
	DeviceEnergyType     *MeterEnergyType `json:"deviceEnergyType,omitempty"`
	MeterSubType         *MeterSubType    `json:"meterSubType,omitempty"`
	FamilyType           *MeterFamilyType `json:"familyType,omitempty"`
	ActivePower          *float32         `json:"activePower,omitempty"`
	ActivePowerL1        *float32         `json:"activePowerL1,omitempty"`
	ActivePowerL2        *float32         `json:"activePowerL2,omitempty"`
	ActivePowerL3        *float32         `json:"activePowerL3,omitempty"`
	ActivePowerUnit      *string          `json:"activePowerUnit,omitempty"`
	CounterReading       *float32         `json:"counterReading,omitempty"`
	CounterReadingUnit   *string          `json:"counterReadingUnit,omitempty"`
	CounterReadingT1     *float32         `json:"counterReadingT1,omitempty"`
	CounterReadingT2     *float32         `json:"counterReadingT2,omitempty"`
	CounterReadingT3     *float32         `json:"counterReadingT3,omitempty"`
	CounterReadingT4     *float32         `json:"counterReadingT4,omitempty"`
	CounterReadingImport *float32         `json:"counterReadingImport,omitempty"`
	CounterReadingExport *float32         `json:"counterReadingExport,omitempty"`
	SwitchOn             *bool            `json:"switchOn,omitempty"`
	SwitchPhaseL10n      *bool            `json:"switchPhaseL10n,omitempty"`
	SwitchPhaseL20n      *bool            `json:"switchPhaseL20n,omitempty"`
	SwitchPhaseL30n      *bool            `json:"switchPhaseL30n,omitempty"`
	Voltage              *float32         `json:"voltage,omitempty"`
	VoltageL1            *float32         `json:"voltageL1,omitempty"`
	VoltageL2            *float32         `json:"voltageL2,omitempty"`
	VoltageL3            *float32         `json:"voltageL3,omitempty"`
	Current              *float32         `json:"current,omitempty"`
	CurrentL1            *float32         `json:"currentL1,omitempty"`
	CurrentL2            *float32         `json:"currentL2,omitempty"`
	CurrentL3            *float32         `json:"currentL3,omitempty"`
}

type Devices []Device

// ===============================================================================================

var logger *log.Logger
var logLevel int

type apiConfiguration struct {
	Host           string
	Authentication string
}

var apiConfig = apiConfiguration{}

func ConfigureApi(
	host string,
	username string,
	password string,
	loggerParam *log.Logger,
	logLevelParam int,
) {
	apiConfig.Host = host
	apiConfig.Authentication = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	logger = loggerParam
	logLevel = logLevelParam
}

func GetDevices() (*Devices, error) {
	httpUrl := fmt.Sprintf("%s%s", apiConfig.Host, "/Devices")
	fmt.Println(httpUrl)
	json, err := loadUrl(httpUrl)
	if err != nil {
		return nil, err
	}
	// var result Devices
	var result Devices
	err = json2.Unmarshal(json, &result)
	return &result, err
}

func GetDevice(deviceId string) (*Device, error) {
	httpUrl := fmt.Sprintf("%s%s%s", apiConfig.Host, "/Devices/", deviceId)
	fmt.Println(httpUrl)
	json, err := loadUrl(httpUrl)
	if err != nil {
		return nil, err
	}
	var result Device
	err = json2.Unmarshal(json, &result)
	return &result, err
}

func loadUrl(httpUrl string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", httpUrl, nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", apiConfig.Authentication)

	if logLevel > 1 {
		logger.Printf("getting %s ...\n", httpUrl)
	}

	var json []byte

	response, err := client.Do(req)
	if err != nil {
		logger.Printf("error getting %s: %s\n", httpUrl, err.Error())
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("GET url %s returned code %d (%s)", httpUrl, response.StatusCode, response.Status)
	}

	json, err = ioutil.ReadAll(response.Body)

	return json, err
}
