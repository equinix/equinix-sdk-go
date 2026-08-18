# AppSubscriptionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AppSubscriptionType**](AppSubscriptionType.md) |  | [default to APPSUBSCRIPTIONTYPE_APP_SUBSCRIPTION]
**Source** | Pointer to [**AppSubscriptionSourceRequest**](AppSubscriptionSourceRequest.md) |  | [optional] 
**Target** | Pointer to [**AppSubscriptionTargetRequest**](AppSubscriptionTargetRequest.md) |  | [optional] 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewAppSubscriptionPostRequest

`func NewAppSubscriptionPostRequest(type_ AppSubscriptionType, project Project, ) *AppSubscriptionPostRequest`

NewAppSubscriptionPostRequest instantiates a new AppSubscriptionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionPostRequestWithDefaults

`func NewAppSubscriptionPostRequestWithDefaults() *AppSubscriptionPostRequest`

NewAppSubscriptionPostRequestWithDefaults instantiates a new AppSubscriptionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppSubscriptionPostRequest) GetType() AppSubscriptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubscriptionPostRequest) GetTypeOk() (*AppSubscriptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubscriptionPostRequest) SetType(v AppSubscriptionType)`

SetType sets Type field to given value.


### GetSource

`func (o *AppSubscriptionPostRequest) GetSource() AppSubscriptionSourceRequest`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppSubscriptionPostRequest) GetSourceOk() (*AppSubscriptionSourceRequest, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppSubscriptionPostRequest) SetSource(v AppSubscriptionSourceRequest)`

SetSource sets Source field to given value.

### HasSource

`func (o *AppSubscriptionPostRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *AppSubscriptionPostRequest) GetTarget() AppSubscriptionTargetRequest`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AppSubscriptionPostRequest) GetTargetOk() (*AppSubscriptionTargetRequest, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AppSubscriptionPostRequest) SetTarget(v AppSubscriptionTargetRequest)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AppSubscriptionPostRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetProject

`func (o *AppSubscriptionPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppSubscriptionPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppSubscriptionPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


