# RoutingProtocolChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RoutingProtocolChangeType**](RoutingProtocolChangeType.md) |  | 
**Href** | Pointer to **string** | Routing Protocol Change URI | [optional] 

## Methods

### NewRoutingProtocolChange

`func NewRoutingProtocolChange(uuid string, type_ RoutingProtocolChangeType, ) *RoutingProtocolChange`

NewRoutingProtocolChange instantiates a new RoutingProtocolChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolChangeWithDefaults

`func NewRoutingProtocolChangeWithDefaults() *RoutingProtocolChange`

NewRoutingProtocolChangeWithDefaults instantiates a new RoutingProtocolChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RoutingProtocolChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RoutingProtocolChange) GetType() RoutingProtocolChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolChange) GetTypeOk() (*RoutingProtocolChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolChange) SetType(v RoutingProtocolChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RoutingProtocolChange) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolChange) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolChange) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolChange) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


