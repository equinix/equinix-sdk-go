# AlertMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**Alert**](Alert.md) |  | 
**Type** | **string** | message type | [default to "custom-alert"]

## Methods

### NewAlertMessageData

`func NewAlertMessageData(data Alert, type_ string, ) *AlertMessageData`

NewAlertMessageData instantiates a new AlertMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertMessageDataWithDefaults

`func NewAlertMessageDataWithDefaults() *AlertMessageData`

NewAlertMessageDataWithDefaults instantiates a new AlertMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AlertMessageData) GetData() Alert`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AlertMessageData) GetDataOk() (*Alert, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AlertMessageData) SetData(v Alert)`

SetData sets Data field to given value.


### GetType

`func (o *AlertMessageData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertMessageData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertMessageData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


