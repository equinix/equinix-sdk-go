# CloudEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Cloud Event message | [optional] 
**Resource** | Pointer to [**ResourceData**](ResourceData.md) |  | [optional] 
**Auth** | Pointer to [**AuthContext**](AuthContext.md) |  | [optional] 

## Methods

### NewCloudEventData

`func NewCloudEventData() *CloudEventData`

NewCloudEventData instantiates a new CloudEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEventDataWithDefaults

`func NewCloudEventDataWithDefaults() *CloudEventData`

NewCloudEventDataWithDefaults instantiates a new CloudEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CloudEventData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudEventData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudEventData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudEventData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResource

`func (o *CloudEventData) GetResource() ResourceData`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CloudEventData) GetResourceOk() (*ResourceData, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CloudEventData) SetResource(v ResourceData)`

SetResource sets Resource field to given value.

### HasResource

`func (o *CloudEventData) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetAuth

`func (o *CloudEventData) GetAuth() AuthContext`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CloudEventData) GetAuthOk() (*AuthContext, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CloudEventData) SetAuth(v AuthContext)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *CloudEventData) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


