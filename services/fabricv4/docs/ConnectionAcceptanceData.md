# ConnectionAcceptanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZSide** | Pointer to [**ConnectionSide**](ConnectionSide.md) |  | [optional] 
**ProviderBandwidth** | Pointer to **int32** | Authorization key bandwidth in Mbps | [optional] [readonly] 

## Methods

### NewConnectionAcceptanceData

`func NewConnectionAcceptanceData() *ConnectionAcceptanceData`

NewConnectionAcceptanceData instantiates a new ConnectionAcceptanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionAcceptanceDataWithDefaults

`func NewConnectionAcceptanceDataWithDefaults() *ConnectionAcceptanceData`

NewConnectionAcceptanceDataWithDefaults instantiates a new ConnectionAcceptanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZSide

`func (o *ConnectionAcceptanceData) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *ConnectionAcceptanceData) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *ConnectionAcceptanceData) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *ConnectionAcceptanceData) HasZSide() bool`

HasZSide returns a boolean if a field has been set.

### GetProviderBandwidth

`func (o *ConnectionAcceptanceData) GetProviderBandwidth() int32`

GetProviderBandwidth returns the ProviderBandwidth field if non-nil, zero value otherwise.

### GetProviderBandwidthOk

`func (o *ConnectionAcceptanceData) GetProviderBandwidthOk() (*int32, bool)`

GetProviderBandwidthOk returns a tuple with the ProviderBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBandwidth

`func (o *ConnectionAcceptanceData) SetProviderBandwidth(v int32)`

SetProviderBandwidth sets ProviderBandwidth field to given value.

### HasProviderBandwidth

`func (o *ConnectionAcceptanceData) HasProviderBandwidth() bool`

HasProviderBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


