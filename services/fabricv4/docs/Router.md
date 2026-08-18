# Router

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Fabric Cloud Router URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Cloud Router UUID | [optional] 
**Type** | Pointer to [**CloudRouterPostRequestBaseType**](CloudRouterPostRequestBaseType.md) |  | [optional] 

## Methods

### NewRouter

`func NewRouter() *Router`

NewRouter instantiates a new Router object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterWithDefaults

`func NewRouterWithDefaults() *Router`

NewRouterWithDefaults instantiates a new Router object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Router) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Router) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Router) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Router) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Router) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Router) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Router) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Router) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *Router) GetType() CloudRouterPostRequestBaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Router) GetTypeOk() (*CloudRouterPostRequestBaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Router) SetType(v CloudRouterPostRequestBaseType)`

SetType sets Type field to given value.

### HasType

`func (o *Router) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


