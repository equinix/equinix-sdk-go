# AdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Reason of the error | [optional] 
**Property** | Pointer to **string** | Request property that caused the error | [optional] 

## Methods

### NewAdditionalInfo

`func NewAdditionalInfo() *AdditionalInfo`

NewAdditionalInfo instantiates a new AdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalInfoWithDefaults

`func NewAdditionalInfoWithDefaults() *AdditionalInfo`

NewAdditionalInfoWithDefaults instantiates a new AdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *AdditionalInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AdditionalInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AdditionalInfo) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *AdditionalInfo) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetProperty

`func (o *AdditionalInfo) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AdditionalInfo) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AdditionalInfo) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *AdditionalInfo) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


