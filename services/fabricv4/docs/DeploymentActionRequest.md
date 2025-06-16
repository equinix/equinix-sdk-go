# DeploymentActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**DeploymentActionType**](DeploymentActionType.md) |  | 
**Data** | [**[]ActionRequest**](ActionRequest.md) |  | 

## Methods

### NewDeploymentActionRequest

`func NewDeploymentActionRequest(type_ DeploymentActionType, data []ActionRequest, ) *DeploymentActionRequest`

NewDeploymentActionRequest instantiates a new DeploymentActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentActionRequestWithDefaults

`func NewDeploymentActionRequestWithDefaults() *DeploymentActionRequest`

NewDeploymentActionRequestWithDefaults instantiates a new DeploymentActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentActionRequest) GetType() DeploymentActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentActionRequest) GetTypeOk() (*DeploymentActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentActionRequest) SetType(v DeploymentActionType)`

SetType sets Type field to given value.


### GetData

`func (o *DeploymentActionRequest) GetData() []ActionRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeploymentActionRequest) GetDataOk() (*[]ActionRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeploymentActionRequest) SetData(v []ActionRequest)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


