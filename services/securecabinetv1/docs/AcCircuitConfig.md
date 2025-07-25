# AcCircuitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Voltage** | **int32** | Voltage numerical value | 
**SoldAmperage** | **float64** | Amperage Numerical Value | 
**Phase** | [**Phase**](Phase.md) |  | 
**Receptacle** | **string** | Receptacle type | 

## Methods

### NewAcCircuitConfig

`func NewAcCircuitConfig(voltage int32, soldAmperage float64, phase Phase, receptacle string, ) *AcCircuitConfig`

NewAcCircuitConfig instantiates a new AcCircuitConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcCircuitConfigWithDefaults

`func NewAcCircuitConfigWithDefaults() *AcCircuitConfig`

NewAcCircuitConfigWithDefaults instantiates a new AcCircuitConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoltage

`func (o *AcCircuitConfig) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *AcCircuitConfig) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *AcCircuitConfig) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.


### GetSoldAmperage

`func (o *AcCircuitConfig) GetSoldAmperage() float64`

GetSoldAmperage returns the SoldAmperage field if non-nil, zero value otherwise.

### GetSoldAmperageOk

`func (o *AcCircuitConfig) GetSoldAmperageOk() (*float64, bool)`

GetSoldAmperageOk returns a tuple with the SoldAmperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAmperage

`func (o *AcCircuitConfig) SetSoldAmperage(v float64)`

SetSoldAmperage sets SoldAmperage field to given value.


### GetPhase

`func (o *AcCircuitConfig) GetPhase() Phase`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *AcCircuitConfig) GetPhaseOk() (*Phase, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *AcCircuitConfig) SetPhase(v Phase)`

SetPhase sets Phase field to given value.


### GetReceptacle

`func (o *AcCircuitConfig) GetReceptacle() string`

GetReceptacle returns the Receptacle field if non-nil, zero value otherwise.

### GetReceptacleOk

`func (o *AcCircuitConfig) GetReceptacleOk() (*string, bool)`

GetReceptacleOk returns a tuple with the Receptacle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceptacle

`func (o *AcCircuitConfig) SetReceptacle(v string)`

SetReceptacle sets Receptacle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


