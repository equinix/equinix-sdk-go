# InternetAccessPatchOperationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**InternetAccessPatchOperationUpdateAllowedOp**](InternetAccessPatchOperationUpdateAllowedOp.md) |  | 
**Path** | **string** | Allowed patch paths for Internet Access update. | 
**Value** | Pointer to **interface{}** | New value for updated parameter. Required for add and replace. | [optional] 

## Methods

### NewInternetAccessPatchOperationUpdate

`func NewInternetAccessPatchOperationUpdate(op InternetAccessPatchOperationUpdateAllowedOp, path string, ) *InternetAccessPatchOperationUpdate`

NewInternetAccessPatchOperationUpdate instantiates a new InternetAccessPatchOperationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessPatchOperationUpdateWithDefaults

`func NewInternetAccessPatchOperationUpdateWithDefaults() *InternetAccessPatchOperationUpdate`

NewInternetAccessPatchOperationUpdateWithDefaults instantiates a new InternetAccessPatchOperationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *InternetAccessPatchOperationUpdate) GetOp() InternetAccessPatchOperationUpdateAllowedOp`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *InternetAccessPatchOperationUpdate) GetOpOk() (*InternetAccessPatchOperationUpdateAllowedOp, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *InternetAccessPatchOperationUpdate) SetOp(v InternetAccessPatchOperationUpdateAllowedOp)`

SetOp sets Op field to given value.


### GetPath

`func (o *InternetAccessPatchOperationUpdate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *InternetAccessPatchOperationUpdate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *InternetAccessPatchOperationUpdate) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *InternetAccessPatchOperationUpdate) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InternetAccessPatchOperationUpdate) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InternetAccessPatchOperationUpdate) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *InternetAccessPatchOperationUpdate) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *InternetAccessPatchOperationUpdate) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *InternetAccessPatchOperationUpdate) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


