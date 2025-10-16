# ConnectionServicesDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**ConnectionServices**](ConnectionServices.md) |  | [optional] 
**MediaTypes** | Pointer to [**[]ConnectionServicesDetailsInnerMediaTypesInner**](ConnectionServicesDetailsInnerMediaTypesInner.md) |  | [optional] 

## Methods

### NewConnectionServicesDetailsInner

`func NewConnectionServicesDetailsInner() *ConnectionServicesDetailsInner`

NewConnectionServicesDetailsInner instantiates a new ConnectionServicesDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServicesDetailsInnerWithDefaults

`func NewConnectionServicesDetailsInnerWithDefaults() *ConnectionServicesDetailsInner`

NewConnectionServicesDetailsInnerWithDefaults instantiates a new ConnectionServicesDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectionServicesDetailsInner) GetName() ConnectionServices`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionServicesDetailsInner) GetNameOk() (*ConnectionServices, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionServicesDetailsInner) SetName(v ConnectionServices)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionServicesDetailsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMediaTypes

`func (o *ConnectionServicesDetailsInner) GetMediaTypes() []ConnectionServicesDetailsInnerMediaTypesInner`

GetMediaTypes returns the MediaTypes field if non-nil, zero value otherwise.

### GetMediaTypesOk

`func (o *ConnectionServicesDetailsInner) GetMediaTypesOk() (*[]ConnectionServicesDetailsInnerMediaTypesInner, bool)`

GetMediaTypesOk returns a tuple with the MediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypes

`func (o *ConnectionServicesDetailsInner) SetMediaTypes(v []ConnectionServicesDetailsInnerMediaTypesInner)`

SetMediaTypes sets MediaTypes field to given value.

### HasMediaTypes

`func (o *ConnectionServicesDetailsInner) HasMediaTypes() bool`

HasMediaTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


