# EnvironmentDataForArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNo** | Pointer to **string** | account number | [optional] 
**Cabinet** | Pointer to **string** | cabinet unique space id | [optional] 
**Cage** | Pointer to **string** | cage unique space id | [optional] 
**Humidity** | Pointer to **string** | current humidity | [optional] 
**HumidityUom** | Pointer to **string** | unit of measure for humidity | [optional] 
**Ibx** | Pointer to **string** | ibx code | [optional] 
**Sensor** | Pointer to **string** | sensor id | [optional] 
**Temperature** | Pointer to **string** | current temperature | [optional] 
**TemperatureUom** | Pointer to **string** | unit of measure for temperature values | [optional] 
**Timestamp** | Pointer to **string** | epoch timestamp when the current reading was read | [optional] 
**Zone** | Pointer to **string** | zone unique space id | [optional] 

## Methods

### NewEnvironmentDataForArray

`func NewEnvironmentDataForArray() *EnvironmentDataForArray`

NewEnvironmentDataForArray instantiates a new EnvironmentDataForArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDataForArrayWithDefaults

`func NewEnvironmentDataForArrayWithDefaults() *EnvironmentDataForArray`

NewEnvironmentDataForArrayWithDefaults instantiates a new EnvironmentDataForArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *EnvironmentDataForArray) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *EnvironmentDataForArray) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *EnvironmentDataForArray) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *EnvironmentDataForArray) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetCabinet

`func (o *EnvironmentDataForArray) GetCabinet() string`

GetCabinet returns the Cabinet field if non-nil, zero value otherwise.

### GetCabinetOk

`func (o *EnvironmentDataForArray) GetCabinetOk() (*string, bool)`

GetCabinetOk returns a tuple with the Cabinet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinet

`func (o *EnvironmentDataForArray) SetCabinet(v string)`

SetCabinet sets Cabinet field to given value.

### HasCabinet

`func (o *EnvironmentDataForArray) HasCabinet() bool`

HasCabinet returns a boolean if a field has been set.

### GetCage

`func (o *EnvironmentDataForArray) GetCage() string`

GetCage returns the Cage field if non-nil, zero value otherwise.

### GetCageOk

`func (o *EnvironmentDataForArray) GetCageOk() (*string, bool)`

GetCageOk returns a tuple with the Cage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCage

`func (o *EnvironmentDataForArray) SetCage(v string)`

SetCage sets Cage field to given value.

### HasCage

`func (o *EnvironmentDataForArray) HasCage() bool`

HasCage returns a boolean if a field has been set.

### GetHumidity

`func (o *EnvironmentDataForArray) GetHumidity() string`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *EnvironmentDataForArray) GetHumidityOk() (*string, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *EnvironmentDataForArray) SetHumidity(v string)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *EnvironmentDataForArray) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetHumidityUom

`func (o *EnvironmentDataForArray) GetHumidityUom() string`

GetHumidityUom returns the HumidityUom field if non-nil, zero value otherwise.

### GetHumidityUomOk

`func (o *EnvironmentDataForArray) GetHumidityUomOk() (*string, bool)`

GetHumidityUomOk returns a tuple with the HumidityUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidityUom

`func (o *EnvironmentDataForArray) SetHumidityUom(v string)`

SetHumidityUom sets HumidityUom field to given value.

### HasHumidityUom

`func (o *EnvironmentDataForArray) HasHumidityUom() bool`

HasHumidityUom returns a boolean if a field has been set.

### GetIbx

`func (o *EnvironmentDataForArray) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *EnvironmentDataForArray) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *EnvironmentDataForArray) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *EnvironmentDataForArray) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetSensor

`func (o *EnvironmentDataForArray) GetSensor() string`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *EnvironmentDataForArray) GetSensorOk() (*string, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *EnvironmentDataForArray) SetSensor(v string)`

SetSensor sets Sensor field to given value.

### HasSensor

`func (o *EnvironmentDataForArray) HasSensor() bool`

HasSensor returns a boolean if a field has been set.

### GetTemperature

`func (o *EnvironmentDataForArray) GetTemperature() string`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *EnvironmentDataForArray) GetTemperatureOk() (*string, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *EnvironmentDataForArray) SetTemperature(v string)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *EnvironmentDataForArray) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTemperatureUom

`func (o *EnvironmentDataForArray) GetTemperatureUom() string`

GetTemperatureUom returns the TemperatureUom field if non-nil, zero value otherwise.

### GetTemperatureUomOk

`func (o *EnvironmentDataForArray) GetTemperatureUomOk() (*string, bool)`

GetTemperatureUomOk returns a tuple with the TemperatureUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureUom

`func (o *EnvironmentDataForArray) SetTemperatureUom(v string)`

SetTemperatureUom sets TemperatureUom field to given value.

### HasTemperatureUom

`func (o *EnvironmentDataForArray) HasTemperatureUom() bool`

HasTemperatureUom returns a boolean if a field has been set.

### GetTimestamp

`func (o *EnvironmentDataForArray) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EnvironmentDataForArray) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EnvironmentDataForArray) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EnvironmentDataForArray) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetZone

`func (o *EnvironmentDataForArray) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *EnvironmentDataForArray) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *EnvironmentDataForArray) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *EnvironmentDataForArray) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


