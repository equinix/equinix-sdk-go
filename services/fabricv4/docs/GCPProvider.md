# GCPProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**GCPProviderType**](GCPProviderType.md) |  | 
**Resources** | [**[]GCPProviderResource**](GCPProviderResource.md) |  | 

## Methods

### NewGCPProvider

`func NewGCPProvider(type_ GCPProviderType, resources []GCPProviderResource, ) *GCPProvider`

NewGCPProvider instantiates a new GCPProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPProviderWithDefaults

`func NewGCPProviderWithDefaults() *GCPProvider`

NewGCPProviderWithDefaults instantiates a new GCPProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GCPProvider) GetType() GCPProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPProvider) GetTypeOk() (*GCPProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPProvider) SetType(v GCPProviderType)`

SetType sets Type field to given value.


### GetResources

`func (o *GCPProvider) GetResources() []GCPProviderResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GCPProvider) GetResourcesOk() (*[]GCPProviderResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GCPProvider) SetResources(v []GCPProviderResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


