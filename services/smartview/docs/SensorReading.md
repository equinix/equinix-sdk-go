# SensorReading

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Humidity** | Pointer to [**ValueWithUnit**](ValueWithUnit.md) |  | [optional] 
**Ibx** | **string** | The ibx identifier where the sensor is placed. | 
**SensorId** | **string** | The sensor identifier. | 
**Temperature** | Pointer to [**ValueWithUnit**](ValueWithUnit.md) |  | [optional] 
**ZoneId** | **string** | The zone where the sensor is placed. | 

## Methods

### NewSensorReading

`func NewSensorReading(ibx string, sensorId string, zoneId string, ) *SensorReading`

NewSensorReading instantiates a new SensorReading object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSensorReadingWithDefaults

`func NewSensorReadingWithDefaults() *SensorReading`

NewSensorReadingWithDefaults instantiates a new SensorReading object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHumidity

`func (o *SensorReading) GetHumidity() ValueWithUnit`

GetHumidity returns the Humidity field if non-nil, zero value otherwise.

### GetHumidityOk

`func (o *SensorReading) GetHumidityOk() (*ValueWithUnit, bool)`

GetHumidityOk returns a tuple with the Humidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidity

`func (o *SensorReading) SetHumidity(v ValueWithUnit)`

SetHumidity sets Humidity field to given value.

### HasHumidity

`func (o *SensorReading) HasHumidity() bool`

HasHumidity returns a boolean if a field has been set.

### GetIbx

`func (o *SensorReading) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *SensorReading) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *SensorReading) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetSensorId

`func (o *SensorReading) GetSensorId() string`

GetSensorId returns the SensorId field if non-nil, zero value otherwise.

### GetSensorIdOk

`func (o *SensorReading) GetSensorIdOk() (*string, bool)`

GetSensorIdOk returns a tuple with the SensorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorId

`func (o *SensorReading) SetSensorId(v string)`

SetSensorId sets SensorId field to given value.


### GetTemperature

`func (o *SensorReading) GetTemperature() ValueWithUnit`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *SensorReading) GetTemperatureOk() (*ValueWithUnit, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *SensorReading) SetTemperature(v ValueWithUnit)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *SensorReading) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetZoneId

`func (o *SensorReading) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *SensorReading) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *SensorReading) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


