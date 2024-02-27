# RouteFilterRulesChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RouteFilterRulesChangeType**](RouteFilterRulesChangeType.md) |  | 
**Href** | Pointer to **string** | Route Filter Change URI | [optional] 

## Methods

### NewRouteFilterRulesChange

`func NewRouteFilterRulesChange(uuid string, type_ RouteFilterRulesChangeType, ) *RouteFilterRulesChange`

NewRouteFilterRulesChange instantiates a new RouteFilterRulesChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRulesChangeWithDefaults

`func NewRouteFilterRulesChangeWithDefaults() *RouteFilterRulesChange`

NewRouteFilterRulesChangeWithDefaults instantiates a new RouteFilterRulesChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RouteFilterRulesChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteFilterRulesChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteFilterRulesChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RouteFilterRulesChange) GetType() RouteFilterRulesChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFilterRulesChange) GetTypeOk() (*RouteFilterRulesChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFilterRulesChange) SetType(v RouteFilterRulesChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteFilterRulesChange) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteFilterRulesChange) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteFilterRulesChange) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteFilterRulesChange) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


