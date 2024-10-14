# CloudRouterActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CloudRouterActionType**](CloudRouterActionType.md) |  | 
**Connection** | Pointer to [**RouterActionsConnection**](RouterActionsConnection.md) |  | [optional] 

## Methods

### NewCloudRouterActionRequest

`func NewCloudRouterActionRequest(type_ CloudRouterActionType, ) *CloudRouterActionRequest`

NewCloudRouterActionRequest instantiates a new CloudRouterActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterActionRequestWithDefaults

`func NewCloudRouterActionRequestWithDefaults() *CloudRouterActionRequest`

NewCloudRouterActionRequestWithDefaults instantiates a new CloudRouterActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudRouterActionRequest) GetType() CloudRouterActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterActionRequest) GetTypeOk() (*CloudRouterActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterActionRequest) SetType(v CloudRouterActionType)`

SetType sets Type field to given value.


### GetConnection

`func (o *CloudRouterActionRequest) GetConnection() RouterActionsConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *CloudRouterActionRequest) GetConnectionOk() (*RouterActionsConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *CloudRouterActionRequest) SetConnection(v RouterActionsConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *CloudRouterActionRequest) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


