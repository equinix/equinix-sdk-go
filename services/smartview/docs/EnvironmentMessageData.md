# EnvironmentMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Environment**](Environment.md) |  | 
**Type** | **string** | message type | [default to "environmental"]

## Methods

### NewEnvironmentMessageData

`func NewEnvironmentMessageData(data Environment, type_ string, ) *EnvironmentMessageData`

NewEnvironmentMessageData instantiates a new EnvironmentMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentMessageDataWithDefaults

`func NewEnvironmentMessageDataWithDefaults() *EnvironmentMessageData`

NewEnvironmentMessageDataWithDefaults instantiates a new EnvironmentMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnvironmentMessageData) GetData() Environment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnvironmentMessageData) GetDataOk() (*Environment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnvironmentMessageData) SetData(v Environment)`

SetData sets Data field to given value.


### GetType

`func (o *EnvironmentMessageData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentMessageData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentMessageData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


