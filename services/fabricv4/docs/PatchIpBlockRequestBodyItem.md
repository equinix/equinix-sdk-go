# PatchIpBlockRequestBodyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**PatchIpBlockRequestBodyItemOp**](PatchIpBlockRequestBodyItemOp.md) |  | 
**Path** | **string** | path | 
**Value** | Pointer to [**PatchIpBlockRequestBodyItemValue**](PatchIpBlockRequestBodyItemValue.md) |  | [optional] 

## Methods

### NewPatchIpBlockRequestBodyItem

`func NewPatchIpBlockRequestBodyItem(op PatchIpBlockRequestBodyItemOp, path string, ) *PatchIpBlockRequestBodyItem`

NewPatchIpBlockRequestBodyItem instantiates a new PatchIpBlockRequestBodyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchIpBlockRequestBodyItemWithDefaults

`func NewPatchIpBlockRequestBodyItemWithDefaults() *PatchIpBlockRequestBodyItem`

NewPatchIpBlockRequestBodyItemWithDefaults instantiates a new PatchIpBlockRequestBodyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchIpBlockRequestBodyItem) GetOp() PatchIpBlockRequestBodyItemOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchIpBlockRequestBodyItem) GetOpOk() (*PatchIpBlockRequestBodyItemOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchIpBlockRequestBodyItem) SetOp(v PatchIpBlockRequestBodyItemOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *PatchIpBlockRequestBodyItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchIpBlockRequestBodyItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchIpBlockRequestBodyItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *PatchIpBlockRequestBodyItem) GetValue() PatchIpBlockRequestBodyItemValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchIpBlockRequestBodyItem) GetValueOk() (*PatchIpBlockRequestBodyItemValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchIpBlockRequestBodyItem) SetValue(v PatchIpBlockRequestBodyItemValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchIpBlockRequestBodyItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


