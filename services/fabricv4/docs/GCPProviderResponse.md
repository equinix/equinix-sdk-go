# GCPProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**GCPProviderType**](GCPProviderType.md) |  | 
**Resources** | Pointer to [**[]GCPProviderResourceResponse**](GCPProviderResourceResponse.md) |  | [optional] 

## Methods

### NewGCPProviderResponse

`func NewGCPProviderResponse(type_ GCPProviderType, ) *GCPProviderResponse`

NewGCPProviderResponse instantiates a new GCPProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPProviderResponseWithDefaults

`func NewGCPProviderResponseWithDefaults() *GCPProviderResponse`

NewGCPProviderResponseWithDefaults instantiates a new GCPProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GCPProviderResponse) GetType() GCPProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPProviderResponse) GetTypeOk() (*GCPProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPProviderResponse) SetType(v GCPProviderType)`

SetType sets Type field to given value.


### GetResources

`func (o *GCPProviderResponse) GetResources() []GCPProviderResourceResponse`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *GCPProviderResponse) GetResourcesOk() (*[]GCPProviderResourceResponse, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *GCPProviderResponse) SetResources(v []GCPProviderResourceResponse)`

SetResources sets Resources field to given value.

### HasResources

`func (o *GCPProviderResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


