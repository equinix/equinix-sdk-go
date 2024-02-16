# AWSFabricProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSFabricProviderType**](AWSFabricProviderType.md) |  | 
**AccountId** | **string** | AWS Account ID | 

## Methods

### NewAWSFabricProvider

`func NewAWSFabricProvider(type_ AWSFabricProviderType, accountId string, ) *AWSFabricProvider`

NewAWSFabricProvider instantiates a new AWSFabricProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSFabricProviderWithDefaults

`func NewAWSFabricProviderWithDefaults() *AWSFabricProvider`

NewAWSFabricProviderWithDefaults instantiates a new AWSFabricProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSFabricProvider) GetType() AWSFabricProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSFabricProvider) GetTypeOk() (*AWSFabricProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSFabricProvider) SetType(v AWSFabricProviderType)`

SetType sets Type field to given value.


### GetAccountId

`func (o *AWSFabricProvider) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSFabricProvider) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSFabricProvider) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


