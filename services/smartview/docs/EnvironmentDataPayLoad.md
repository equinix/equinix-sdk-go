# EnvironmentDataPayLoad

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

### NewEnvironmentDataPayLoad

`func NewEnvironmentDataPayLoad() *EnvironmentDataPayLoad`

NewEnvironmentDataPayLoad instantiates a new EnvironmentDataPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDataPayLoadWithDefaults

`func NewEnvironmentDataPayLoadWithDefaults() *EnvironmentDataPayLoad`

NewEnvironmentDataPayLoadWithDefaults instantiates a new EnvironmentDataPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNo

`func (o *EnvironmentDataPayLoad) GetAccountNo() string`

GetAccountNo returns the AccountNo field if non-nil, zero value otherwise.

### GetAccountNoOk

`func (o *EnvironmentDataPayLoad) GetAccountNoOk() (*string, bool)`

GetAccountNoOk returns a tuple with the AccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNo

`func (o *EnvironmentDataPayLoad) SetAccountNo(v string)`

SetAccountNo sets AccountNo field to given value.

### HasAccountNo

`func (o *EnvironmentDataPayLoad) HasAccountNo() bool`

HasAccountNo returns a boolean if a field has been set.

### GetCabinet

`func (o *EnvironmentDataPayLoad) GetCabinet() string`

GetCabinet returns the Cabinet field if non-nil, zero value otherwise.

### GetCabinetOk

`func (o *EnvironmentDataPayLoad) GetCabinetOk() (*string, bool)`

GetCabinetOk returns a tuple with the Cabinet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinet

`func (o *EnvironmentDataPayLoad) SetCabinet(v string)`

SetCabinet sets Cabinet field to given value.

### HasCabinet

`func (o *EnvironmentDataPayLoad) HasCabinet() bool`

HasCabinet returns a boolean if a field has been set.

### GetCage

`func (o *EnvironmentDataPayLoad) GetCage() string`

GetCage returns the Cage field if non-nil, zero value otherwise.

### GetCageOk

`func (o *EnvironmentDataPayLoad) GetCageOk() (*string, bool)`

GetCageOk returns a tuple with the Cage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCage

`func (o *EnvironmentDataPayLoad) SetCage(v string)`

SetCage sets Cage field to given value.

### HasCage

`func (o *EnvironmentDataPayLoad) HasCage() bool`

HasCage returns a boolean if a field has been set.

### GetHumidity

`func (o *EnvironmentDataPayLoad) GetHumidity() string`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *EnvironmentDataPayLoad) GetHumidityOk() (*string, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *EnvironmentDataPayLoad) SetHumidity(v string)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *EnvironmentDataPayLoad) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetHumidityUom

`func (o *EnvironmentDataPayLoad) GetHumidityUom() string`

GetHumidityUom returns the HumidityUom field if non-nil, zero value otherwise.

### GetHumidityUomOk

`func (o *EnvironmentDataPayLoad) GetHumidityUomOk() (*string, bool)`

GetHumidityUomOk returns a tuple with the HumidityUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidityUom

`func (o *EnvironmentDataPayLoad) SetHumidityUom(v string)`

SetHumidityUom sets HumidityUom field to given value.

### HasHumidityUom

`func (o *EnvironmentDataPayLoad) HasHumidityUom() bool`

HasHumidityUom returns a boolean if a field has been set.

### GetIbx

`func (o *EnvironmentDataPayLoad) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *EnvironmentDataPayLoad) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *EnvironmentDataPayLoad) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *EnvironmentDataPayLoad) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetSensor

`func (o *EnvironmentDataPayLoad) GetSensor() string`

GetSensor returns the Sensor field if non-nil, zero value otherwise.

### GetSensorOk

`func (o *EnvironmentDataPayLoad) GetSensorOk() (*string, bool)`

GetSensorOk returns a tuple with the Sensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensor

`func (o *EnvironmentDataPayLoad) SetSensor(v string)`

SetSensor sets Sensor field to given value.

### HasSensor

`func (o *EnvironmentDataPayLoad) HasSensor() bool`

HasSensor returns a boolean if a field has been set.

### GetTemperature

`func (o *EnvironmentDataPayLoad) GetTemperature() string`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *EnvironmentDataPayLoad) GetTemperatureOk() (*string, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *EnvironmentDataPayLoad) SetTemperature(v string)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *EnvironmentDataPayLoad) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTemperatureUom

`func (o *EnvironmentDataPayLoad) GetTemperatureUom() string`

GetTemperatureUom returns the TemperatureUom field if non-nil, zero value otherwise.

### GetTemperatureUomOk

`func (o *EnvironmentDataPayLoad) GetTemperatureUomOk() (*string, bool)`

GetTemperatureUomOk returns a tuple with the TemperatureUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureUom

`func (o *EnvironmentDataPayLoad) SetTemperatureUom(v string)`

SetTemperatureUom sets TemperatureUom field to given value.

### HasTemperatureUom

`func (o *EnvironmentDataPayLoad) HasTemperatureUom() bool`

HasTemperatureUom returns a boolean if a field has been set.

### GetTimestamp

`func (o *EnvironmentDataPayLoad) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EnvironmentDataPayLoad) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EnvironmentDataPayLoad) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EnvironmentDataPayLoad) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetZone

`func (o *EnvironmentDataPayLoad) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *EnvironmentDataPayLoad) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *EnvironmentDataPayLoad) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *EnvironmentDataPayLoad) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


