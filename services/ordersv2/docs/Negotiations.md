# Negotiations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | Pointer to **string** | Reference id / internal id for customer to identify the actionable order line | [optional] 
**OrderRequestedDateTime** | Pointer to **time.Time** | Customer requested order date and time at point of order. | [optional] 
**ProposedDateTime** | Pointer to **time.Time** | Alternative / revised date and time proposed by Equinix for fulfilling the order request. | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Validity date and time of the negotiation. Negotiations request after this date will no longer be valid. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Negotiations message created date and time by operations team | [optional] 
**Expedited** | Pointer to **bool** | When &#x60;true,&#x60; requested order is an expedite order | [optional] 
**Message** | Pointer to **string** | Negotiations message | [optional] 

## Methods

### NewNegotiations

`func NewNegotiations() *Negotiations`

NewNegotiations instantiates a new Negotiations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsWithDefaults

`func NewNegotiationsWithDefaults() *Negotiations`

NewNegotiationsWithDefaults instantiates a new Negotiations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *Negotiations) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Negotiations) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Negotiations) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Negotiations) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetOrderRequestedDateTime

`func (o *Negotiations) GetOrderRequestedDateTime() time.Time`

GetOrderRequestedDateTime returns the OrderRequestedDateTime field if non-nil, zero value otherwise.

### GetOrderRequestedDateTimeOk

`func (o *Negotiations) GetOrderRequestedDateTimeOk() (*time.Time, bool)`

GetOrderRequestedDateTimeOk returns a tuple with the OrderRequestedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRequestedDateTime

`func (o *Negotiations) SetOrderRequestedDateTime(v time.Time)`

SetOrderRequestedDateTime sets OrderRequestedDateTime field to given value.

### HasOrderRequestedDateTime

`func (o *Negotiations) HasOrderRequestedDateTime() bool`

HasOrderRequestedDateTime returns a boolean if a field has been set.

### GetProposedDateTime

`func (o *Negotiations) GetProposedDateTime() time.Time`

GetProposedDateTime returns the ProposedDateTime field if non-nil, zero value otherwise.

### GetProposedDateTimeOk

`func (o *Negotiations) GetProposedDateTimeOk() (*time.Time, bool)`

GetProposedDateTimeOk returns a tuple with the ProposedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedDateTime

`func (o *Negotiations) SetProposedDateTime(v time.Time)`

SetProposedDateTime sets ProposedDateTime field to given value.

### HasProposedDateTime

`func (o *Negotiations) HasProposedDateTime() bool`

HasProposedDateTime returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *Negotiations) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *Negotiations) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *Negotiations) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *Negotiations) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Negotiations) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Negotiations) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Negotiations) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Negotiations) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetExpedited

`func (o *Negotiations) GetExpedited() bool`

GetExpedited returns the Expedited field if non-nil, zero value otherwise.

### GetExpeditedOk

`func (o *Negotiations) GetExpeditedOk() (*bool, bool)`

GetExpeditedOk returns a tuple with the Expedited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpedited

`func (o *Negotiations) SetExpedited(v bool)`

SetExpedited sets Expedited field to given value.

### HasExpedited

`func (o *Negotiations) HasExpedited() bool`

HasExpedited returns a boolean if a field has been set.

### GetMessage

`func (o *Negotiations) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Negotiations) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Negotiations) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Negotiations) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


