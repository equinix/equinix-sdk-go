# PowerMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Power**](Power.md) |  | 
**Type** | **string** | message type | [default to "power"]

## Methods

### NewPowerMessageData

`func NewPowerMessageData(data Power, type_ string, ) *PowerMessageData`

NewPowerMessageData instantiates a new PowerMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerMessageDataWithDefaults

`func NewPowerMessageDataWithDefaults() *PowerMessageData`

NewPowerMessageDataWithDefaults instantiates a new PowerMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PowerMessageData) GetData() Power`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PowerMessageData) GetDataOk() (*Power, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PowerMessageData) SetData(v Power)`

SetData sets Data field to given value.


### GetType

`func (o *PowerMessageData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerMessageData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerMessageData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


