# RouterActionsRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Router UUID | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**CloudRouterPostRequestBaseType**](CloudRouterPostRequestBaseType.md) |  | [optional] 
**Operation** | Pointer to [**Operation**](Operation.md) |  | [optional] 

## Methods

### NewRouterActionsRouter

`func NewRouterActionsRouter() *RouterActionsRouter`

NewRouterActionsRouter instantiates a new RouterActionsRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterActionsRouterWithDefaults

`func NewRouterActionsRouterWithDefaults() *RouterActionsRouter`

NewRouterActionsRouterWithDefaults instantiates a new RouterActionsRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RouterActionsRouter) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouterActionsRouter) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouterActionsRouter) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouterActionsRouter) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHref

`func (o *RouterActionsRouter) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouterActionsRouter) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouterActionsRouter) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouterActionsRouter) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouterActionsRouter) GetType() CloudRouterPostRequestBaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouterActionsRouter) GetTypeOk() (*CloudRouterPostRequestBaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouterActionsRouter) SetType(v CloudRouterPostRequestBaseType)`

SetType sets Type field to given value.

### HasType

`func (o *RouterActionsRouter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOperation

`func (o *RouterActionsRouter) GetOperation() Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RouterActionsRouter) GetOperationOk() (*Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RouterActionsRouter) SetOperation(v Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RouterActionsRouter) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


