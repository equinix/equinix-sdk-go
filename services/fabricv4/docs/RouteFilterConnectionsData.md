# RouteFilterConnectionsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Connection URI | [optional] 
**Type** | Pointer to [**ConnectionType**](ConnectionType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route filter identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewRouteFilterConnectionsData

`func NewRouteFilterConnectionsData() *RouteFilterConnectionsData`

NewRouteFilterConnectionsData instantiates a new RouteFilterConnectionsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterConnectionsDataWithDefaults

`func NewRouteFilterConnectionsDataWithDefaults() *RouteFilterConnectionsData`

NewRouteFilterConnectionsDataWithDefaults instantiates a new RouteFilterConnectionsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RouteFilterConnectionsData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteFilterConnectionsData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteFilterConnectionsData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteFilterConnectionsData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouteFilterConnectionsData) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFilterConnectionsData) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFilterConnectionsData) SetType(v ConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *RouteFilterConnectionsData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RouteFilterConnectionsData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteFilterConnectionsData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteFilterConnectionsData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouteFilterConnectionsData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RouteFilterConnectionsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteFilterConnectionsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteFilterConnectionsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteFilterConnectionsData) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


