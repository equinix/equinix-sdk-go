# RouterActionsConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Connection UUID | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ConnectionType**](ConnectionType.md) |  | [optional] 
**Operation** | Pointer to [**Operation**](Operation.md) |  | [optional] 

## Methods

### NewRouterActionsConnection

`func NewRouterActionsConnection() *RouterActionsConnection`

NewRouterActionsConnection instantiates a new RouterActionsConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterActionsConnectionWithDefaults

`func NewRouterActionsConnectionWithDefaults() *RouterActionsConnection`

NewRouterActionsConnectionWithDefaults instantiates a new RouterActionsConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RouterActionsConnection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouterActionsConnection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouterActionsConnection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouterActionsConnection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHref

`func (o *RouterActionsConnection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouterActionsConnection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouterActionsConnection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouterActionsConnection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouterActionsConnection) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouterActionsConnection) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouterActionsConnection) SetType(v ConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *RouterActionsConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOperation

`func (o *RouterActionsConnection) GetOperation() Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RouterActionsConnection) GetOperationOk() (*Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RouterActionsConnection) SetOperation(v Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RouterActionsConnection) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


