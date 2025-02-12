# SpotMarketRequestCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicesMax** | Pointer to **int32** |  | [optional] 
**DevicesMin** | Pointer to **int32** |  | [optional] 
**EndAt** | Pointer to **time.Time** |  | [optional] 
**InstanceParameters** | Pointer to [**SpotMarketRequestCreateInputInstanceParameters**](SpotMarketRequestCreateInputInstanceParameters.md) |  | [optional] 
**MaxBidPrice** | Pointer to **float32** |  | [optional] 
**Metro** | Pointer to **string** | The metro ID or code the spot market request will be created in. | [optional] 

## Methods

### NewSpotMarketRequestCreateInput

`func NewSpotMarketRequestCreateInput() *SpotMarketRequestCreateInput`

NewSpotMarketRequestCreateInput instantiates a new SpotMarketRequestCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpotMarketRequestCreateInputWithDefaults

`func NewSpotMarketRequestCreateInputWithDefaults() *SpotMarketRequestCreateInput`

NewSpotMarketRequestCreateInputWithDefaults instantiates a new SpotMarketRequestCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicesMax

`func (o *SpotMarketRequestCreateInput) GetDevicesMax() int32`

GetDevicesMax returns the DevicesMax field if non-nil, zero value otherwise.

### GetDevicesMaxOk

`func (o *SpotMarketRequestCreateInput) GetDevicesMaxOk() (*int32, bool)`

GetDevicesMaxOk returns a tuple with the DevicesMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMax

`func (o *SpotMarketRequestCreateInput) SetDevicesMax(v int32)`

SetDevicesMax sets DevicesMax field to given value.

### HasDevicesMax

`func (o *SpotMarketRequestCreateInput) HasDevicesMax() bool`

HasDevicesMax returns a boolean if a field has been set.

### GetDevicesMin

`func (o *SpotMarketRequestCreateInput) GetDevicesMin() int32`

GetDevicesMin returns the DevicesMin field if non-nil, zero value otherwise.

### GetDevicesMinOk

`func (o *SpotMarketRequestCreateInput) GetDevicesMinOk() (*int32, bool)`

GetDevicesMinOk returns a tuple with the DevicesMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesMin

`func (o *SpotMarketRequestCreateInput) SetDevicesMin(v int32)`

SetDevicesMin sets DevicesMin field to given value.

### HasDevicesMin

`func (o *SpotMarketRequestCreateInput) HasDevicesMin() bool`

HasDevicesMin returns a boolean if a field has been set.

### GetEndAt

`func (o *SpotMarketRequestCreateInput) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *SpotMarketRequestCreateInput) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *SpotMarketRequestCreateInput) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *SpotMarketRequestCreateInput) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetInstanceParameters

`func (o *SpotMarketRequestCreateInput) GetInstanceParameters() SpotMarketRequestCreateInputInstanceParameters`

GetInstanceParameters returns the InstanceParameters field if non-nil, zero value otherwise.

### GetInstanceParametersOk

`func (o *SpotMarketRequestCreateInput) GetInstanceParametersOk() (*SpotMarketRequestCreateInputInstanceParameters, bool)`

GetInstanceParametersOk returns a tuple with the InstanceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceParameters

`func (o *SpotMarketRequestCreateInput) SetInstanceParameters(v SpotMarketRequestCreateInputInstanceParameters)`

SetInstanceParameters sets InstanceParameters field to given value.

### HasInstanceParameters

`func (o *SpotMarketRequestCreateInput) HasInstanceParameters() bool`

HasInstanceParameters returns a boolean if a field has been set.

### GetMaxBidPrice

`func (o *SpotMarketRequestCreateInput) GetMaxBidPrice() float32`

GetMaxBidPrice returns the MaxBidPrice field if non-nil, zero value otherwise.

### GetMaxBidPriceOk

`func (o *SpotMarketRequestCreateInput) GetMaxBidPriceOk() (*float32, bool)`

GetMaxBidPriceOk returns a tuple with the MaxBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBidPrice

`func (o *SpotMarketRequestCreateInput) SetMaxBidPrice(v float32)`

SetMaxBidPrice sets MaxBidPrice field to given value.

### HasMaxBidPrice

`func (o *SpotMarketRequestCreateInput) HasMaxBidPrice() bool`

HasMaxBidPrice returns a boolean if a field has been set.

### GetMetro

`func (o *SpotMarketRequestCreateInput) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *SpotMarketRequestCreateInput) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *SpotMarketRequestCreateInput) SetMetro(v string)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *SpotMarketRequestCreateInput) HasMetro() bool`

HasMetro returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


