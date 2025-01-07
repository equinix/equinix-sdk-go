# SubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**Channel**](Channel.md) |  | [optional] 
**MessageType** | Pointer to [**MessageType**](MessageType.md) |  | [optional] 

## Methods

### NewSubscriptionRequest

`func NewSubscriptionRequest() *SubscriptionRequest`

NewSubscriptionRequest instantiates a new SubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionRequestWithDefaults

`func NewSubscriptionRequestWithDefaults() *SubscriptionRequest`

NewSubscriptionRequestWithDefaults instantiates a new SubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SubscriptionRequest) GetChannel() Channel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SubscriptionRequest) GetChannelOk() (*Channel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SubscriptionRequest) SetChannel(v Channel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SubscriptionRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMessageType

`func (o *SubscriptionRequest) GetMessageType() MessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SubscriptionRequest) GetMessageTypeOk() (*MessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SubscriptionRequest) SetMessageType(v MessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *SubscriptionRequest) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


