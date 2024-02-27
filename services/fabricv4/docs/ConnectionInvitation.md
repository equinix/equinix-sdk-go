# ConnectionInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | invitee email | [optional] 
**Message** | Pointer to **string** | invitation message | [optional] 
**CtrDraftOrderId** | Pointer to **string** | draft order id for invitation | [optional] 

## Methods

### NewConnectionInvitation

`func NewConnectionInvitation() *ConnectionInvitation`

NewConnectionInvitation instantiates a new ConnectionInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionInvitationWithDefaults

`func NewConnectionInvitationWithDefaults() *ConnectionInvitation`

NewConnectionInvitationWithDefaults instantiates a new ConnectionInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ConnectionInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConnectionInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConnectionInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ConnectionInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMessage

`func (o *ConnectionInvitation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectionInvitation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectionInvitation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConnectionInvitation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCtrDraftOrderId

`func (o *ConnectionInvitation) GetCtrDraftOrderId() string`

GetCtrDraftOrderId returns the CtrDraftOrderId field if non-nil, zero value otherwise.

### GetCtrDraftOrderIdOk

`func (o *ConnectionInvitation) GetCtrDraftOrderIdOk() (*string, bool)`

GetCtrDraftOrderIdOk returns a tuple with the CtrDraftOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtrDraftOrderId

`func (o *ConnectionInvitation) SetCtrDraftOrderId(v string)`

SetCtrDraftOrderId sets CtrDraftOrderId field to given value.

### HasCtrDraftOrderId

`func (o *ConnectionInvitation) HasCtrDraftOrderId() bool`

HasCtrDraftOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


