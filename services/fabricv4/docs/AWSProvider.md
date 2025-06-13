# AWSProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSProviderType**](AWSProviderType.md) |  | 
**Resources** | [**[]AWSProviderResource**](AWSProviderResource.md) |  | 

## Methods

### NewAWSProvider

`func NewAWSProvider(type_ AWSProviderType, resources []AWSProviderResource, ) *AWSProvider`

NewAWSProvider instantiates a new AWSProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSProviderWithDefaults

`func NewAWSProviderWithDefaults() *AWSProvider`

NewAWSProviderWithDefaults instantiates a new AWSProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSProvider) GetType() AWSProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSProvider) GetTypeOk() (*AWSProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSProvider) SetType(v AWSProviderType)`

SetType sets Type field to given value.


### GetResources

`func (o *AWSProvider) GetResources() []AWSProviderResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AWSProvider) GetResourcesOk() (*[]AWSProviderResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AWSProvider) SetResources(v []AWSProviderResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


