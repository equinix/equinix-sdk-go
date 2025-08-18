# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ibx** | **string** | Ibx Name | 
**Cages** | [**[]Cage**](Cage.md) | List of cage | 

## Methods

### NewLocation

`func NewLocation(ibx string, cages []Cage, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbx

`func (o *Location) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *Location) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *Location) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetCages

`func (o *Location) GetCages() []Cage`

GetCages returns the Cages field if non-nil, zero value otherwise.

### GetCagesOk

`func (o *Location) GetCagesOk() (*[]Cage, bool)`

GetCagesOk returns a tuple with the Cages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCages

`func (o *Location) SetCages(v []Cage)`

SetCages sets Cages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


