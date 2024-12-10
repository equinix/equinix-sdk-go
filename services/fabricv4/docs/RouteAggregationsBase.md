# RouteAggregationsBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RouteAggregationsBaseType**](RouteAggregationsBaseType.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** | Customer-provided connection description | [optional] 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewRouteAggregationsBase

`func NewRouteAggregationsBase(type_ RouteAggregationsBaseType, name string, project Project, ) *RouteAggregationsBase`

NewRouteAggregationsBase instantiates a new RouteAggregationsBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationsBaseWithDefaults

`func NewRouteAggregationsBaseWithDefaults() *RouteAggregationsBase`

NewRouteAggregationsBaseWithDefaults instantiates a new RouteAggregationsBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RouteAggregationsBase) GetType() RouteAggregationsBaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationsBase) GetTypeOk() (*RouteAggregationsBaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationsBase) SetType(v RouteAggregationsBaseType)`

SetType sets Type field to given value.


### GetName

`func (o *RouteAggregationsBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteAggregationsBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteAggregationsBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RouteAggregationsBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteAggregationsBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteAggregationsBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteAggregationsBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *RouteAggregationsBase) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RouteAggregationsBase) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RouteAggregationsBase) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


