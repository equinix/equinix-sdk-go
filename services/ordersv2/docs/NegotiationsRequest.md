# NegotiationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **string** | Unique identifier to reference specific activity or order line id.  | 
**Action** | [**NegotiationsRequestAction**](NegotiationsRequestAction.md) |  | 
**Reason** | Pointer to **string** | Reason for cancelling the negotiation | [optional] 

## Methods

### NewNegotiationsRequest

`func NewNegotiationsRequest(referenceId string, action NegotiationsRequestAction, ) *NegotiationsRequest`

NewNegotiationsRequest instantiates a new NegotiationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsRequestWithDefaults

`func NewNegotiationsRequestWithDefaults() *NegotiationsRequest`

NewNegotiationsRequestWithDefaults instantiates a new NegotiationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *NegotiationsRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *NegotiationsRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *NegotiationsRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetAction

`func (o *NegotiationsRequest) GetAction() NegotiationsRequestAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NegotiationsRequest) GetActionOk() (*NegotiationsRequestAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NegotiationsRequest) SetAction(v NegotiationsRequestAction)`

SetAction sets Action field to given value.


### GetReason

`func (o *NegotiationsRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NegotiationsRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NegotiationsRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NegotiationsRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


