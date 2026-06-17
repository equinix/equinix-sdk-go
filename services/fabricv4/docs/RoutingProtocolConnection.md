# RoutingProtocolConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Equinix-assigned connection identifier | [optional] 
**PlatformUuid** | Pointer to **string** | Equinix-assigned platform connection identifier | [optional] 

## Methods

### NewRoutingProtocolConnection

`func NewRoutingProtocolConnection() *RoutingProtocolConnection`

NewRoutingProtocolConnection instantiates a new RoutingProtocolConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolConnectionWithDefaults

`func NewRoutingProtocolConnectionWithDefaults() *RoutingProtocolConnection`

NewRoutingProtocolConnectionWithDefaults instantiates a new RoutingProtocolConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RoutingProtocolConnection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolConnection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolConnection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolConnection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetPlatformUuid

`func (o *RoutingProtocolConnection) GetPlatformUuid() string`

GetPlatformUuid returns the PlatformUuid field if non-nil, zero value otherwise.

### GetPlatformUuidOk

`func (o *RoutingProtocolConnection) GetPlatformUuidOk() (*string, bool)`

GetPlatformUuidOk returns a tuple with the PlatformUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUuid

`func (o *RoutingProtocolConnection) SetPlatformUuid(v string)`

SetPlatformUuid sets PlatformUuid field to given value.

### HasPlatformUuid

`func (o *RoutingProtocolConnection) HasPlatformUuid() bool`

HasPlatformUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


