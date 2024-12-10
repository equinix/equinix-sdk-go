# RouteAggregationConnectionsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Connection URI | [optional] 
**Type** | Pointer to [**ConnectionType**](ConnectionType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route Aggregation identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewRouteAggregationConnectionsData

`func NewRouteAggregationConnectionsData() *RouteAggregationConnectionsData`

NewRouteAggregationConnectionsData instantiates a new RouteAggregationConnectionsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationConnectionsDataWithDefaults

`func NewRouteAggregationConnectionsDataWithDefaults() *RouteAggregationConnectionsData`

NewRouteAggregationConnectionsDataWithDefaults instantiates a new RouteAggregationConnectionsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RouteAggregationConnectionsData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteAggregationConnectionsData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteAggregationConnectionsData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteAggregationConnectionsData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouteAggregationConnectionsData) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationConnectionsData) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationConnectionsData) SetType(v ConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *RouteAggregationConnectionsData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RouteAggregationConnectionsData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteAggregationConnectionsData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteAggregationConnectionsData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouteAggregationConnectionsData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RouteAggregationConnectionsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteAggregationConnectionsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteAggregationConnectionsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteAggregationConnectionsData) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


