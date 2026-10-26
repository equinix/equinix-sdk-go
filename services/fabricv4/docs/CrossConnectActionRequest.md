# CrossConnectActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CrossConnectActionType**](CrossConnectActionType.md) |  | 
**Data** | Pointer to [**CrossConnectActionData**](CrossConnectActionData.md) |  | [optional] 

## Methods

### NewCrossConnectActionRequest

`func NewCrossConnectActionRequest(type_ CrossConnectActionType, ) *CrossConnectActionRequest`

NewCrossConnectActionRequest instantiates a new CrossConnectActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionRequestWithDefaults

`func NewCrossConnectActionRequestWithDefaults() *CrossConnectActionRequest`

NewCrossConnectActionRequestWithDefaults instantiates a new CrossConnectActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CrossConnectActionRequest) GetType() CrossConnectActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectActionRequest) GetTypeOk() (*CrossConnectActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectActionRequest) SetType(v CrossConnectActionType)`

SetType sets Type field to given value.


### GetData

`func (o *CrossConnectActionRequest) GetData() CrossConnectActionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CrossConnectActionRequest) GetDataOk() (*CrossConnectActionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CrossConnectActionRequest) SetData(v CrossConnectActionData)`

SetData sets Data field to given value.

### HasData

`func (o *CrossConnectActionRequest) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


