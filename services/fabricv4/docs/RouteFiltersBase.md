# RouteFiltersBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**RouteFiltersBaseType**](RouteFiltersBaseType.md) |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** | Customer-provided connection description | [optional] 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewRouteFiltersBase

`func NewRouteFiltersBase(type_ RouteFiltersBaseType, name string, project Project, ) *RouteFiltersBase`

NewRouteFiltersBase instantiates a new RouteFiltersBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersBaseWithDefaults

`func NewRouteFiltersBaseWithDefaults() *RouteFiltersBase`

NewRouteFiltersBaseWithDefaults instantiates a new RouteFiltersBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RouteFiltersBase) GetType() RouteFiltersBaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFiltersBase) GetTypeOk() (*RouteFiltersBaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFiltersBase) SetType(v RouteFiltersBaseType)`

SetType sets Type field to given value.


### GetName

`func (o *RouteFiltersBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteFiltersBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteFiltersBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RouteFiltersBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteFiltersBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteFiltersBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteFiltersBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *RouteFiltersBase) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RouteFiltersBase) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RouteFiltersBase) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


