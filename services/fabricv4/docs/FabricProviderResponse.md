# FabricProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FabricProviderType**](FabricProviderType.md) |  | 
**Resources** | Pointer to [**[]FabricProviderResourceResponse**](FabricProviderResourceResponse.md) |  | [optional] 

## Methods

### NewFabricProviderResponse

`func NewFabricProviderResponse(type_ FabricProviderType, ) *FabricProviderResponse`

NewFabricProviderResponse instantiates a new FabricProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricProviderResponseWithDefaults

`func NewFabricProviderResponseWithDefaults() *FabricProviderResponse`

NewFabricProviderResponseWithDefaults instantiates a new FabricProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricProviderResponse) GetType() FabricProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricProviderResponse) GetTypeOk() (*FabricProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricProviderResponse) SetType(v FabricProviderType)`

SetType sets Type field to given value.


### GetResources

`func (o *FabricProviderResponse) GetResources() []FabricProviderResourceResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *FabricProviderResponse) GetResourcesOk() (*[]FabricProviderResourceResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *FabricProviderResponse) SetResources(v []FabricProviderResourceResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *FabricProviderResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


