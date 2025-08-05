# CancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | Reason for cancellation | 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | File(s) attached | [optional] 
**LineIds** | Pointer to **[]string** | Refers to the &#x60;lineId&#x60; of product/service. | [optional] 

## Methods

### NewCancelRequest

`func NewCancelRequest(reason string, ) *CancelRequest`

NewCancelRequest instantiates a new CancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelRequestWithDefaults

`func NewCancelRequestWithDefaults() *CancelRequest`

NewCancelRequestWithDefaults instantiates a new CancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CancelRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAttachments

`func (o *CancelRequest) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CancelRequest) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CancelRequest) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CancelRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetLineIds

`func (o *CancelRequest) GetLineIds() []string`

GetLineIds returns the LineIds field if non-nil, zero value otherwise.

### GetLineIdsOk

`func (o *CancelRequest) GetLineIdsOk() (*[]string, bool)`

GetLineIdsOk returns a tuple with the LineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIds

`func (o *CancelRequest) SetLineIds(v []string)`

SetLineIds sets LineIds field to given value.

### HasLineIds

`func (o *CancelRequest) HasLineIds() bool`

HasLineIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


