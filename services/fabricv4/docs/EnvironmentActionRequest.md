# EnvironmentActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnvironmentActionTypeEnum**](EnvironmentActionTypeEnum.md) |  | 
**KeyDetails** | [**ActivationKeyDetails**](ActivationKeyDetails.md) |  | 

## Methods

### NewEnvironmentActionRequest

`func NewEnvironmentActionRequest(type_ EnvironmentActionTypeEnum, keyDetails ActivationKeyDetails, ) *EnvironmentActionRequest`

NewEnvironmentActionRequest instantiates a new EnvironmentActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentActionRequestWithDefaults

`func NewEnvironmentActionRequestWithDefaults() *EnvironmentActionRequest`

NewEnvironmentActionRequestWithDefaults instantiates a new EnvironmentActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnvironmentActionRequest) GetType() EnvironmentActionTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentActionRequest) GetTypeOk() (*EnvironmentActionTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentActionRequest) SetType(v EnvironmentActionTypeEnum)`

SetType sets Type field to given value.


### GetKeyDetails

`func (o *EnvironmentActionRequest) GetKeyDetails() ActivationKeyDetails`

GetKeyDetails returns the KeyDetails field if non-nil, zero value otherwise.

### GetKeyDetailsOk

`func (o *EnvironmentActionRequest) GetKeyDetailsOk() (*ActivationKeyDetails, bool)`

GetKeyDetailsOk returns a tuple with the KeyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDetails

`func (o *EnvironmentActionRequest) SetKeyDetails(v ActivationKeyDetails)`

SetKeyDetails sets KeyDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


