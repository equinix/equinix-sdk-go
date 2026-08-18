# LoaActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LoaActionType**](LoaActionType.md) |  | 
**Data** | Pointer to [**LoaActionData**](LoaActionData.md) |  | [optional] 

## Methods

### NewLoaActionRequest

`func NewLoaActionRequest(type_ LoaActionType, ) *LoaActionRequest`

NewLoaActionRequest instantiates a new LoaActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaActionRequestWithDefaults

`func NewLoaActionRequestWithDefaults() *LoaActionRequest`

NewLoaActionRequestWithDefaults instantiates a new LoaActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoaActionRequest) GetType() LoaActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoaActionRequest) GetTypeOk() (*LoaActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoaActionRequest) SetType(v LoaActionType)`

SetType sets Type field to given value.


### GetData

`func (o *LoaActionRequest) GetData() LoaActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoaActionRequest) GetDataOk() (*LoaActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoaActionRequest) SetData(v LoaActionData)`

SetData sets Data field to given value.

### HasData

`func (o *LoaActionRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


