# VirtualPortLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ibx** | Pointer to **string** | Code assigned to the Equinix International Business Exchange (IBX) data center from which the port is ordered. &lt;br&gt; The port might be in a different location. | [optional] 

## Methods

### NewVirtualPortLocation

`func NewVirtualPortLocation() *VirtualPortLocation`

NewVirtualPortLocation instantiates a new VirtualPortLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualPortLocationWithDefaults

`func NewVirtualPortLocationWithDefaults() *VirtualPortLocation`

NewVirtualPortLocationWithDefaults instantiates a new VirtualPortLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbx

`func (o *VirtualPortLocation) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *VirtualPortLocation) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *VirtualPortLocation) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *VirtualPortLocation) HasIbx() bool`

HasIbx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


