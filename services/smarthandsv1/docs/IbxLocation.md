# IbxLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ibx** | **string** |  | 
**Cages** | [**[]IbxLocationCagesInner**](IbxLocationCagesInner.md) |  | 

## Methods

### NewIbxLocation

`func NewIbxLocation(ibx string, cages []IbxLocationCagesInner, ) *IbxLocation`

NewIbxLocation instantiates a new IbxLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbxLocationWithDefaults

`func NewIbxLocationWithDefaults() *IbxLocation`

NewIbxLocationWithDefaults instantiates a new IbxLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbx

`func (o *IbxLocation) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *IbxLocation) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *IbxLocation) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetCages

`func (o *IbxLocation) GetCages() []IbxLocationCagesInner`

GetCages returns the Cages field if non-nil, zero value otherwise.

### GetCagesOk

`func (o *IbxLocation) GetCagesOk() (*[]IbxLocationCagesInner, bool)`

GetCagesOk returns a tuple with the Cages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCages

`func (o *IbxLocation) SetCages(v []IbxLocationCagesInner)`

SetCages sets Cages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


