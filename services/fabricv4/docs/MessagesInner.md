# MessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Role of the message sender user or assistant | [optional] 
**Content** | Pointer to **string** | Content of the chat message | [optional] 

## Methods

### NewMessagesInner

`func NewMessagesInner() *MessagesInner`

NewMessagesInner instantiates a new MessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesInnerWithDefaults

`func NewMessagesInnerWithDefaults() *MessagesInner`

NewMessagesInnerWithDefaults instantiates a new MessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessagesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessagesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessagesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MessagesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *MessagesInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessagesInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessagesInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MessagesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


