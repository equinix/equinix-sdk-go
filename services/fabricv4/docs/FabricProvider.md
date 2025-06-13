# FabricProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FabricProviderType**](FabricProviderType.md) |  | 
**Resources** | [**[]FabricProviderResource**](FabricProviderResource.md) |  | 

## Methods

### NewFabricProvider

`func NewFabricProvider(type_ FabricProviderType, resources []FabricProviderResource, ) *FabricProvider`

NewFabricProvider instantiates a new FabricProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricProviderWithDefaults

`func NewFabricProviderWithDefaults() *FabricProvider`

NewFabricProviderWithDefaults instantiates a new FabricProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FabricProvider) GetType() FabricProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricProvider) GetTypeOk() (*FabricProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricProvider) SetType(v FabricProviderType)`

SetType sets Type field to given value.


### GetResources

`func (o *FabricProvider) GetResources() []FabricProviderResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *FabricProvider) GetResourcesOk() (*[]FabricProviderResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *FabricProvider) SetResources(v []FabricProviderResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


