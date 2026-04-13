# ServiceProfileActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Action type. Example values: PROFILE_UPDATE_ACCEPTANCE, PROFILE_UPDATE_REJECTION | 
**Description** | Pointer to **string** | Action description | [optional] 

## Methods

### NewServiceProfileActionRequest

`func NewServiceProfileActionRequest(type_ string, ) *ServiceProfileActionRequest`

NewServiceProfileActionRequest instantiates a new ServiceProfileActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileActionRequestWithDefaults

`func NewServiceProfileActionRequestWithDefaults() *ServiceProfileActionRequest`

NewServiceProfileActionRequestWithDefaults instantiates a new ServiceProfileActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceProfileActionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileActionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileActionRequest) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *ServiceProfileActionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceProfileActionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceProfileActionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceProfileActionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


