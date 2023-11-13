# PortAssignInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vnid** | Pointer to **string** | Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself.  | [optional] 

## Methods

### NewPortAssignInput

`func NewPortAssignInput() *PortAssignInput`

NewPortAssignInput instantiates a new PortAssignInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortAssignInputWithDefaults

`func NewPortAssignInputWithDefaults() *PortAssignInput`

NewPortAssignInputWithDefaults instantiates a new PortAssignInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVnid

`func (o *PortAssignInput) GetVnid() string`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *PortAssignInput) GetVnidOk() (*string, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *PortAssignInput) SetVnid(v string)`

SetVnid sets Vnid field to given value.

### HasVnid

`func (o *PortAssignInput) HasVnid() bool`

HasVnid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


