package smartmeapi

import (
	// "bytes"
	"encoding/base64"
	json2 "encoding/json"
	"fmt"
	"time"

	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

type ChargeStationState int32

const (
	Booting             ChargeStationState = 0
	ReadyNoCarConnected ChargeStationState = 1
	ReadyCarConnected   ChargeStationState = 2
	StartedWaitForCar   ChargeStationState = 3
	Charging            ChargeStationState = 4
	Installation        ChargeStationState = 5
	Authorize           ChargeStationState = 6
	Offline             ChargeStationState = 7
)

type Device struct {
	Id                          *string             `json:"id,omitempty"`
	Name                        *string             `json:"name,omitempty"`
	Serial                      *int64              `json:"serial,omitempty"`
	DeviceEnergyType            *MeterEnergyType    `json:"deviceEnergyType,omitempty"`
	MeterSubType                *MeterSubType       `json:"meterSubType,omitempty"`
	FamilyType                  *MeterFamilyType    `json:"familyType,omitempty"`
	ActivePower                 *float64            `json:"activePower,omitempty"`
	ActivePowerL1               *float64            `json:"activePowerL1,omitempty"`
	ActivePowerL2               *float64            `json:"activePowerL2,omitempty"`
	ActivePowerL3               *float64            `json:"activePowerL3,omitempty"`
	ActivePowerUnit             *string             `json:"activePowerUnit,omitempty"`
	CounterReading              *float64            `json:"counterReading,omitempty"`
	CounterReadingUnit          *string             `json:"counterReadingUnit,omitempty"`
	CounterReadingT1            *float64            `json:"counterReadingT1,omitempty"`
	CounterReadingT2            *float64            `json:"counterReadingT2,omitempty"`
	CounterReadingT3            *float64            `json:"counterReadingT3,omitempty"`
	CounterReadingT4            *float64            `json:"counterReadingT4,omitempty"`
	CounterReadingImport        *float64            `json:"counterReadingImport,omitempty"`
	CounterReadingExport        *float64            `json:"counterReadingExport,omitempty"`
	SwitchOn                    *bool               `json:"switchOn,omitempty"`
	SwitchPhaseL10n             *bool               `json:"switchPhaseL10n,omitempty"`
	SwitchPhaseL20n             *bool               `json:"switchPhaseL20n,omitempty"`
	SwitchPhaseL30n             *bool               `json:"switchPhaseL30n,omitempty"`
	Voltage                     *float64            `json:"voltage,omitempty"`
	VoltageL1                   *float64            `json:"voltageL1,omitempty"`
	VoltageL2                   *float64            `json:"voltageL2,omitempty"`
	VoltageL3                   *float64            `json:"voltageL3,omitempty"`
	Current                     *float64            `json:"current,omitempty"`
	CurrentL1                   *float64            `json:"currentL1,omitempty"`
	CurrentL2                   *float64            `json:"currentL2,omitempty"`
	CurrentL3                   *float64            `json:"currentL3,omitempty"`
	PowerFactor                 *float64            `json:"powerFactor,omitempty"`
	PowerFactorL1               *float64            `json:"powerFactorL1,omitempty"`
	PowerFactorL2               *float64            `json:"powerFactorL2,omitempty"`
	PowerFactorL3               *float64            `json:"powerFactorL3,omitempty"`
	Temperature                 *float64            `json:"temperature,omitempty"`
	ActiveTariff                *int32              `json:"activeTariff,omitempty"`
	DigitalOutput1              *bool               `json:"digitalOutput1,omitempty"`
	DigitalOutput2              *bool               `json:"digitalOutput2,omitempty"`
	AnalogOutput1               *int32              `json:"analogOutput1,omitempty"`
	AnalogOutput2               *int32              `json:"analogOutput2,omitempty"`
	DigitalInput1               *bool               `json:"digitalInput1,omitempty"`
	DigitalInput2               *bool               `json:"digitalInput2,omitempty"`
	ValueDate                   *string             `json:"valueDate,omitempty"`
	AdditionalMeterSerialNumber *string             `json:"additionalMeterSerialNumber,omitempty"`
	FlowRate                    *float64            `json:"flowRate,omitempty"`
	ChargeStationState          *ChargeStationState `json:"chargeStationState"`
}

type Devices []Device

type ValueData struct {
	Obis  *string  `json:"obis,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

type ValuesData struct {
	DeviceId *string `json:"deviceId,omitempty"`
	Date     *string `json:"date,omitempty"`
	Values   []ValueData
}

// ===============================================================================================

var logger *log.Logger
var logLevel int

type apiConfiguration struct {
	Url            *url.URL
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
	var err error
	apiConfig.Url, err = url.Parse(host)
	if err != nil {
		log.Fatal(err)
	}
	apiConfig.Authentication = "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	logger = loggerParam
	logLevel = logLevelParam
}

func GetDevices() (*[]Device, error) {
	httpUrl := apiConfig.Url.JoinPath("Devices")
	fmt.Println(httpUrl)
	json, err := loadUrl(httpUrl.String())
	if err != nil {
		return nil, err
	}
	var result []Device
	err = json2.Unmarshal(json, &result)
	return &result, err
}

func GetDevice(deviceId string) (*Device, error) {
	httpUrl := apiConfig.Url.JoinPath("Devices", deviceId)
	fmt.Println(httpUrl)
	json, err := loadUrl(httpUrl.String())
	if err != nil {
		return nil, err
	}
	var result Device
	err = json2.Unmarshal(json, &result)
	return &result, err
}

func GetValues(deviceId string) (*ValuesData, error) {
	httpUrl := apiConfig.Url.JoinPath("Values", deviceId)
	fmt.Println(httpUrl)
	json, err := loadUrl(httpUrl.String())
	if err != nil {
		return nil, err
	}
	var result ValuesData
	err = json2.Unmarshal(json, &result)
	return &result, err
}

func GetValuesInPast(deviceId string, date time.Time) (*ValuesData, error) {
	httpUrl := apiConfig.Url.JoinPath("ValuesInPast", deviceId)
	values := url.Values{}
	values.Add("date", date.Format((time.RFC3339)))
	httpUrl.RawQuery = values.Encode()
	fmt.Println(httpUrl)
	json, err := loadUrl(httpUrl.String())
	if err != nil {
		return nil, err
	}
	var result ValuesData
	err = json2.Unmarshal(json, &result)
	return &result, err
}

func GetValuesInPastMultiple(deviceId string, startDate time.Time, endDate time.Time, interval int) (*[]ValuesData, error) {
	httpUrl := apiConfig.Url.JoinPath("ValuesInPastMultiple", deviceId)
	values := url.Values{}
	values.Add("startDate", startDate.Format((time.RFC3339)))
	values.Add("endDate", endDate.Format((time.RFC3339)))
	values.Add("interval", strconv.Itoa(interval))
	httpUrl.RawQuery = values.Encode()
	fmt.Println(httpUrl)
	json, err := loadUrl(httpUrl.String())
	if err != nil {
		return nil, err
	}
	var result []ValuesData
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

	json, err = io.ReadAll(response.Body)

	return json, err
}
