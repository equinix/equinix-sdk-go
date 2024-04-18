# MetroResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetroCode** | Pointer to **string** | Metro code | [optional] 
**MetroDescription** | Pointer to **string** | Metro description | [optional] 
**Region** | Pointer to **string** | Region within which the metro is located | [optional] 
**ClusterSupported** | Pointer to **bool** | Whether this metro supports cluster devices | [optional] 

## Methods

### NewMetroResponse

`func NewMetroResponse() *MetroResponse`

NewMetroResponse instantiates a new MetroResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroResponseWithDefaults

`func NewMetroResponseWithDefaults() *MetroResponse`

NewMetroResponseWithDefaults instantiates a new MetroResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetroCode

`func (o *MetroResponse) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *MetroResponse) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *MetroResponse) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *MetroResponse) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetMetroDescription

`func (o *MetroResponse) GetMetroDescription() string`

GetMetroDescription returns the MetroDescription field if non-nil, zero value otherwise.

### GetMetroDescriptionOk

`func (o *MetroResponse) GetMetroDescriptionOk() (*string, bool)`

GetMetroDescriptionOk returns a tuple with the MetroDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroDescription

`func (o *MetroResponse) SetMetroDescription(v string)`

SetMetroDescription sets MetroDescription field to given value.

### HasMetroDescription

`func (o *MetroResponse) HasMetroDescription() bool`

HasMetroDescription returns a boolean if a field has been set.

### GetRegion

`func (o *MetroResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MetroResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MetroResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MetroResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetClusterSupported

`func (o *MetroResponse) GetClusterSupported() bool`

GetClusterSupported returns the ClusterSupported field if non-nil, zero value otherwise.

### GetClusterSupportedOk

`func (o *MetroResponse) GetClusterSupportedOk() (*bool, bool)`

GetClusterSupportedOk returns a tuple with the ClusterSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSupported

`func (o *MetroResponse) SetClusterSupported(v bool)`

SetClusterSupported sets ClusterSupported field to given value.

### HasClusterSupported

`func (o *MetroResponse) HasClusterSupported() bool`

HasClusterSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


