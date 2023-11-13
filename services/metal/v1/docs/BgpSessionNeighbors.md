# BgpSessionNeighbors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BgpNeighbors** | Pointer to [**[]BgpNeighborData**](BgpNeighborData.md) | A list of BGP session neighbor data | [optional] 

## Methods

### NewBgpSessionNeighbors

`func NewBgpSessionNeighbors() *BgpSessionNeighbors`

NewBgpSessionNeighbors instantiates a new BgpSessionNeighbors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpSessionNeighborsWithDefaults

`func NewBgpSessionNeighborsWithDefaults() *BgpSessionNeighbors`

NewBgpSessionNeighborsWithDefaults instantiates a new BgpSessionNeighbors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBgpNeighbors

`func (o *BgpSessionNeighbors) GetBgpNeighbors() []BgpNeighborData`

GetBgpNeighbors returns the BgpNeighbors field if non-nil, zero value otherwise.

### GetBgpNeighborsOk

`func (o *BgpSessionNeighbors) GetBgpNeighborsOk() (*[]BgpNeighborData, bool)`

GetBgpNeighborsOk returns a tuple with the BgpNeighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpNeighbors

`func (o *BgpSessionNeighbors) SetBgpNeighbors(v []BgpNeighborData)`

SetBgpNeighbors sets BgpNeighbors field to given value.

### HasBgpNeighbors

`func (o *BgpSessionNeighbors) HasBgpNeighbors() bool`

HasBgpNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


