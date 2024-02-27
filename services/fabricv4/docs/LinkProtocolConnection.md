# LinkProtocolConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Connection URI | [optional] [readonly] 
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Bandwidth** | Pointer to **int64** |  | [optional] 

## Methods

### NewLinkProtocolConnection

`func NewLinkProtocolConnection() *LinkProtocolConnection`

NewLinkProtocolConnection instantiates a new LinkProtocolConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolConnectionWithDefaults

`func NewLinkProtocolConnectionWithDefaults() *LinkProtocolConnection`

NewLinkProtocolConnectionWithDefaults instantiates a new LinkProtocolConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LinkProtocolConnection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LinkProtocolConnection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LinkProtocolConnection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LinkProtocolConnection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *LinkProtocolConnection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LinkProtocolConnection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LinkProtocolConnection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LinkProtocolConnection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *LinkProtocolConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkProtocolConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkProtocolConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkProtocolConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBandwidth

`func (o *LinkProtocolConnection) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *LinkProtocolConnection) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *LinkProtocolConnection) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *LinkProtocolConnection) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


