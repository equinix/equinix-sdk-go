# ValidateConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Equinix-assigned connection identifier | [optional] 
**Bandwidth** | Pointer to **int32** | Connection bandwidth in Mbps | [optional] 
**Redundancy** | Pointer to [**ConnectionRedundancy**](ConnectionRedundancy.md) |  | [optional] 
**ASide** | Pointer to [**ConnectionSide**](ConnectionSide.md) |  | [optional] 
**ZSide** | Pointer to [**ConnectionSide**](ConnectionSide.md) |  | [optional] 

## Methods

### NewValidateConnectionResponse

`func NewValidateConnectionResponse() *ValidateConnectionResponse`

NewValidateConnectionResponse instantiates a new ValidateConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateConnectionResponseWithDefaults

`func NewValidateConnectionResponseWithDefaults() *ValidateConnectionResponse`

NewValidateConnectionResponseWithDefaults instantiates a new ValidateConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ValidateConnectionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ValidateConnectionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ValidateConnectionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ValidateConnectionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetBandwidth

`func (o *ValidateConnectionResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ValidateConnectionResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ValidateConnectionResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ValidateConnectionResponse) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetRedundancy

`func (o *ValidateConnectionResponse) GetRedundancy() ConnectionRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *ValidateConnectionResponse) GetRedundancyOk() (*ConnectionRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *ValidateConnectionResponse) SetRedundancy(v ConnectionRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *ValidateConnectionResponse) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetASide

`func (o *ValidateConnectionResponse) GetASide() ConnectionSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *ValidateConnectionResponse) GetASideOk() (*ConnectionSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *ValidateConnectionResponse) SetASide(v ConnectionSide)`

SetASide sets ASide field to given value.

### HasASide

`func (o *ValidateConnectionResponse) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *ValidateConnectionResponse) GetZSide() ConnectionSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *ValidateConnectionResponse) GetZSideOk() (*ConnectionSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *ValidateConnectionResponse) SetZSide(v ConnectionSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *ValidateConnectionResponse) HasZSide() bool`

HasZSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


