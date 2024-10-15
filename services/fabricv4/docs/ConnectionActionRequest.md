# ConnectionActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**Actions**](Actions.md) |  | 
**Description** | Pointer to **string** | Connection rejection reason detail | [optional] 
**Data** | Pointer to [**ConnectionAcceptanceData**](ConnectionAcceptanceData.md) |  | [optional] 

## Methods

### NewConnectionActionRequest

`func NewConnectionActionRequest(type_ Actions, ) *ConnectionActionRequest`

NewConnectionActionRequest instantiates a new ConnectionActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionActionRequestWithDefaults

`func NewConnectionActionRequestWithDefaults() *ConnectionActionRequest`

NewConnectionActionRequestWithDefaults instantiates a new ConnectionActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionActionRequest) GetType() Actions`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionActionRequest) GetTypeOk() (*Actions, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionActionRequest) SetType(v Actions)`

SetType sets Type field to given value.


### GetDescription

`func (o *ConnectionActionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionActionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionActionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionActionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetData

`func (o *ConnectionActionRequest) GetData() ConnectionAcceptanceData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectionActionRequest) GetDataOk() (*ConnectionAcceptanceData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectionActionRequest) SetData(v ConnectionAcceptanceData)`

SetData sets Data field to given value.

### HasData

`func (o *ConnectionActionRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


