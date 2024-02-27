# SimplifiedNetworkChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Network URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | Pointer to [**NetworkChangeType**](NetworkChangeType.md) |  | [optional] 

## Methods

### NewSimplifiedNetworkChange

`func NewSimplifiedNetworkChange() *SimplifiedNetworkChange`

NewSimplifiedNetworkChange instantiates a new SimplifiedNetworkChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedNetworkChangeWithDefaults

`func NewSimplifiedNetworkChangeWithDefaults() *SimplifiedNetworkChange`

NewSimplifiedNetworkChangeWithDefaults instantiates a new SimplifiedNetworkChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SimplifiedNetworkChange) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedNetworkChange) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedNetworkChange) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedNetworkChange) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *SimplifiedNetworkChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SimplifiedNetworkChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SimplifiedNetworkChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SimplifiedNetworkChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedNetworkChange) GetType() NetworkChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedNetworkChange) GetTypeOk() (*NetworkChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedNetworkChange) SetType(v NetworkChangeType)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedNetworkChange) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


