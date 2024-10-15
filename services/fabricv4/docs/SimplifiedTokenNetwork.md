# SimplifiedTokenNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | url to entity | [optional] 
**Uuid** | Pointer to **string** | Network Identifier | [optional] 
**Type** | Pointer to [**SimplifiedTokenNetworkType**](SimplifiedTokenNetworkType.md) |  | [optional] 
**Name** | Pointer to **string** | Network Name | [optional] 
**Scope** | Pointer to [**SimplifiedTokenNetworkScope**](SimplifiedTokenNetworkScope.md) |  | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 

## Methods

### NewSimplifiedTokenNetwork

`func NewSimplifiedTokenNetwork() *SimplifiedTokenNetwork`

NewSimplifiedTokenNetwork instantiates a new SimplifiedTokenNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedTokenNetworkWithDefaults

`func NewSimplifiedTokenNetworkWithDefaults() *SimplifiedTokenNetwork`

NewSimplifiedTokenNetworkWithDefaults instantiates a new SimplifiedTokenNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SimplifiedTokenNetwork) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedTokenNetwork) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedTokenNetwork) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedTokenNetwork) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *SimplifiedTokenNetwork) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SimplifiedTokenNetwork) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SimplifiedTokenNetwork) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SimplifiedTokenNetwork) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedTokenNetwork) GetType() SimplifiedTokenNetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedTokenNetwork) GetTypeOk() (*SimplifiedTokenNetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedTokenNetwork) SetType(v SimplifiedTokenNetworkType)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedTokenNetwork) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedTokenNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedTokenNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedTokenNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedTokenNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *SimplifiedTokenNetwork) GetScope() SimplifiedTokenNetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SimplifiedTokenNetwork) GetScopeOk() (*SimplifiedTokenNetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SimplifiedTokenNetwork) SetScope(v SimplifiedTokenNetworkScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SimplifiedTokenNetwork) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetLocation

`func (o *SimplifiedTokenNetwork) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SimplifiedTokenNetwork) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SimplifiedTokenNetwork) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SimplifiedTokenNetwork) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


