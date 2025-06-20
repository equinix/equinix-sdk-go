# AWSProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AWSProviderType**](AWSProviderType.md) |  | 
**Resources** | Pointer to [**[]AWSProviderResourceResponse**](AWSProviderResourceResponse.md) |  | [optional] 

## Methods

### NewAWSProviderResponse

`func NewAWSProviderResponse(type_ AWSProviderType, ) *AWSProviderResponse`

NewAWSProviderResponse instantiates a new AWSProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSProviderResponseWithDefaults

`func NewAWSProviderResponseWithDefaults() *AWSProviderResponse`

NewAWSProviderResponseWithDefaults instantiates a new AWSProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AWSProviderResponse) GetType() AWSProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AWSProviderResponse) GetTypeOk() (*AWSProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AWSProviderResponse) SetType(v AWSProviderType)`

SetType sets Type field to given value.


### GetResources

`func (o *AWSProviderResponse) GetResources() []AWSProviderResourceResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AWSProviderResponse) GetResourcesOk() (*[]AWSProviderResourceResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AWSProviderResponse) SetResources(v []AWSProviderResourceResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AWSProviderResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


