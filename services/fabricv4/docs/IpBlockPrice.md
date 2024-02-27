# IpBlockPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Either uuid or rest of attributes are required | [optional] 
**Type** | Pointer to [**IpBlockType**](IpBlockType.md) |  | [optional] 
**PrefixLength** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to [**PriceLocation**](PriceLocation.md) |  | [optional] 

## Methods

### NewIpBlockPrice

`func NewIpBlockPrice() *IpBlockPrice`

NewIpBlockPrice instantiates a new IpBlockPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockPriceWithDefaults

`func NewIpBlockPriceWithDefaults() *IpBlockPrice`

NewIpBlockPriceWithDefaults instantiates a new IpBlockPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *IpBlockPrice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IpBlockPrice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IpBlockPrice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IpBlockPrice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *IpBlockPrice) GetType() IpBlockType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpBlockPrice) GetTypeOk() (*IpBlockType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpBlockPrice) SetType(v IpBlockType)`

SetType sets Type field to given value.

### HasType

`func (o *IpBlockPrice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrefixLength

`func (o *IpBlockPrice) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *IpBlockPrice) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *IpBlockPrice) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *IpBlockPrice) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetLocation

`func (o *IpBlockPrice) GetLocation() PriceLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IpBlockPrice) GetLocationOk() (*PriceLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IpBlockPrice) SetLocation(v PriceLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IpBlockPrice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


