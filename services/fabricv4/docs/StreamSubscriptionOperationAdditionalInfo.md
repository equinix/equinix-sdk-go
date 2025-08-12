# StreamSubscriptionOperationAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Additional attribute for error information | [optional] 
**StatusCode** | Pointer to **string** | HTTP error status code response from sink type if error occurred | [optional] 
**Reason** | Pointer to **string** | HTTP error message response from sink type if error occurred | [optional] 

## Methods

### NewStreamSubscriptionOperationAdditionalInfo

`func NewStreamSubscriptionOperationAdditionalInfo() *StreamSubscriptionOperationAdditionalInfo`

NewStreamSubscriptionOperationAdditionalInfo instantiates a new StreamSubscriptionOperationAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionOperationAdditionalInfoWithDefaults

`func NewStreamSubscriptionOperationAdditionalInfoWithDefaults() *StreamSubscriptionOperationAdditionalInfo`

NewStreamSubscriptionOperationAdditionalInfoWithDefaults instantiates a new StreamSubscriptionOperationAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *StreamSubscriptionOperationAdditionalInfo) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamSubscriptionOperationAdditionalInfo) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamSubscriptionOperationAdditionalInfo) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamSubscriptionOperationAdditionalInfo) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetStatusCode

`func (o *StreamSubscriptionOperationAdditionalInfo) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *StreamSubscriptionOperationAdditionalInfo) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *StreamSubscriptionOperationAdditionalInfo) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *StreamSubscriptionOperationAdditionalInfo) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetReason

`func (o *StreamSubscriptionOperationAdditionalInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StreamSubscriptionOperationAdditionalInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StreamSubscriptionOperationAdditionalInfo) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *StreamSubscriptionOperationAdditionalInfo) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


