# DeleteConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**PrimaryConnectionId** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteConnectionResponse

`func NewDeleteConnectionResponse() *DeleteConnectionResponse`

NewDeleteConnectionResponse instantiates a new DeleteConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteConnectionResponseWithDefaults

`func NewDeleteConnectionResponseWithDefaults() *DeleteConnectionResponse`

NewDeleteConnectionResponseWithDefaults instantiates a new DeleteConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DeleteConnectionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeleteConnectionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeleteConnectionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DeleteConnectionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPrimaryConnectionId

`func (o *DeleteConnectionResponse) GetPrimaryConnectionId() string`

GetPrimaryConnectionId returns the PrimaryConnectionId field if non-nil, zero value otherwise.

### GetPrimaryConnectionIdOk

`func (o *DeleteConnectionResponse) GetPrimaryConnectionIdOk() (*string, bool)`

GetPrimaryConnectionIdOk returns a tuple with the PrimaryConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryConnectionId

`func (o *DeleteConnectionResponse) SetPrimaryConnectionId(v string)`

SetPrimaryConnectionId sets PrimaryConnectionId field to given value.

### HasPrimaryConnectionId

`func (o *DeleteConnectionResponse) HasPrimaryConnectionId() bool`

HasPrimaryConnectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


