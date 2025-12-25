# Metro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetroCode** | Pointer to **string** | Metro code | [optional] 
**MetroDescription** | Pointer to **string** | Metro description | [optional] 
**Region** | Pointer to **string** | A region have several metros. | [optional] 
**ClusterSupported** | Pointer to **bool** | Whether this metro supports cluster devices | [optional] 

## Methods

### NewMetro

`func NewMetro() *Metro`

NewMetro instantiates a new Metro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroWithDefaults

`func NewMetroWithDefaults() *Metro`

NewMetroWithDefaults instantiates a new Metro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetroCode

`func (o *Metro) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *Metro) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *Metro) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *Metro) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetMetroDescription

`func (o *Metro) GetMetroDescription() string`

GetMetroDescription returns the MetroDescription field if non-nil, zero value otherwise.

### GetMetroDescriptionOk

`func (o *Metro) GetMetroDescriptionOk() (*string, bool)`

GetMetroDescriptionOk returns a tuple with the MetroDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroDescription

`func (o *Metro) SetMetroDescription(v string)`

SetMetroDescription sets MetroDescription field to given value.

### HasMetroDescription

`func (o *Metro) HasMetroDescription() bool`

HasMetroDescription returns a boolean if a field has been set.

### GetRegion

`func (o *Metro) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Metro) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Metro) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Metro) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetClusterSupported

`func (o *Metro) GetClusterSupported() bool`

GetClusterSupported returns the ClusterSupported field if non-nil, zero value otherwise.

### GetClusterSupportedOk

`func (o *Metro) GetClusterSupportedOk() (*bool, bool)`

GetClusterSupportedOk returns a tuple with the ClusterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSupported

`func (o *Metro) SetClusterSupported(v bool)`

SetClusterSupported sets ClusterSupported field to given value.

### HasClusterSupported

`func (o *Metro) HasClusterSupported() bool`

HasClusterSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


