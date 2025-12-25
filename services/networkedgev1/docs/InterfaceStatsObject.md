# InterfaceStatsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stats** | Pointer to [**InterfaceStatsDetailObject**](InterfaceStatsDetailObject.md) |  | [optional] 

## Methods

### NewInterfaceStatsObject

`func NewInterfaceStatsObject() *InterfaceStatsObject`

NewInterfaceStatsObject instantiates a new InterfaceStatsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceStatsObjectWithDefaults

`func NewInterfaceStatsObjectWithDefaults() *InterfaceStatsObject`

NewInterfaceStatsObjectWithDefaults instantiates a new InterfaceStatsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStats

`func (o *InterfaceStatsObject) GetStats() InterfaceStatsDetailObject`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *InterfaceStatsObject) GetStatsOk() (*InterfaceStatsDetailObject, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *InterfaceStatsObject) SetStats(v InterfaceStatsDetailObject)`

SetStats sets Stats field to given value.

### HasStats

`func (o *InterfaceStatsObject) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


