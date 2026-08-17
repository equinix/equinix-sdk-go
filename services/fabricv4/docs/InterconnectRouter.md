# InterconnectRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Router URI | [optional] 
**Type** | Pointer to [**InterconnectRouterType**](InterconnectRouterType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Router identifier | [optional] 

## Methods

### NewInterconnectRouter

`func NewInterconnectRouter() *InterconnectRouter`

NewInterconnectRouter instantiates a new InterconnectRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectRouterWithDefaults

`func NewInterconnectRouterWithDefaults() *InterconnectRouter`

NewInterconnectRouterWithDefaults instantiates a new InterconnectRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *InterconnectRouter) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InterconnectRouter) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InterconnectRouter) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *InterconnectRouter) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *InterconnectRouter) GetType() InterconnectRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterconnectRouter) GetTypeOk() (*InterconnectRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterconnectRouter) SetType(v InterconnectRouterType)`

SetType sets Type field to given value.

### HasType

`func (o *InterconnectRouter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *InterconnectRouter) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InterconnectRouter) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InterconnectRouter) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InterconnectRouter) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


